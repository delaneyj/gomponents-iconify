package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HoodieBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m246 121.34l-56.31-84.43A20 20 0 0 0 173.05 28H83a20 20 0 0 0-16.64 8.91L10 121.34a12 12 0 0 0-1.73 9.28l18.39 82.11c0 .18.09.36.14.53A20 20 0 0 0 46.11 228H76a20 20 0 0 0 20-20v-8h64v8a20 20 0 0 0 20 20h29.89a20 20 0 0 0 19.29-14.74c0-.17.1-.35.14-.53l18.39-82.11a12 12 0 0 0-1.71-9.28ZM128 74.4L86 52h84ZM80 176V76l16 8.53V136a12 12 0 0 0 24 0V97.33l2.35 1.26a12 12 0 0 0 11.3 0l2.35-1.26V128a12 12 0 0 0 24 0V84.53L176 76v100Zm-30.68 28l-16.49-73.61L56 95.63V180a20 20 0 0 0 16 19.6v4.4Zm157.36 0H184v-4.4a20 20 0 0 0 16-19.6V95.63l23.17 34.76Z"/>`),
		g.Group(children),
	)
}