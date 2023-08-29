package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LicenseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 6H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h15l.57-.7l.93-1.14l-.09-.16H4V8h28v8.56a8.41 8.41 0 0 1 2 1.81V8a2 2 0 0 0-2-2Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M7 12h17v1.6H7z" class="clr-i-outline clr-i-outline-path-2"/><path fill="currentColor" d="M7 16h11v1.6H7z" class="clr-i-outline clr-i-outline-path-3"/><path fill="currentColor" d="M7 23h10v1.6H7z" class="clr-i-outline clr-i-outline-path-4"/><path fill="currentColor" d="M27.46 17.23a6.36 6.36 0 0 0-4.4 11l-1.94 2.37l.9 3.61l3.66-4.46a6.26 6.26 0 0 0 3.55 0l3.66 4.46l.9-3.61l-1.94-2.37a6.36 6.36 0 0 0-4.4-11Zm0 10.68a4.31 4.31 0 1 1 4.37-4.31a4.35 4.35 0 0 1-4.37 4.31Z" class="clr-i-outline clr-i-outline-path-5"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}