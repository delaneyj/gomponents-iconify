package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1088 0q185 0 316.5 93.5T1536 320v896q0 130-125.5 222t-305.5 97l213 202q16 15 8 35t-30 20H240q-22 0-30-20t8-35l213-202q-180-5-305.5-97T0 1216V320Q0 187 131.5 93.5T448 0h640zM288 1312q66 0 113-47t47-113t-47-113t-113-47t-113 47t-47 113t47 113t113 47zm416-544V256H160v512h544zm544 544q66 0 113-47t47-113t-47-113t-113-47t-113 47t-47 113t47 113t113 47zm160-544V256H832v512h576z"/>`),
		g.Group(children),
	)
}