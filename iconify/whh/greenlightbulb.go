package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Greenlightbulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M719 391L147 639q-26 8-49.5-6.5T67 591t6.5-51.5T113 507l572-248q27-7 50 7t30.5 41.5t-6.5 52t-40 31.5zm-64-192L83 447q-26 8-49.5-6.5T3 399t6.5-51.5T49 316L621 67q27-7 50 7t30.5 41.5t-6.5 52t-40 31.5zm-352-64l-156 56q-26 8-49.5-6T67 143.5t6.5-52T113 60L270 3q26-7 49.5 7T350 51.5t-6.5 52T303 135zm-21 513q10-9 23-13l316-120q27-7 50 7t30.5 41.5t-6.5 52t-40 31.5L385 750q-1 13-1 19h64q27 0 45.5 19t18.5 45.5t-18.5 45T448 897H320q-6 0-12-1q-49-5-82.5-41.5T192 769q0-42 25-75.5t65-45.5zm102 377q-53 0-90.5-18.5T256 961h256q0 27-37.5 45.5T384 1025z"/>`),
		g.Group(children),
	)
}