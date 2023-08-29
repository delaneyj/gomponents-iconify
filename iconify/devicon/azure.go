package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Azure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><linearGradient id="deviconAzure0" x1="60.919" x2="18.667" y1="9.602" y2="134.423" gradientUnits="userSpaceOnUse"><stop stop-color="#114A8B"/><stop offset="1" stop-color="#0669BC"/></linearGradient><linearGradient id="deviconAzure1" x1="74.117" x2="64.344" y1="67.772" y2="71.076" gradientUnits="userSpaceOnUse"><stop stop-opacity=".3"/><stop offset=".071" stop-opacity=".2"/><stop offset=".321" stop-opacity=".1"/><stop offset=".623" stop-opacity=".05"/><stop offset="1" stop-opacity="0"/></linearGradient><linearGradient id="deviconAzure2" x1="68.742" x2="115.122" y1="5.961" y2="129.525" gradientUnits="userSpaceOnUse"><stop stop-color="#3CCBF4"/><stop offset="1" stop-color="#2892DF"/></linearGradient></defs><path fill="url(#deviconAzure0)" d="M46.09.002h40.685L44.541 125.137a6.485 6.485 0 0 1-6.146 4.413H6.733a6.482 6.482 0 0 1-5.262-2.699a6.474 6.474 0 0 1-.876-5.848L39.944 4.414A6.488 6.488 0 0 1 46.09 0z" transform="translate(.587 4.468) scale(.91904)"/><path fill="#0078d4" d="M97.28 81.607H37.987a2.743 2.743 0 0 0-1.874 4.751l38.1 35.562a5.991 5.991 0 0 0 4.087 1.61h33.574z"/><path fill="url(#deviconAzure1)" d="M46.09.002A6.434 6.434 0 0 0 39.93 4.5L.644 120.897a6.469 6.469 0 0 0 6.106 8.653h32.48a6.942 6.942 0 0 0 5.328-4.531l7.834-23.089l27.985 26.101a6.618 6.618 0 0 0 4.165 1.519h36.396l-15.963-45.616l-46.533.011L86.922.002z" transform="translate(.587 4.468) scale(.91904)"/><path fill="url(#deviconAzure2)" d="M98.055 4.408A6.476 6.476 0 0 0 91.917.002H46.575a6.478 6.478 0 0 1 6.137 4.406l39.35 116.594a6.476 6.476 0 0 1-6.137 8.55h45.344a6.48 6.48 0 0 0 6.136-8.55z" transform="translate(.587 4.468) scale(.91904)"/>`),
		g.Group(children),
	)
}