package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeftaFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#039" d="M0 0h640v480H0z"/><circle cx="320" cy="249.8" r="30.4" fill="none" stroke="#fc0" stroke-width="27.5"/><circle cx="320" cy="249.8" r="88.3" fill="none" stroke="#fc0" stroke-width="27.5"/><path fill="#039" d="m404.7 165.1l84.7 84.7l-84.7 84.7l-84.7-84.7z"/><path fill="#fc0" d="M175.7 236.1h59.2v27.5h-59.2zm259.1 0h88.3v27.5h-88.3zM363 187.4l38.8-38.8l19.4 19.5l-38.7 38.7zM306.3 48.6h27.5v107.1h-27.5z"/><circle cx="225.1" cy="159.6" r="13.7" fill="#fc0"/><circle cx="144.3" cy="249.8" r="13.7" fill="#fc0"/><circle cx="320" cy="381.4" r="13.7" fill="#fc0"/><circle cx="320" cy="425.5" r="13.7" fill="#fc0"/><circle cx="408.3" cy="249.8" r="13.7" fill="#fc0"/><path fill="#fc0" d="m208.3 341.5l19.5-19.4l19.4 19.4l-19.4 19.5zm204.7 21l19.5-19.5l19.5 19.5l-19.5 19.4z"/>`),
		g.Group(children),
	)
}