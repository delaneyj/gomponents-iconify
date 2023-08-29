package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostOffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#b2c1c0" d="M63 60H1c-.6 0-1 .5-1 1v2c0 .5.4 1 1 1h62c.5 0 1-.5 1-1v-2c0-.5-.5-1-1-1"/><path fill="#428bc1" d="M64 17c0 .5-.5 1-1 1H1c-.6 0-1-.5-1-1V3c0-.5.4-1 1-1h62c.5 0 1 .5 1 1v14"/><path fill="#62727a" d="M2 18h60v42H2z"/><path fill="#fff" d="M2 20h60v10H2zm0 12h60v10H2z"/><path fill="#d0d0d0" d="M2 0h60v2H2z"/><path fill="#ffe62e" d="M22.8 6c0 2.8-1.6 5.2-3.9 6.3c.6-.9 1-1.9 1-3.1c0-2.9-2.4-5.3-5.3-5.3c-2.8 0-5.1 2.2-5.3 4.9c-.4-.8-.6-1.7-.6-2.8H4c0 5.5 4.5 10 10 10s10-4.5 10-10h-1.2m-8.2-.2c1.9 0 3.5 1.6 3.5 3.5s-1.6 3.5-3.5 3.5s-3.5-1.6-3.5-3.5c0-2 1.5-3.5 3.5-3.5"/><path fill="#b4d7ee" d="M4.5 22h4v6h-4zm6 0h4v6h-4zm9 0h4v6h-4zm6 0h4v6h-4zm9 0h4v6h-4zm6 0h4v6h-4zm9 0h4v6h-4zm6 0h4v6h-4zm-51 12h4v6h-4zm6 0h4v6h-4zm9 0h4v6h-4zm6 0h4v6h-4zm9 0h4v6h-4zm6 0h4v6h-4zm9 0h4v6h-4zm6 0h4v6h-4z"/><path fill="#ffe62e" d="M21 44v16h2V48h18v12h2V44z"/><path fill="#b4d7ee" d="M27 48h4v12h-4zm6 0h4v12h-4z"/><g fill="#fff"><path d="M31 48h2v12h-2zm6 0h2v12h-2zm-12 0h2v12h-2z"/><path d="M26 53h12v2H26zm17-5h16v10H43zM5 48h16v10H5z"/></g><path fill="#b4d7ee" d="M43 50h14v6H43zM7 50h14v6H7z"/><path fill="#fff" d="M13 49h2v8h-2zm36 0h2v8h-2z"/><path fill="#ffe62e" d="M28 5.3c.6-.1 1.5-.2 2.8-.2c1.3 0 2.2.3 2.8.8c.6.5 1 1.3 1 2.2s-.3 1.8-.8 2.3c-.7.7-1.7 1-2.9 1h-.7v3.5h-2V5.3zm2 4.4c.2 0 .4.1.7.1c1.1 0 1.8-.6 1.8-1.6c0-.9-.6-1.4-1.6-1.4c-.4 0-.7 0-.8.1v2.8zm14.4.2c0 3.2-1.8 5.1-4.4 5.1c-2.7 0-4.2-2.2-4.2-4.9c0-2.9 1.7-5.1 4.4-5.1s4.2 2.2 4.2 4.9m-6.5.1c0 1.9.8 3.2 2.2 3.2c1.4 0 2.2-1.4 2.2-3.3c0-1.7-.8-3.2-2.2-3.2c-1.4 0-2.2 1.4-2.2 3.3m8.3 2.6c.6.3 1.4.6 2.3.6c.9 0 1.4-.4 1.4-1s-.4-.9-1.5-1.4c-1.5-.6-2.5-1.5-2.5-2.9c0-1.7 1.3-2.9 3.4-2.9c1 0 1.8.2 2.3.5l-.5 1.8c-.4-.2-1-.5-1.9-.5c-.9 0-1.3.4-1.3.9c0 .6.5.9 1.7 1.4c1.6.6 2.4 1.5 2.4 2.9c0 1.6-1.2 3-3.6 3c-1 0-2.1-.3-2.6-.6l.4-1.8M55.5 7h-2.4V5.2H60V7h-2.5v7.8h-2V7z"/>`),
		g.Group(children),
	)
}