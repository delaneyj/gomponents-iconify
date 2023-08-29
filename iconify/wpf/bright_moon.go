package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="m9.5.063l-.625 2.28l-2.25-.593L8.25 3.469L6.625 5.156l2.25-.593l.625 2.312l.625-2.313l2.25.594L10.75 3.47l1.625-1.719l-2.25.594L9.5.063zM2.469 6.5l-.5 1.813l-1.813-.47L1.47 9.189L.156 10.53l1.813-.469l.5 1.813L3 10.062l1.813.47L3.5 9.186l1.313-1.343L3 8.312L2.469 6.5zm15.625.688a8.273 8.273 0 0 1 4.937 7.593c0 4.581-3.685 8.281-8.25 8.281a8.29 8.29 0 0 1-7.594-4.968a9.39 9.39 0 0 0 9.25 7.718a9.361 9.361 0 0 0 9.375-9.375a9.396 9.396 0 0 0-7.718-9.25z"/>`),
		g.Group(children),
	)
}