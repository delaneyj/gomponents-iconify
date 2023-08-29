package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ws(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackWs0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackWs2)"><use href="#flagpackWs0"/><path fill="#C51918" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackWs1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackWs1)"><path fill="#2E42A5" d="M0 0v14h16V0H0Z"/><path fill="#FEFFFF" d="m3.566 7.772l-1.06.64l.241-1.249l-.882-.934l1.195-.051l.506-1.166l.505 1.166h1.194l-.881.985l.265 1.249l-1.083-.64Zm8 0l-1.06.64l.241-1.249l-.882-.934l1.195-.051l.506-1.166l.505 1.166h1.194l-.881.985l.265 1.249l-1.083-.64Zm-4.1-3.634l-.998.602l.227-1.175l-.83-.88l1.125-.047l.476-1.098l.476 1.098h1.123l-.83.927l.25 1.175l-1.02-.602Zm1.08 4.287l-.623.377l.142-.735l-.52-.55l.704-.03l.297-.685l.298.686h.702l-.518.58l.156.734l-.638-.377Zm-1.137 4.403l-1.497.904l.342-1.763l-1.247-1.32l1.688-.071l.714-1.646l.714 1.646h1.684l-1.243 1.39l.374 1.764l-1.53-.904Z"/></g></g><defs><clipPath id="flagpackWs2"><use href="#flagpackWs0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}