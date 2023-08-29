package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCombiner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.5 12.374c.771-.097 1.543.482 1.543 1.254h0l2.509 23.831c.096.772-.483 1.544-1.255 1.544L8.563 41.512c-.772.096-1.544-.483-1.544-1.255h0L4.511 16.523c-.097-.772.482-1.544 1.254-1.544L29.5 12.374Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.905 18.935c-1.737.193-3.088 1.736-2.895 3.57c.193 1.736 1.737 3.087 3.57 2.894c1.737-.193 3.087-1.737 2.894-3.57s-1.833-3.087-3.57-2.894h0Zm8.78 5.596l-4.632 5.692c-.193.193-.482.29-.675.097h0l-1.061-.869c-.193-.193-.58-.096-.676.097l-5.017 6.175c-.193.193-.193.579.097.675c-.097.097.096.097.193.097l19.971-2.123c.29 0 .483-.29.483-.579c0-.096-.097-.193-.097-.29l-7.815-9.069c-.29-.096-.579-.096-.772.097h0Zm-5.404-10.71l.676-6.078c.096-.772.772-1.351 1.544-1.255l23.734 2.509c.772.096 1.351.772 1.255 1.544h0l-2.51 23.734c0 .676-.772 1.255-1.544 1.255l-6.175-.676"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.78 29.934l5.306.579c.29 0 .483-.193.58-.483c0-.096 0-.193-.097-.29l-5.79-10.516c-.096-.193-.482-.29-.675-.193h0l-.386.386"/>`),
		g.Group(children),
	)
}