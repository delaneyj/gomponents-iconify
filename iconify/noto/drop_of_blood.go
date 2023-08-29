package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropOfBlood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#E5322E" d="M88.04 44.26C75.44 28.81 69.87 15.67 66.68 7.83c-1.01-2.47-4.36-2.47-5.36 0c-3.19 7.84-8.76 20.98-21.36 36.43c-7.89 9.68-15.67 23.85-15.67 37.81c0 23.82 17.19 39.95 39.71 39.95s39.71-16.14 39.71-39.95c0-14.51-7.78-28.13-15.67-37.81z"/><path fill="#FF6050" d="M74.98 79.84c6.35-12.08 5.45-23.9 10.47-21.77c6.82 2.91 14.37 17.86 11.54 31.41c-2.02 9.66-8.54 15.51-16.85 12.72c-6.71-2.25-10.88-11.48-5.16-22.36z"/>`),
		g.Group(children),
	)
}