package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yourwallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 14.502a2.805 2.805 0 0 0-.01 2.847l.01.02l1.921 3.32h-7.685a2.814 2.814 0 0 1-2.453-1.438a2.72 2.72 0 0 1-.365-1.38V6.334a2.817 2.817 0 0 1 2.818-2.818h11.528a2.817 2.817 0 0 1 2.818 2.818v2.798l-7.567 4.365A2.835 2.835 0 0 0 24 14.502Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.079 20.69L20.177 24l-1.911 3.31l-.02.03a2.8 2.8 0 0 1-2.463 1.41a2.95 2.95 0 0 1-1.39-.385L6.837 24l-2.424-1.4a2.82 2.82 0 0 1-1.034-3.852l5.764-9.98a2.82 2.82 0 0 1 3.852-1.035l2.424 1.399v8.74c0 .502.128.975.365 1.379a2.814 2.814 0 0 0 2.453 1.439h3.843Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 33.498a2.837 2.837 0 0 1-1.015 1.005l-7.567 4.365l-2.424 1.4a2.82 2.82 0 0 1-3.852-1.035l-5.764-9.981a2.82 2.82 0 0 1 1.034-3.853L6.836 24l7.557 4.365a2.95 2.95 0 0 0 1.39.384a2.8 2.8 0 0 0 2.463-1.409l.02-.03l1.91-3.31l1.903 3.31l1.92 3.321l.01.02c.533.916.493 2-.01 2.847Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.582 30.128v11.538a2.817 2.817 0 0 1-2.818 2.818H18.236a2.817 2.817 0 0 1-2.818-2.818v-2.798l7.567-4.365A2.835 2.835 0 0 0 24 33.498a2.805 2.805 0 0 0 .01-2.847l-.01-.02l-1.921-3.32h7.685c1.054 0 1.97.581 2.453 1.438c.237.404.365.877.365 1.38Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m44.622 29.252l-5.764 9.98a2.82 2.82 0 0 1-3.852 1.035l-2.424-1.399v-8.74a2.72 2.72 0 0 0-.365-1.379a2.814 2.814 0 0 0-2.453-1.438H25.92L27.823 24l1.911-3.31l.02-.03a2.8 2.8 0 0 1 2.463-1.41c.473.01.956.139 1.39.385L41.163 24l2.424 1.4a2.82 2.82 0 0 1 1.034 3.852Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.588 22.6L41.164 24l-7.557-4.365a2.95 2.95 0 0 0-1.39-.384a2.8 2.8 0 0 0-2.463 1.409l-.02.03L27.823 24l-1.902-3.31l-1.92-3.321l-.01-.02c-.533-.916-.493-2 .01-2.847a2.837 2.837 0 0 1 1.014-1.005l7.567-4.365l2.424-1.4a2.82 2.82 0 0 1 3.852 1.035l5.764 9.981a2.82 2.82 0 0 1-1.034 3.853Z"/>`),
		g.Group(children),
	)
}