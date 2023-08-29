package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveUndo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1728 1152q67 0 125 25t102 69t68 103t25 126q0 91-34 171t-97 146q-26 26-53 50t-54 48q-45 40-88 79t-88 79l-98-91q46-39 91-78t90-79q28-24 55-48t54-51q47-48 70-104t24-125q0-41-15-76t-42-61t-62-40t-76-15q-66 0-116 21t-92 58t-77 82t-69 95h293v128h-512v-512h128v286q46-61 91-113t96-90t115-61t146-22zM0 128h2048v640h-128v256h-128V768H256v1024h1152v128H128V768H0V128zm1920 512V256H128v384h1792zm-896 512H640v-128h384v128z"/>`),
		g.Group(children),
	)
}