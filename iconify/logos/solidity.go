package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Solidity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="m191.513 0l-63.867 113.512H0L63.823 0h127.69" opacity=".45"/><path d="M127.646 113.512h127.691L191.513 0H63.823z" opacity=".6"/><path d="m63.823 226.981l63.823-113.469L63.823 0L0 113.512z" opacity=".8"/><path d="m64.442 397.25l63.867-113.513H256L192.132 397.25H64.442" opacity=".45"/><path d="M128.309 283.737H.618L64.441 397.25h127.691z" opacity=".6"/><path d="m192.132 170.269l-63.823 113.468l63.823 113.513L256 283.737z" opacity=".8"/>`),
		g.Group(children),
	)
}