package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trumpet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024q-73-89-148-176q-16 33-37 54q-58 58-139 58t-139-58L314 784q-57-58-57-139t57-139q17-17 43-32q-143-90-292-90L9 372q-9-10-9-23t9-23L326 9q10-9 23-9t23 9l13 56q0 103 46.5 212.5T578 487l152 152q40 40 113.5 104.5t127 108.5l53.5 44zM393 585q-24 25-24 60t24 59l119 119q25 25 60 25t59-25q33-33 22-79q-9-10-14-14L487 578q-3-3-8.5-8t-7.5-7q-45-10-78 22z"/>`),
		g.Group(children),
	)
}