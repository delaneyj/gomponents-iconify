package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Memo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiMemo0" d="m38.07 47.191l2.386 2.386l-3.464 1.28z"/></defs><path fill="#FFF" d="M16.405 11.378H55.97v49.066H16.405z"/><path fill="#f4aa41" d="m39.931 40.222l11.294-11.294l7.376 7.376l-11.198 11.198"/><path fill="#a57939" d="m37.941 46.819l1.961-5.649l6.454 6.454l-5.648 1.961"/><path fill="#EA5A47" d="m54.609 25.052l3.978-3.979l7.859 7.859l-3.945 3.945"/><path fill="#d0cfce" d="m50.494 29.659l4.608-4.607l7.375 7.375l-4.568 4.569"/><use href="#openmojiMemo0"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M22.322 19.875h27m-27 8h25.6m-25.6 8h17.109m-17.109 8h12.206m-12.206 8h10.146"/><use href="#openmojiMemo0"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m40.692 49.591l5.664-1.967l15.59-15.59l-6.454-6.454l-15.59 15.59l-1.974 5.671zm17.394-26.605l1.555-1.555l6.454 6.454l-1.632 1.632M40.279 40.793l6.454 6.454"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m37.932 46.828l-1.383 4.149l4.159-1.392m10.594-19.428l6.202 6.202"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2.021" d="M55.322 44.228v15.647h-39v-48h39v8.456"/>`),
		g.Group(children),
	)
}