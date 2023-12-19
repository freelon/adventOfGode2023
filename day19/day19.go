package day19

import (
	"maps"
	"strconv"
	"strings"
)

func Part1(input string) string {
	workflows, parts := parse(input)
	sum := 0
	for _, part := range parts {
		var cw Workflow
		next := "in"
		for next != "A" && next != "R" {
			cw = workflows[next]
			for _, rule := range cw.rules {
				if rule.condition == nil || matches(rule.condition, part) {
					next = rule.result
					break
				}
			}
		}
		if next == "A" {
			sum += part["x"]
			sum += part["m"]
			sum += part["a"]
			sum += part["s"]
		}
	}
	return strconv.Itoa(sum)
}

func matches(condition *Condition, part map[string]int) bool {
	v := part[condition.variable]
	if condition.comparator == "<" {
		return v < condition.target
	} else {
		return v > condition.target
	}
}

func parse(input string) (workflows map[string]Workflow, machineParts []map[string]int) {
	workflows = make(map[string]Workflow)
	blocks := strings.Split(input, "\n\n")
	for _, line := range strings.Split(blocks[0], "\n") {
		name := line[:strings.Index(line, "{")]
		w := Workflow{name: name, rules: make([]Rule, 0)}
		for _, r := range strings.Split(strings.TrimSuffix(line[len(name)+1:], "}"), ",") {
			var result string
			var condition *Condition
			if parts := strings.Split(r, ":"); len(parts) == 1 {
				result = r
			} else {
				result = parts[1]
				var comparator = ">"
				i := strings.Index(parts[0], comparator)
				if i < 0 {
					comparator = "<"
					i = strings.Index(parts[0], comparator)
				}
				condition = &Condition{}
				condition.variable = parts[0][:i]
				condition.comparator = comparator
				v, _ := strconv.Atoi(parts[0][i+1:])
				condition.target = v
			}
			rule := Rule{condition: condition, result: result}
			w.rules = append(w.rules, rule)
		}
		workflows[w.name] = w
	}
	for _, line := range strings.Split(blocks[1], "\n") {
		line = line[1 : len(line)-1]
		part := make(map[string]int)
		for _, s := range strings.Split(line, ",") {
			i := strings.Index(s, "=")
			k := s[:i]
			v, _ := strconv.Atoi(s[i+1:])
			part[k] = v
		}
		machineParts = append(machineParts, part)
	}
	return
}

type Workflow struct {
	name  string
	rules []Rule
}

type Rule struct {
	condition *Condition
	result    string
}

type Condition struct {
	variable   string
	comparator string
	target     int
}

func (c Condition) invert() *Condition {
	var com string
	var target int
	if c.comparator == "<" {
		com = ">"
		target = c.target - 1
	} else {
		com = "<"
		target = c.target + 1
	}
	r := Condition{
		variable:   c.variable,
		comparator: com,
		target:     target,
	}
	return &r
}

func Part2(input string) string {
	workflows, _ := parse(input)
	All := Foo{mins: map[string]int{"x": 1, "m": 1, "a": 1, "s": 1}, maxs: map[string]int{"x": 4000, "m": 4000, "a": 4000, "s": 4000}}
	acceptables := count(All, workflows["in"], workflows)
	sum := 0
	for _, acceptable := range acceptables {
		x := 1
		for k, v := range acceptable.mins {
			x *= acceptable.maxs[k] - v + 1
		}
		sum += x
	}
	return strconv.Itoa(sum)
}

func count(foo Foo, workflow Workflow, workflows map[string]Workflow) (result []Foo) {
	if foo.empty() {
		return make([]Foo, 0)
	}
	foo = foo.clone()
	for _, rule := range workflow.rules {
		next := rule.result
		newFoo := foo.restrict(rule.condition)
		if rule.condition != nil {
			foo = foo.restrict(rule.condition.invert())
		}
		if next == "R" {
			continue
		}
		if next == "A" {
			result = append(result, newFoo)
			continue
		}
		subresult := count(newFoo, workflows[next], workflows)
		result = append(result, subresult...)
	}
	return
}

type Foo struct {
	mins map[string]int
	maxs map[string]int
}

func (f Foo) empty() bool {
	for k, v := range f.mins {
		if v > f.maxs[k] {
			return true
		}
	}
	return false
}

func (f Foo) restrict(condition *Condition) Foo {
	if condition == nil {
		return f
	}
	result := f.clone()
	if condition.comparator == "<" {
		result.maxs[condition.variable] = condition.target - 1
	} else {
		result.mins[condition.variable] = condition.target + 1
	}
	return result
}

func (f Foo) clone() Foo {
	return Foo{
		mins: maps.Clone(f.mins),
		maxs: maps.Clone(f.maxs),
	}
}
