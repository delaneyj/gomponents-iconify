package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M344.766 288h-87.741l38.054 25.748a21.894 21.894 0 0 1 9.558 16.187a20.653 20.653 0 0 1-6.058 16.824C294.7 350.584 286.705 357 276.677 357h-104v54h104c21.722 0 42.972-9.165 59.835-25.808a74.638 74.638 0 0 0 21.9-60.218A75.831 75.831 0 0 0 344.766 288Zm-129.34-93.347c-12.195-8.251-9.725-20.755-8.677-24.389c1.269-4.4 5.647-14.643 18.964-15.248c6.468-.018 118.281 0 118.281 0l.012-54s-119.176-.006-119.795.015c-32.949 1.1-60.171 22.419-69.351 54.278c-7.125 24.726-1.819 49.847 13.481 68.688H80v32h352V224H258.8Z"/>`),
		g.Group(children),
	)
}