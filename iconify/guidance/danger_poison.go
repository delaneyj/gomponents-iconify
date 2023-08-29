package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DangerPoison(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14 18.501c-.185 0-.712-.229-.945-.46a3.952 3.952 0 0 1-.762-1.064C12.143 16.67 12 16.299 12 16c0 .299-.144.67-.293.977c-.2.409-.462.765-.762 1.063c-.233.232-.76.461-.945.461M2.5 18.5V10a9.5 9.5 0 1 1 19 0v8.5a5 5 0 0 0-5 5h-9a5 5 0 0 0-5-5Zm8-8V14c-1 1-4 1.5-5 1.5v-1.015c0-1.591.69-3.117 2.225-3.538c1.087-.298 2.23-.447 2.775-.447Zm3 0V14c1 1 4 1.5 5 1.5v-1.015c0-1.591-.69-3.117-2.225-3.538c-1.087-.298-2.23-.447-2.775-.447Z"/>`),
		g.Group(children),
	)
}