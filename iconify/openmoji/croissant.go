package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Croissant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiCroissant0" d="m59.107 36.112l-8.018 4.769A4.993 4.993 0 0 0 56 45a5 5 0 0 0 5-5c0-1.58-.748-2.972-1.893-3.888zM42 17a2 2 0 0 1 2 2l-3 15a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2l-3-15a2 2 0 0 1 2-2h12z"/><path id="openmojiCroissant1" d="m53.496 23.576l-9.205-3.906a2.073 2.073 0 0 0-.4-.122L41 34c0 .882-.574 1.622-1.367 1.888c.201.254.459.467.777.602l3.682 1.563a2 2 0 0 0 2.623-1.06L54.557 26.2a2 2 0 0 0-1.06-2.623z"/><path id="openmojiCroissant2" d="m60.412 32.478l-5.756-5.556a1.926 1.926 0 0 0-.409-.297l-7.532 10.368a1.98 1.98 0 0 1-.887.967c-.13.646.057 1.344.567 1.835l1.439 1.39a2 2 0 0 0 2.828-.05l9.799-5.828a2 2 0 0 0-.049-2.829zm-47.519 3.634l8.018 4.769A4.993 4.993 0 0 1 16 45a5 5 0 0 1-5-5c0-1.58.748-2.972 1.893-3.888zm5.611-12.536l9.205-3.906c.13-.056.266-.095.4-.122L31 34c0 .882.574 1.622 1.367 1.888a1.982 1.982 0 0 1-.777.602l-3.682 1.563a2 2 0 0 1-2.623-1.06L17.443 26.2a2 2 0 0 1 1.06-2.623z"/><path id="openmojiCroissant3" d="m11.588 32.478l5.756-5.556c.125-.121.263-.216.409-.297l7.532 10.368c.183.432.504.758.887.967a1.992 1.992 0 0 1-.567 1.835l-1.439 1.39a2 2 0 0 1-2.828-.05l-9.799-5.828a2 2 0 0 1 .049-2.829z"/></defs><g fill="#f4aa41"><use href="#openmojiCroissant0"/><use href="#openmojiCroissant1"/><use href="#openmojiCroissant2"/><use href="#openmojiCroissant3"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiCroissant0"/><use href="#openmojiCroissant1"/><use href="#openmojiCroissant2"/><use href="#openmojiCroissant3"/></g>`),
		g.Group(children),
	)
}