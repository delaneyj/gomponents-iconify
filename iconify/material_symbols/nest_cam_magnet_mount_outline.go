package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamMagnetMountOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 16.3q-.825 0-1.413-.588T21 14.3v-4q0-.85.588-1.425T23 8.3h1v8h-1ZM7.275 18.275q-.375 0-.75-.137T5.85 17.7L1.6 13.475q-.3-.3-.45-.662T1 12.075q0-.375.138-.75t.437-.675l2.875-2.9q.725-.725 1.638-1.1T8 6.275q2.1 0 3.55 1.45t1.45 3.55q0 1-.375 1.913t-1.1 1.637l-2.85 2.85q-.3.3-.663.45t-.737.15Zm.725-10q-.575 0-1.125.225t-1 .675L3 12.05l4.25 4.25l2.875-2.875q.45-.45.675-1t.225-1.125q0-1.275-.875-2.15T8 8.275Zm-.975 4Zm9.925 3.075L15.9 14.3l1.2-1.2H14v-1.5h3.2l-1.25-1.25L17 9.3l3 3l-3.05 3.05Z"/>`),
		g.Group(children),
	)
}