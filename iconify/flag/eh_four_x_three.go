package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EhFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagEh4x30"><path fill-opacity=".7" d="M-158.7 0H524v512h-682.7z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagEh4x30)" transform="translate(148.8) scale(.94)"><path d="M-158.3 0h680.9v255.3h-680.9z"/><path fill="#007a3d" d="M-158.3 255.3h680.9v255.3h-680.9z"/><path fill="#fff" d="M-158.3 148.9h680.9v212.8h-680.9z"/><path fill="#c4111b" d="m-158.3 0l340.4 255.3l-340.4 255.3Z"/><circle cx="352.3" cy="255.3" r="68.1" fill="#c4111b"/><circle cx="377.9" cy="255.3" r="68.1" fill="#fff"/><path fill="#c4111b" d="m334 296.5l29.1-20.7l28.8 21l-10.8-34l29-20.9l-35.7-.2l-11-34l-11.2 33.9l-35.7-.2l28.7 21.2l-11.1 34z"/></g>`),
		g.Group(children),
	)
}