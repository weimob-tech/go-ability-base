package hook

var (
	PreStartHook  = make([]func(), 0)
	PostStartHook = make([]func(), 0)
)

func AddPreStartHook(hook func()) {
	PreStartHook = append(PreStartHook, hook)
}

func ApplyPreStartHook() {
	for _, hook := range PreStartHook {
		hook()
	}
}

func AddPostStartHook(hook func()) {
	PostStartHook = append(PostStartHook, hook)
}

func ApplyPostStartHook() {
	for _, hook := range PostStartHook {
		hook()
	}
}
