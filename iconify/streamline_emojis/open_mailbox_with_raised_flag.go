package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenMailboxWithRaisedFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00b8f0" d="M14.53 5.24h18.33a8.53 8.53 0 0 1 8.53 8.53v15.1a1.25 1.25 0 0 1-1.25 1.25H7.25A1.25 1.25 0 0 1 6 28.87v-15.1a8.53 8.53 0 0 1 8.53-8.53Z"/><path fill="#4acfff" d="M32.86 5.24H14.53A8.53 8.53 0 0 0 6 13.77v3.75A8.53 8.53 0 0 1 14.53 9h18.33a8.53 8.53 0 0 1 8.53 8.53v-3.76a8.53 8.53 0 0 0-8.53-8.53Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.53 5.24h18.33a8.53 8.53 0 0 1 8.53 8.53v15.1a1.25 1.25 0 0 1-1.25 1.25H7.25A1.25 1.25 0 0 1 6 28.87v-15.1a8.53 8.53 0 0 1 8.53-8.53Z"/><path fill="#45413c" d="M11.08 44.13a14.37 1.87 0 1 0 28.74 0a14.37 1.87 0 1 0-28.74 0Z" opacity=".15"/><path fill="#c0dceb" d="M14.54 5.24A8.54 8.54 0 0 0 6 13.78v15.09a1.25 1.25 0 0 0 1.25 1.25h14.57a1.25 1.25 0 0 0 1.25-1.25V13.78a8.53 8.53 0 0 0-8.53-8.54Z"/><path fill="#8ca4b8" d="M14.54 5.24a8.49 8.49 0 0 0-5.65 2.14a8.52 8.52 0 0 1 9.88 8.43v14.31h3.05a1.25 1.25 0 0 0 1.25-1.25V13.78a8.53 8.53 0 0 0-8.53-8.54Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M14.54 5.24A8.54 8.54 0 0 0 6 13.78v15.09a1.25 1.25 0 0 0 1.25 1.25h14.57a1.25 1.25 0 0 0 1.25-1.25V13.78a8.53 8.53 0 0 0-8.53-8.54Z"/><path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31 1.37h-3.25a1.24 1.24 0 0 0-1.25 1.24V18.3a1.88 1.88 0 1 0 3.75 0V8.64H31a1.24 1.24 0 0 0 1.28-1.24V2.61A1.24 1.24 0 0 0 31 1.37Z"/><path fill="#debb7e" d="M22.16 30.12h5.94v12.63A1.25 1.25 0 0 1 26.85 44h-3.44a1.25 1.25 0 0 1-1.25-1.25V30.12Z"/><path fill="#b89558" d="M22.16 30.12h5.93v3.12h-5.93z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M22.16 30.12h5.94v12.63A1.25 1.25 0 0 1 26.85 44h-3.44a1.25 1.25 0 0 1-1.25-1.25V30.12h0Z"/><path fill="#fff" d="M11.17 13.35a1.24 1.24 0 0 0-1.24 1.25v14.27a1.24 1.24 0 0 0 1.24 1.25h10.65a1.25 1.25 0 0 0 1.25-1.25V13.35Z"/><path fill="#f0f0f0" d="M23.06 13.35h-3.38v15.52a1.25 1.25 0 0 1-1.25 1.25h3.38a1.25 1.25 0 0 0 1.25-1.25V13.78c.01-.15.01-.29 0-.43Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M11.17 13.35a1.24 1.24 0 0 0-1.24 1.25v14.27a1.24 1.24 0 0 0 1.24 1.25h10.65a1.25 1.25 0 0 0 1.25-1.25V13.35Zm-.93.43l12.83 9.52m-12.83 6.39l10.01-8.48"/>`),
		g.Group(children),
	)
}