package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KanjiStudy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.193 23.74v-3.02h13.511v3.115m-6.755-5.759v2.644m-4.11 3.304h8.031l-3.78 2.927v6.797h-2.74m-4.157-5.193h13.511"/><rect width="19.593" height="28.186" x="23.907" y="12.458" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.05"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.288 7.417a1.23 1.23 0 0 0-.478.098L14.03 14.22a1.225 1.225 0 0 0-.648 1.609l10.047 23.689a1.218 1.218 0 0 0 .768.687a1.22 1.22 0 0 1-.29-.785V13.683a1.222 1.222 0 0 1 1.225-1.225h8.105l-1.82-4.292a1.224 1.224 0 0 0-1.129-.749Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.847 7.356a1.219 1.219 0 0 0-.866.359L4.86 19.848a1.223 1.223 0 0 0 0 1.733l18.181 18.198a1.214 1.214 0 0 0 .967.35a1.216 1.216 0 0 1-.58-.61L13.382 15.826a1.222 1.222 0 0 1 .649-1.606l7.851-3.335l-3.169-3.171a1.219 1.219 0 0 0-.865-.36Z"/>`),
		g.Group(children),
	)
}