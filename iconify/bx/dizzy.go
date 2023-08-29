package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dizzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm0 18c-4.411 0-8-3.589-8-8s3.589-8 8-8s8 3.589 8 8s-3.589 8-8 8z"/><path fill="currentColor" d="M10.707 12.293L9.414 11l1.293-1.293l-1.414-1.414L8 9.586L6.707 8.293L5.293 9.707L6.586 11l-1.293 1.293l1.414 1.414L8 12.414l1.293 1.293zm6.586-4L16 9.586l-1.293-1.293l-1.414 1.414L14.586 11l-1.293 1.293l1.414 1.414L16 12.414l1.293 1.293l1.414-1.414L17.414 11l1.293-1.293zM10 16h4v2h-4z"/>`),
		g.Group(children),
	)
}