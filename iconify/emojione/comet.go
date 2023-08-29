package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#62bbc7"><path d="M25.5 8.3L4.9 48.4v10.7h10.7z"/><path d="M55.7 38.5L4.9 48.4v10.7h10.7z"/></g><g fill="#6dd1de"><path d="M39.7 8.3L4.9 48.4v10.7h10.7z"/><path d="M55.7 24.3L15.6 59.1H4.9V48.4z"/></g><path fill="#7df0ff" d="M62 2L4.9 48.4v10.7h10.7z"/><path fill="#c2f8ff" d="M33.4 30.6L4.9 48.4v10.7h10.7z"/><circle cx="9.6" cy="54.4" r="7.6" fill="#7df0ff"/><circle cx="9.6" cy="54.4" r="6.7" fill="#5995c1"/><circle cx="9.4" cy="49.9" r="1.5" fill="#4b7ea3"/><circle cx="9.3" cy="49.6" r="1.2" fill="#5995c1"/><circle cx="14.2" cy="52.7" r="1" fill="#4b7ea3"/><circle cx="14.1" cy="52.5" r=".8" fill="#5995c1"/><circle cx="11.7" cy="58.8" r="1.5" fill="#4b7ea3"/><circle cx="11.7" cy="58.5" r="1.2" fill="#5995c1"/><circle cx="4.7" cy="52.9" r="1" fill="#4b7ea3"/><circle cx="4.6" cy="52.7" r=".8" fill="#5995c1"/><circle cx="7.9" cy="56.2" r="2.5" fill="#4b7ea3"/><circle cx="7.7" cy="55.7" r="2" fill="#5995c1"/><path fill="#c2f8ff" d="m10.839 32.64l2.05-2.05l2.051 2.05l-2.05 2.05zm18.538-8.42l2.05-2.05l2.05 2.049l-2.049 2.051zm-1.452-11.259l2.05-2.051l2.051 2.05l-2.05 2.051zm22.381 7.152l2.051-2.05l2.05 2.05l-2.05 2.051zM44.1 36.943l2.05-2.05l2.05 2.05l-2.05 2.05zm-23.874-6.982l.99-.99l.99.99l-.99.99zm-2.076-5.646l.989-.99l.99.99l-.99.99zm16.727 14.192l.99-.99l.99.991l-.99.99zm-3.841 8.887l.99-.99l.99.99l-.99.99zm9.286-4.068l.991-.99l.99.991l-.991.99zm5.816-18.399l.99-.99l.99.991l-.99.99zm7.208 5.085l.99-.99l.99.99l-.99.99zM42.317 11.024l.99-.99l.99.99l-.99.99zm2.786-7.384l.99-.99l.99.99l-.99.99zM33.896 8.297l.99-.99l.99.99l-.99.99z"/>`),
		g.Group(children),
	)
}