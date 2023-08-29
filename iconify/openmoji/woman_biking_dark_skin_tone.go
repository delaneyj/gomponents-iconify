package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomanBikingDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#6a462f"><circle cx="32.793" cy="10.581" r="2.86"/><path d="m32.793 19.923l1.906-2.764l3.623.095l6.005 13.441l-2.002 3.623l-2.383 2.383l-2.955 10.677l-3.241 6.387l-2.288-1.907l.858-6.768l-3.622-6.673l.095-3.146l7.626-4.766l-2.67-7.15"/></g><g fill="none" stroke="#000" stroke-width="2"><circle cx="19.161" cy="49.57" r="9.723" stroke-miterlimit="10"/><circle cx="49.666" cy="49.57" r="9.723" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m19.256 49.57l1.907-22.878h3.622"/><circle cx="32.793" cy="10.581" r="3" stroke-miterlimit="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m25.071 26.787l4.1-2.765a11.819 11.819 0 0 0 2.668-2.669l1.812-2.574a3.942 3.942 0 0 1 2.955-1.525a3.055 3.055 0 0 1 2.574 1.811l4.385 10.01a4.742 4.742 0 0 1-.095 3.432l-.096.19a5.875 5.875 0 0 1-2.67 2.383l-5.242 2.288a1.935 1.935 0 0 0-1.049 2.478l2.479 6.101a2.285 2.285 0 0 1-.668 2.67a1.73 1.73 0 0 1-2.383-.763l-4.194-7.817a8.648 8.648 0 0 1-.858-3.241a3.666 3.666 0 0 1 1.621-2.574l4.385-2.764a2.784 2.784 0 0 0 1.144-2.86l-1.43-5.72"/><path stroke-miterlimit="10" d="m36.224 48.617l-.953 2.955a2.66 2.66 0 0 1-2.288 1.811a1.622 1.622 0 0 1-1.525-1.906l.858-6.387m8.103-9.914l-3.432 10.962"/><path stroke-linecap="round" stroke-linejoin="round" d="M36.125 8.278s.458-.472.898-.323c.517.175 1.248 1.289 1.913 1.638a3.364 3.364 0 0 0 2.14.366"/></g>`),
		g.Group(children),
	)
}