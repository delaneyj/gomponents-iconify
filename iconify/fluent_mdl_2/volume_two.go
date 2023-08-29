package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1402 390q126 128 194 291t68 343q0 180-68 343t-194 291l-91-91q109-109 167-249t58-294q0-154-58-294t-167-249l91-91zm-272 272q73 73 111 166t39 196q0 103-38 196t-112 166l-90-90q54-54 83-124t29-148q0-77-29-147t-83-125l90-90zM677 256h91v1536h-90l-385-384H0V640h293l384-384zm-37 219L347 768H128v512h219l293 293V475z"/>`),
		g.Group(children),
	)
}