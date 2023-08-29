package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Greece(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#009fd9" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M38.68 32.79A45.62 45.62 0 0 1 24 30.41A46 46 0 0 0 9.37 28H2v3.49h7.37A46 46 0 0 1 24 33.9a45.62 45.62 0 0 0 14.65 2.38h7.37v-3.49ZM9.93 7.52H3c-.58 0-1 .35-1 .79v6.06h7.93Zm0 10.33H2v6.85h7.93Zm28.75 1.25A45.62 45.62 0 0 1 24 16.72c-.79-.26-1.59-.5-2.39-.72v3.49c.8.22 1.6.45 2.39.72a45.62 45.62 0 0 0 14.65 2.38h7.37V19.1ZM45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45 45 0 0 0 13.71 7.7v6.87A46.36 46.36 0 0 1 21.64 16v-3.36c.8.22 1.6.46 2.39.72a45.62 45.62 0 0 0 14.65 2.38h7.37V13a.94.94 0 0 0-1.05-.75Z"/><path fill="#009fd9" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 27.06a45.89 45.89 0 0 0 14.65 2.37h7.37v-3.49h-7.34A45.89 45.89 0 0 1 24 23.57c-.79-.27-1.59-.5-2.39-.73v-3.35a45.41 45.41 0 0 0-7.93-1.43v6.85A44.93 44.93 0 0 1 24 27.06Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}