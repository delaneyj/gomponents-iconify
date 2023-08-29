package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DripCoffeeMaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M16.617 37.555A3 3 0 0 1 19.584 35h13.832a3 3 0 0 1 2.967 2.555l2.1 14A3 3 0 0 1 35.516 55H17.484a3 3 0 0 1-2.967-3.445l2.1-14Z"/><path fill="#fff" fill-rule="evenodd" d="M20.744 35h-1.16a3 3 0 0 0-2.967 2.555l-2.1 14a3 3 0 0 0 .016.993A4.5 4.5 0 0 1 12 48.5v-12a4.5 4.5 0 0 1 8.744-1.5Z" clip-rule="evenodd"/><path fill="#6A462F" fill-rule="evenodd" d="M18 55h17a3 3 0 0 0 3-3v-7H15v7a3 3 0 0 0 3 3Z" clip-rule="evenodd"/><path fill="#fff" fill-rule="evenodd" d="M13 12a1 1 0 0 0-1 1v13h31v29H12v4a1 1 0 0 0 1 1h46a1 1 0 0 0 1-1V13a1 1 0 0 0-1-1H13Z" clip-rule="evenodd"/><rect width="4" height="20" x="49" y="19" fill="#D0CFCE" rx="1"/><circle cx="51.5" cy="46.5" r="2.5" fill="#D22F27" stroke="#D0CFCE" stroke-width="2"/><path fill="#fff" d="M17 29h19v6H17z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M19.584 35a3 3 0 0 0-2.967 2.555l-2.1 14a3 3 0 0 0 .016.993A4.5 4.5 0 0 1 12 48.5v-12a4.5 4.5 0 0 1 4.5-4.5"/><path d="M12 26h-1a1 1 0 0 0 1 1v-1Zm31 0h1a1 1 0 0 0-1-1v1Zm0 29v1a1 1 0 0 0 1-1h-1Zm-31 0v-1a1 1 0 0 0-1 1h1Zm1-42v-2a2 2 0 0 0-2 2h2Zm0 13V13h-2v13h2Zm-1 1h31v-2H12v2Zm30-1v29h2V26h-2Zm1 28H12v2h31v-2Zm-30 5v-4h-2v4h2Zm0 0h-2a2 2 0 0 0 2 2v-2Zm46 0H13v2h46v-2Zm0 0v2a2 2 0 0 0 2-2h-2Zm0-46v46h2V13h-2Zm0 0h2a2 2 0 0 0-2-2v2Zm-46 0h46v-2H13v2Z"/><path fill="none" stroke="#000" stroke-width="2" d="M12 26h30.5M40 26V12.5"/><path stroke="#000" stroke-width="2" d="M12 55h48"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M51.5 49a2.5 2.5 0 0 1 0-5"/><path fill="none" stroke="#000" stroke-width="2" d="M16.617 37.555A3 3 0 0 1 19.584 35h13.832a3 3 0 0 1 2.967 2.555l2.1 14A3 3 0 0 1 35.516 55H17.484a3 3 0 0 1-2.967-3.445l2.1-14Z"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M17 29h19v6H17z"/><path stroke="#000" stroke-linecap="round" stroke-width="2" d="M50 20v18"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M50 20h2m-2 18h2m-2-9h2"/>`),
		g.Group(children),
	)
}