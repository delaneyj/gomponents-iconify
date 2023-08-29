package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FramedPicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiFramedPicture0" d="M18 18h36v36H18z"/></defs><g stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path fill="#A57939" d="M12 12h48v48H12z"/><path fill="#92D3F5" d="M18 18h36v36H18z"/><circle cx="26" cy="30" r="4" fill="#FCEA2B"/><path fill="#5C9E31" stroke="#5C9E31" d="M50 35c-2.896-.854-6.28-7.989-8-8c-4.208-.028-6.254 5.836-11 9c-3 2-3.374 2.85-6 4c-2.282 1-3 3-3.25 3.64c-.303.777-1.475 2.522-.958 3.11c1.208 1.375 2.458 1.5 5 .718c2.147-.66 4.977-4.808 6.875-6.218c2.27-1.688 4.645-2.5 8.083-2c2.48.36 6.66 3.172 7.813 3.063c1.312-.126-1.594-2.563-.532-4.188c1.133-1.732 2.91.153 3.636-1.083C51.984 36.5 50.632 35.186 50 35z"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M12 12h48v48H12z"/><use href="#openmojiFramedPicture0"/><circle cx="26" cy="30" r="4"/><use href="#openmojiFramedPicture0"/><path d="M22 43c.526-1.02.728-1.967 3-3c2.61-1.186 3-2 6-4c4.746-3.164 6.792-9.028 11-9c1.72.011 5 7 8 8"/></g>`),
		g.Group(children),
	)
}