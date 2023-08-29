package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monstra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 544q0 109-52.5 202T638 895q2 18 2 33q0 40-9.5 68t-22.5 28q-15 0-48.5-23.5T515 948q-50 12-99 12t-99-12q-11 29-44.5 52.5T224 1024q-13 0-22.5-28t-9.5-68q0-16 2-33q-89-56-141.5-149T0 544q0-108 51-200t139-149t194-65v-12q-25-14-44.5-43.5T320 32t28-22.5T416 0t68 9.5T512 32t-19.5 42.5T448 118v12q106 8 194 65t139 149t51 200zM416 896q60 0 94-16t66-48H256q29 29 65 46.5t95 17.5zm0-576q-93 0-158.5 65.5T192 544t65.5 158.5T416 768t158.5-65.5T640 544t-65.5-158.5T416 320zm0 320q-40 0-68-28t-28-68t28-68t68-28q21 0 41 10q-9 9-9 22t9.5 22.5T480 512t22-9q10 20 10 41q0 40-28 68t-68 28z"/>`),
		g.Group(children),
	)
}