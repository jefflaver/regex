// regex is a test package of a simple regex, supporting '.', '*', 'a-z'
// using a stack instead of
package regex

type regexState struct {
	regex, test string
}

// Match takes a regex expression and tests to see if it matches the
// provided input
func Match(regex, test string) bool {
	stateStack := make([]regexState, 0)
	stateStack = append(stateStack, regexState{regex, test})
	for {
		if len(stateStack) == 0 {
			break
		}
		// Get the current state information
		state := regexState{}
		state, stateStack = stateStack[len(stateStack)-1],
			stateStack[:len(stateStack)-1]
		regex, test = state.regex, state.test

		// Check for a complete match
		if len(regex) == 0 && len(test) == 0 {
			return true
		} else if len(regex) == 0 {
			continue
		}

		// For each of the possible next states, push them onto the stack
		if len(regex) > 1 && regex[1] == '*' {
			if len(test) > 0 && test[0] == regex[0] {
				stateStack = append(stateStack, regexState{regex, test[1:]})
			}
			stateStack = append(stateStack, regexState{regex[2:], test})
		} else if len(test) == 0 {
			continue
		} else if regex[0] == '.' || regex[0] == test[0] {
			stateStack = append(stateStack, regexState{regex[1:], test[1:]})
		}
	}

	return false
}
