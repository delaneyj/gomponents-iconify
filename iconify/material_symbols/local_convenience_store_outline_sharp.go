package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalConvenienceStoreOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.025 11.05V21h-18v-9.95q-.575-.525-.887-1.35t-.013-1.8L3.65 3H20.4l1.525 4.9q.3.975-.013 1.775t-.887 1.375Zm-6.8-1.05q.675 0 1.025-.463t.275-1.037l-.55-3.5h-1.95v3.7q0 .525.35.913t.85.387Zm-4.5 0q.575 0 .938-.388t.362-.912V5h-1.95l-.55 3.5q-.1.6.262 1.05t.938.45Zm-4.45 0q.45 0 .788-.325t.412-.825L7.025 5h-1.95l-1 3.35q-.15.5.163 1.075T5.275 10Zm13.5 0q.725 0 1.05-.575t.15-1.075L18.925 5h-1.9l.55 3.85q.075.5.413.825t.787.325Zm-13.75 9h14v-7.05q-.125.05-.162.05h-.088q-.675 0-1.187-.225t-1.013-.725q-.45.45-1.025.7t-1.225.25q-.675 0-1.263-.25t-1.037-.7q-.425.45-.988.7T9.825 12q-.725 0-1.312-.25t-1.038-.7q-.525.525-1.037.738T5.274 12h-.113q-.062 0-.137-.05V19Zm14 0h-14h14Zm-11-1h3v-1h-2v-1h2v-3h-3v1h2v1h-2v3Zm7 0h1v-5h-1v2h-1v-2h-1v3h2v2Z"/>`),
		g.Group(children),
	)
}