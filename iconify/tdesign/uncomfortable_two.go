package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UncomfortableTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm6.619-4.862L11 9.04v1.46l-3.2 2.4l-1.2-1.6l1.86-1.394L6.637 8.88l.98-1.743Zm9.743 1.743L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902l.98 1.743ZM9.5 15.5a1 1 0 0 0-1 1v1h-2v-1a3 3 0 0 1 5.065-2.177a.987.987 0 0 0 .206.158a.225.225 0 0 0 .044.019h.37a.223.223 0 0 0 .044-.019a.987.987 0 0 0 .207-.158A3 3 0 0 1 17.5 16.5v1h-2.001v-1a1 1 0 0 0-1.688-.726c-.322.306-.878.726-1.621.726h-.382c-.743 0-1.299-.42-1.621-.726A.993.993 0 0 0 9.5 15.5Zm2.682-1h.001Zm-.364 0h-.001Z"/>`),
		g.Group(children),
	)
}