package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 1024H128q-53 0-90.5-37.5T0 896V512q0-63 18.5-122.5T67 293t61-37q27 0 45.5-18.5t18.5-45t-15-66.5q-21-5-35-22.5T128 64q0-26 19-45t45-19h512q26 0 45 19t19 45q0 22-14 39.5T719 126q-15 40-15 66q0 27 18.5 45.5T768 256q31 0 61 37t48.5 96.5T896 512v384q0 53-37.5 90.5T768 1024zm0-469q0-72-16-121.5T704 384q-39 0-83.5-34.5T576 288V128H320v160q0 27-44.5 61.5T192 384q-32 0-48 50t-16 121v256q0 35 27 60t64 25h458q37 0 64-25t27-60V555z"/>`),
		g.Group(children),
	)
}