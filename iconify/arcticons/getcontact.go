package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Getcontact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.62 20.285a9.232 9.232 0 0 0 .34-1.696a1.136 1.136 0 0 1 1.137-.999h2.673c.688 0 1.217.594 1.157 1.278c-.61 6.897-6.104 12.391-13 13a1.165 1.165 0 0 1-1.279-1.154v-2.383c0-.87.426-1.368 1-1.428a9.317 9.317 0 0 0 1.695-.341a1.897 1.897 0 0 1 1.884.481l1.144 1.144a12.098 12.098 0 0 0 4.874-4.874L27.1 22.169a1.897 1.897 0 0 1-.481-1.884h0Zm-4.469-2.949a12.098 12.098 0 0 0-4.873 4.874"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.33 34.359c6.227-6.038 6.227-14.68 0-20.718a372.725 372.725 0 0 1-3.97-3.971c-6.037-6.227-14.68-6.227-20.718 0a372.725 372.725 0 0 1-3.972 3.97c-6.227 6.038-6.227 14.68 0 20.718a371.68 371.68 0 0 1 3.971 3.971c6.037 6.227 14.68 6.227 20.718 0a371.68 371.68 0 0 1 3.971-3.971Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.684 31.442a12.909 12.909 0 0 0 2.344-7.446c0-7.19-5.827-13.028-13.028-13.028s-13.028 5.837-13.028 13.028c0 7.2 5.827 13.039 13.028 13.039c2.78 0 5.347-.863 7.457-2.344h-.001l7.228 7.228a2.288 2.288 0 0 0 3.235 0h0a2.288 2.288 0 0 0 0-3.235l-7.238-7.239l.003-.003Z"/>`),
		g.Group(children),
	)
}