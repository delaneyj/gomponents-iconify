package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagLebanon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 17h62v7H5zm0 31h62v7H5z"/><path fill="#5c9e31" stroke="#5c9e31" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m36 27l-7 5l7-1l7 1l-7-5z"/><path fill="#5c9e31" stroke="#5c9e31" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m36 31l-8 5l8-1l8 1l-8-5z"/><path fill="#5c9e31" stroke="#5c9e31" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m36 35l-9 5l9-1l9 1l-9-5z"/><path fill="#5c9e31" d="M38 39v3.5a2.227 2.227 0 0 0 1.423 1.974l.154.052l-.154-.052A11.113 11.113 0 0 0 36.5 44h-1a11.113 11.113 0 0 0-2.923.474l-.154.052l.154-.052A2.227 2.227 0 0 0 34 42.5V39"/><path fill="none" stroke="#5c9e31" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M38 39v3.5a2.227 2.227 0 0 0 1.423 1.974l.154.052h0l-.154-.052A11.113 11.113 0 0 0 36.5 44h-1a11.113 11.113 0 0 0-2.923.474l-.154.052h0l.154-.052A2.227 2.227 0 0 0 34 42.5V39"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}