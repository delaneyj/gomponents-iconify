package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfficeBuilding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#3e4347" d="M26 6h20v48H26z"/><path fill="#dae3ea" d="M16 2h32v4H16z"/><path fill="#3e4347" d="M4 6h10v46H4z"/><path fill="#b2c1c0" d="M26 52h28v8H26z"/><path fill="#b4d7ee" d="M30 52h4v8h-4zm8 0h4v8h-4zm8 0h4v8h-4zm0-10h10v4H46zm0-10h10v4H46zm0-10h10v4H46zm0-10h10v4H46zm-8 30h4v4h-4zm0-10h4v4h-4zm0-10h4v4h-4zm0-10h4v4h-4zm-8 30h4v4h-4zm0-10h4v4h-4zm0-10h4v4h-4zm0-10h4v4h-4z"/><path fill="#3e4347" d="M26 46h32v6H26zm0-10h32v6H26zm0-10h32v6H26zm0-10h32v6H26zm0-10h32v6H26z"/><path fill="#b4d7ee" d="M14 8h12v44H14zM4 52h22v10H4z"/><path fill="#dae3ea" d="M2 50h24v2H2z"/><path fill="#d6eef0" d="M14 38h12v2H14zm0-10h12v2H14zm0-10h12v2H14zm0-8h12v2H14z"/><path fill="#f5f5f5" d="M14 6h12v2H14z"/><path fill="#b2c1c0" d="M4 52h2v10H4z"/><g fill="#f5f5f5"><path d="M10 52h1.3v10H10z"/><path d="M10 52h12v1.1H10zm0 8.9h12V62H10z"/><path d="M15.3 52h1.3v10h-1.3zm5.4 0H22v10h-1.3z"/></g><path fill="#83bf4f" d="M26 58h36v2H26z"/><path fill="#dae3ea" d="M24 60h38v2H24z"/><path fill="#83bf4f" d="M61 57.8c.2-.3.3-.7.3-1.2c0-1.3-1-2.3-2.3-2.3h-.1V54c0-.6-.5-1.2-1.2-1.2c-.3 0-.5.1-.7.3c-.1-.8-.9-1.4-1.7-1.4c-1 0-1.7.8-1.7 1.8v.4c-.3-.2-.7-.4-1.1-.4c-.7 0-1.3.4-1.6 1c-.2-.1-.5-.1-.8-.1c-1.3 0-2.3 1-2.3 2.3c0 .5.2 1 .5 1.4c-.3.3-.5.7-.5 1.2c0 1 .8 1.8 1.7 1.8h9.2c.3.3.8.5 1.2.5c1.1 0 2-.9 2-2c.1-.8-.3-1.4-.9-1.8"/><path fill="#699635" d="M51.7 59.6c0 .6.5-.3 1.2-.3c.6 0 1.2.9 1.2.3c0-.6-.5-1.2-1.2-1.2s-1.2.6-1.2 1.2m3.4-2c0 .3.3 0 .6 0s.6.3.6 0s-.3-.6-.6-.6s-.6.2-.6.6m4.3 2.1c0 .3.3 0 .6 0s.6.3.6 0s-.3-.6-.6-.6c-.3.1-.6.3-.6.6"/><path fill="#b5f478" d="M56.9 56c0 .3.3-.1.6-.1s.6.4.6.1c0-.3-.3-.6-.6-.6s-.6.2-.6.6m2 1.7c0 .3.3-.1.6-.1s.6.4.6.1c0-.3-.3-.6-.6-.6c-.4 0-.6.2-.6.6"/><path fill="#699635" d="M49.2 56.6c0 .7.5-.4 1.2-.4s1.2 1 1.2.4s-.5-1.2-1.2-1.2s-1.2.5-1.2 1.2m-.1 2.7c0 .3.3 0 .6 0s.6.3.6 0s-.3-.6-.6-.6c-.3.1-.6.3-.6.6"/><path fill="#b5f478" d="M54.3 54.7c0 .4.3-.2.8-.2c.4 0 .8.7.8.2c0-.4-.3-.8-.8-.8c-.4 0-.8.3-.8.8"/><path fill="#699635" d="M56.5 59.9c0 .4.3-.2.8-.2c.4 0 .8.7.8.2c0-.4-.3-.8-.8-.8s-.8.3-.8.8"/>`),
		g.Group(children),
	)
}