package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommonDataServiceCDS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1523 0l525 525v998l-525 525H525L0 1523V525L525 0h998zM512 1854l13-331l-397-384v331l384 384zm0-512v-331L128 627v331l384 384zm0-512V194L194 512l318 318zm128-702v768h830l384-384l-384-384H640zm0 1011v269h269l-269-269zm0 397v268l268-268H640zm1280-66v-331l-397 397h-434l-384 384h765l450-450zm0-512V627l-397 397H706l384 384h380l450-450z"/>`),
		g.Group(children),
	)
}