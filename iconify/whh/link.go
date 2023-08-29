package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M989 820L820 989q-34 35-83.5 35T652 989L483 820q-32-32-34-77.5t26-80.5l214 215q20 19 47.5 19t46.5-19l94-94q19-19 19-46.5T877 689L662 475q35-28 80.5-26t77.5 34l169 169q35 35 35 84.5T989 820zM341 341q19-19 45-20t44 17l257 258q18 17 17 43.5T684 685t-45 19.5t-44-16.5L338 430q-18-18-17-44t20-45zm-6-193q-19-20-46.5-20T241 148l-93 93q-20 20-20 47.5t20 46.5l214 214q-35 28-80.5 26T204 541L35 372Q0 337 0 287.5T35 203L204 35q34-35 83.5-35T372 35l169 168q32 33 34 78.5T549 362z"/>`),
		g.Group(children),
	)
}