package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveFromTrash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1086 91L281 896h999v960q0 40-15 75t-41 61t-61 41t-75 15H320q-40 0-75-15t-61-41t-41-61t-15-75v-807l-37 37l-91-90l272-272l-91-90q-18-18-27-41t-10-50q0-51 37-90l271-272q18-18 41-27t50-10q26 0 49 9t42 28l90 91L996 0l90 91zm66 1765v-832H256v832q0 26 19 45t45 19h768q26 0 45-19t19-45zM543 272L272 543l90 91l272-272l-91-90zm1377-16h128v512h-512V640h292q-77-60-167-91t-188-31q-115 0-219 43t-185 124l-90-90q100-100 226-152t268-53q123 0 238 41t209 119V256zm-896 896v640H896v-640h128zm-256 0v640H640v-640h128zm-256 0v640H384v-640h128z"/>`),
		g.Group(children),
	)
}