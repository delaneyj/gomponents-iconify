package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accordion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.838 35.7c-1.46.53-2.639-.437-3.188-2.18L4.894 21.592a2.969 2.969 0 0 1 1.229-3.484l2.606-1.176l.934-.412c1.086-.48 1.802-.21 2.184 1.012l4.562 14.619c.383 1.225.275 1.795-.718 2.154Zm4.572-3.55c5.683-2.449 8.606-2.893 14.936-.893m-19.4-13.807l.92-2.94l2.254 1.978l2.275 8.384"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.73 16.932l.707 2.23l-2.11.747"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.437 19.162l.707 2.23l-2.11.747m2.11-.747l.707 2.23l-2.11.747"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.851 23.621l.708 2.23l-2.11.748"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.559 25.851l.707 2.23l-2.11.748m2.11-.748l.707 2.23l-2.11.748m2.11-.748l.708 2.23l-2.11.748m2.109-.748l.708 2.23m.732-18.283l1.158-2.844l2.083 2.178l1.575 8.562"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.36 15.822l1.39-2.727l1.894 2.362l.863 8.675"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.644 15.457l1.608-2.588l1.693 2.528l.144 8.72m-.144-8.721l1.815-2.43l1.478 2.674l-.576 8.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.238 15.64l2.01-2.254l1.251 2.8l-1.291 8.614m1.292-8.614l2.186-2.06l1.016 2.904M31.05 32.507l.295-1.25l3.357-14.227l.375-1.583"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.077 15.447l6.903 1.822a2.485 2.485 0 0 1 1.69 3.074L40.798 32.51a2.38 2.38 0 0 1-3.196 1.727l-6.552-1.73"/><circle cx="37.225" cy="18.926" r=".75" fill="currentColor"/><circle cx="36.235" cy="22.798" r=".75" fill="currentColor"/><circle cx="35.386" cy="26.824" r=".75" fill="currentColor"/><circle cx="34.316" cy="30.756" r=".75" fill="currentColor"/><circle cx="38.217" cy="29.844" r=".75" fill="currentColor"/><circle cx="39.255" cy="25.66" r=".75" fill="currentColor"/><circle cx="40.167" cy="21.508" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}