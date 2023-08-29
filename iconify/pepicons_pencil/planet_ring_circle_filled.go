package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetRingCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPlanetRingCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.126 6.923a5.976 5.976 0 1 0 0 11.952a5.976 5.976 0 0 0 0-11.952ZM6.15 12.9a6.976 6.976 0 1 1 13.951 0a6.976 6.976 0 0 1-13.951 0Z"/><path d="M5.3 16.528c-.337.667-.288.985-.208 1.102c.067.098.276.238.838.224c.536-.014 1.26-.168 2.132-.469c1.739-.6 3.958-1.745 6.235-3.302c2.268-1.55 3.985-3.189 5.022-4.574c.52-.695.854-1.305 1.01-1.792c.16-.502.104-.772.017-.898c-.014-.021-.047-.063-.193-.078c-.162-.017-.402.008-.739.092c-.63.157-1.441.476-2.422.862l-.207.081l-.366-.93l.234-.093c.95-.374 1.822-.717 2.52-.89c.378-.094.75-.151 1.082-.117c.348.036.69.177.917.508c.345.505.306 1.152.109 1.768c-.203.632-.605 1.343-1.162 2.087c-1.117 1.492-2.922 3.203-5.258 4.8c-2.329 1.592-4.628 2.785-6.473 3.421c-.918.317-1.752.506-2.433.523c-.656.017-1.324-.125-1.688-.658c-.42-.614-.216-1.413.14-2.119c.379-.748 1.033-1.602 1.88-2.494l.725.69c-.815.857-1.396 1.63-1.712 2.256Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPlanetRingCircleFilled0)"/></g>`),
		g.Group(children),
	)
}