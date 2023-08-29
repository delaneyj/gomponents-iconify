package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhysicalTherapyNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPhysicalTherapyNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM6 9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Zm25 4.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0ZM17.313 29h2.935l3.864 3.813l1.92 4.427a2 2 0 1 0 3.67-1.591l-2.075-4.784a2 2 0 0 0-.43-.627L25.943 29H34a2 2 0 0 1 2 2v7h2v-7a4 4 0 0 0-4-4h-3.992a2 2 0 0 0 .038-4c-1.625-.037-2.08-.307-2.397-.593c-.225-.202-.482-.516-.846-1.074c-.261-.402-.53-.853-.863-1.41c-.14-.237-.294-.493-.462-.772l-.04-.066a8.402 8.402 0 0 0-.423-.657a2.735 2.735 0 0 0-1.142-.923c-.6-.25-1.156-.206-1.493-.153c-.319.05-.66.148-.932.226c-.847.243-1.543.882-1.99 1.367A9.575 9.575 0 0 0 18.039 21c-.76 1.47-1.435 3.716-.483 5.823c.027.061.058.12.09.177H14a4 4 0 0 0-4 4v7h2v-7a2 2 0 0 1 2-2h3.16l-.126.5l-1.025 3.46l-1.508 2.43a2 2 0 0 0 3.399 2.109l1.65-2.66a2 2 0 0 0 .217-.486l.693-2.335L17.313 29Zm12.681-2h-6.078l-1.183-1.167l1.247-1.557c.295.393.62.768.993 1.104c1.346 1.212 2.937 1.572 4.98 1.62a2.431 2.431 0 0 0 .041 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPhysicalTherapyNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}