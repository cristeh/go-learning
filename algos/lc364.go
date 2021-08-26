package algos

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

type IntWeight struct {
	Int    int
	Weight int
}

func depthSumInverse(nestedList []*NestedInteger) int {
	maxDepth := 1
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	ints := make([]IntWeight, 0, 8)
	var helper func(nestedList []*NestedInteger, depth int)
	helper = func(nestedList []*NestedInteger, depth int) {
		maxDepth = max(maxDepth, depth)
		for _, nestedInt := range nestedList {
			if nestedInt.IsInteger() {
				ints = append(ints, IntWeight{
					Int:    nestedInt.GetInteger(),
					Weight: depth,
				})
			} else {
				helper(nestedInt.GetList(), depth+1)
			}
		}
	}
	helper(nestedList, 1)
	result := 0
	for _, intWeight := range ints {
		result += intWeight.Int * (maxDepth - intWeight.Weight + 1)
	}
	return result
}
