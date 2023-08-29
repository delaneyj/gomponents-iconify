package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KwgtPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="5.978" height="28.394" x="13.933" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.196" ry="1.196"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.469 4.5h-.598a4.782 4.782 0 0 0-4.782 4.782v3.173a.598.598 0 0 1-.196.442l-5.407 4.915a1.196 1.196 0 0 0 0 1.77l5.407 4.915a.598.598 0 0 1 .196.442v3.173a4.782 4.782 0 0 0 4.782 4.782h.598c.33 0 .598-.267.598-.597v-9.518c0-.317-.126-.621-.35-.845l-2.814-2.814a.598.598 0 0 1 0-.846l2.813-2.813c.225-.225.35-.529.35-.846V5.098a.598.598 0 0 0-.597-.598Z"/><circle cx="16.922" cy="40.511" r="2.989" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.911 40.511h14.156V43.5m-4.817-1.02v-1.969"/>`),
		g.Group(children),
	)
}