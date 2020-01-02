package accUnexp

// The identifier below is unexported (it starts with a lowercase letter)
type alertCounter int

// The function below is exported (it starts with an uppercase letter)
// New creates and returns values of the UNEXPORTED TYPE alertCounter
func New(value int) alertCounter {
    return alertCounter(value)
}
