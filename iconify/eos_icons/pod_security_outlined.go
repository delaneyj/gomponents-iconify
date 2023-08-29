package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodSecurityOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.75 12l-2.25-1l-2.25 1L13 13v3a6.405 6.405 0 0 0 1.282 3.804A5.776 5.776 0 0 0 17.5 22a5.776 5.776 0 0 0 3.218-2.196A6.405 6.405 0 0 0 22 16v-3Zm.098 7.287a4.787 4.787 0 0 1-2.348 1.678V16.5H14v-2.85l1.75-.777l1.75-.778v4.4H21a5.803 5.803 0 0 1-1.152 2.792ZM13 4a1.004 1.004 0 0 0-1-1H8a1.004 1.004 0 0 0-1 1v1h6Z"/><circle cx="10" cy="11" r="1.5" fill="currentColor"/><path fill="currentColor" d="M12 17v-2H5.1l3.18-7h3.42l1.704 3.726l1.825-.811L12.99 6H7L2 17h10zm2 4a5.797 5.797 0 0 1-.519-.597A7.488 7.488 0 0 1 12.68 19H6v-1H4v1.003A1.998 1.998 0 0 0 6 21"/>`),
		g.Group(children),
	)
}