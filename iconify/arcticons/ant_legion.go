package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntLegion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.892 4.5a9.048 9.048 0 0 1-2.473 5.293l3.558 4.555"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.372 16.024c-1.8-.705-5.931-4.296-7.26 2.141c-2.409 1.271-4.861 2.496-4.903 6.247c-1.021.134-2.642-.14-.91 1.865l-1.39.13a5.016 5.016 0 0 0-4.598-1.995c-1.335.212-2.29 1.561-3.036 2.69a5.632 5.632 0 0 0-1.041 3.123a7.28 7.28 0 0 0 1.258 3.254c.623 1.056 1.037 2.012 2.56.911a46.44 46.44 0 0 0 4.468-4.165c5.233.661 5.007 9.472 7.505 13.275"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.496 14.408a10.363 10.363 0 0 0 3.08-5.093c3.55 1.197 8.556-.325 11.191 1.389m-10.395 5.32c-2.151 1.34-3.62 2.753-6.003 3.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.37 19.423l4.25 1.215l3.298-1.301l2.429.303m-9.977-.217l-.738 4.599m.408-2.547l3.279 3.501l2.863-2.646l3.514.911m-9.864-.468c-2.798 1.028-1.1 2.655-.677 4.112c-3.365.316-5.218 1.726-5.228 4.471m5.228-4.471c4.253-.055 5.717.503 7.158 1.605m-7.158-1.605a12.813 12.813 0 0 1 2.134 3.928a33.229 33.229 0 0 0 .902 4.618"/>`),
		g.Group(children),
	)
}