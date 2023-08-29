package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeTrackingOnSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M2 4a2 2 0 0 1 2-2h1.5a.5.5 0 0 1 0 1H4a1 1 0 0 0-1 1v1.5a.5.5 0 0 1-1 0V4z" fill="currentColor"/><path d="M2 12a2 2 0 0 0 2 2h1.5a.5.5 0 0 0 0-1H4a1 1 0 0 1-1-1v-1.5a.5.5 0 0 0-1 0V12z" fill="currentColor"/><path d="M12 2a2 2 0 0 1 2 2v1.5a.5.5 0 0 1-1 0V4a1 1 0 0 0-1-1h-1.5a.5.5 0 0 1 0-1H12z" fill="currentColor"/><path d="M14 12a2 2 0 0 1-2 2h-1.5a.5.5 0 0 1 0-1H12a1 1 0 0 0 1-1v-1.5a.5.5 0 0 1 1 0V12z" fill="currentColor"/><path d="M5.5 9a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0z" fill="currentColor"/><path d="M3.314 7.964a.5.5 0 0 0 .65-.278v.001l.005-.011l.028-.06c.027-.055.072-.137.136-.24a3.73 3.73 0 0 1 .64-.756C5.376 6.073 6.384 5.5 8 5.5c1.617 0 2.624.573 3.226 1.12c.305.277.512.553.64.757a2.752 2.752 0 0 1 .165.299s.244.452.655.288a.5.5 0 0 0 .278-.65c-.107-.223 0 0 0 0v-.002l-.002-.004l-.005-.01a1.809 1.809 0 0 0-.06-.129a3.783 3.783 0 0 0-.186-.327a4.733 4.733 0 0 0-.812-.962c-.774-.703-2.017-1.38-3.9-1.38c-1.882 0-3.125.677-3.898 1.38a4.73 4.73 0 0 0-.813.962a3.753 3.753 0 0 0-.246.456l-.004.01l-.002.005v.001a.5.5 0 0 0 .278.65z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}