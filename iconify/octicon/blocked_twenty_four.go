package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockedTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.638 2.22a.749.749 0 0 1 .53-.22h7.664c.199 0 .389.079.53.22l5.418 5.418c.141.14.22.332.22.53v7.664a.749.749 0 0 1-.22.53l-5.418 5.418a.749.749 0 0 1-.53.22H8.168a.749.749 0 0 1-.53-.22l-5.42-5.418a.752.752 0 0 1-.219-.53V8.168c0-.199.079-.389.22-.53l5.418-5.42ZM8.48 3.5L3.5 8.48v7.04l4.98 4.98h7.04l4.98-4.98V8.48L15.52 3.5ZM7 11.75a.75.75 0 0 1 .75-.75h8.5a.75.75 0 0 1 0 1.5h-8.5a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}