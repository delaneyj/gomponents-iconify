package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const IconifyVersion = ""

func IconFromName(name string) g.Node {
	switch name {
	case "account":
		return Account()
	case "accountBox":
		return AccountBox()
	case "alert":
		return Alert()
	case "alertBox":
		return AlertBox()
	case "alertBoxFill":
		return AlertBoxFill()
	case "alertCircle":
		return AlertCircle()
	case "alertCircleFill":
		return AlertCircleFill()
	case "alertRhombus":
		return AlertRhombus()
	case "alertRhombusFill":
		return AlertRhombusFill()
	case "alphaA":
		return AlphaA()
	case "alphaAFill":
		return AlphaAFill()
	case "alphaB":
		return AlphaB()
	case "alphaBFill":
		return AlphaBFill()
	case "alphaC":
		return AlphaC()
	case "alphaCFill":
		return AlphaCFill()
	case "alphaD":
		return AlphaD()
	case "alphaDFill":
		return AlphaDFill()
	case "alphaE":
		return AlphaE()
	case "alphaEFill":
		return AlphaEFill()
	case "alphaF":
		return AlphaF()
	case "alphaFFill":
		return AlphaFFill()
	case "alphaG":
		return AlphaG()
	case "alphaGFill":
		return AlphaGFill()
	case "alphaH":
		return AlphaH()
	case "alphaHFill":
		return AlphaHFill()
	case "alphaI":
		return AlphaI()
	case "alphaIFill":
		return AlphaIFill()
	case "alphaJ":
		return AlphaJ()
	case "alphaJFill":
		return AlphaJFill()
	case "alphaK":
		return AlphaK()
	case "alphaKFill":
		return AlphaKFill()
	case "alphaL":
		return AlphaL()
	case "alphaLFill":
		return AlphaLFill()
	case "alphaM":
		return AlphaM()
	case "alphaMFill":
		return AlphaMFill()
	case "alphaN":
		return AlphaN()
	case "alphaNFill":
		return AlphaNFill()
	case "alphaO":
		return AlphaO()
	case "alphaOFill":
		return AlphaOFill()
	case "alphaP":
		return AlphaP()
	case "alphaPFill":
		return AlphaPFill()
	case "alphaQ":
		return AlphaQ()
	case "alphaQFill":
		return AlphaQFill()
	case "alphaR":
		return AlphaR()
	case "alphaRFill":
		return AlphaRFill()
	case "alphaS":
		return AlphaS()
	case "alphaSFill":
		return AlphaSFill()
	case "alphaT":
		return AlphaT()
	case "alphaTFill":
		return AlphaTFill()
	case "alphaU":
		return AlphaU()
	case "alphaUFill":
		return AlphaUFill()
	case "alphaV":
		return AlphaV()
	case "alphaVFill":
		return AlphaVFill()
	case "alphaW":
		return AlphaW()
	case "alphaWFill":
		return AlphaWFill()
	case "alphaX":
		return AlphaX()
	case "alphaXFill":
		return AlphaXFill()
	case "alphaY":
		return AlphaY()
	case "alphaYFill":
		return AlphaYFill()
	case "alphaZ":
		return AlphaZ()
	case "alphaZFill":
		return AlphaZFill()
	case "application":
		return Application()
	case "applicationCode":
		return ApplicationCode()
	case "archive":
		return Archive()
	case "arrow":
		return Arrow()
	case "arrowBottomLeft":
		return ArrowBottomLeft()
	case "arrowBottomLeftCircle":
		return ArrowBottomLeftCircle()
	case "arrowBottomRight":
		return ArrowBottomRight()
	case "arrowBottomRightCircle":
		return ArrowBottomRightCircle()
	case "arrowDown":
		return ArrowDown()
	case "arrowDownBold":
		return ArrowDownBold()
	case "arrowDownCircle":
		return ArrowDownCircle()
	case "arrowDownLeft":
		return ArrowDownLeft()
	case "arrowDownRight":
		return ArrowDownRight()
	case "arrowLeft":
		return ArrowLeft()
	case "arrowLeftBold":
		return ArrowLeftBold()
	case "arrowLeftCircle":
		return ArrowLeftCircle()
	case "arrowLeftDown":
		return ArrowLeftDown()
	case "arrowLeftRight":
		return ArrowLeftRight()
	case "arrowLeftUp":
		return ArrowLeftUp()
	case "arrowRight":
		return ArrowRight()
	case "arrowRightBold":
		return ArrowRightBold()
	case "arrowRightCircle":
		return ArrowRightCircle()
	case "arrowRightDown":
		return ArrowRightDown()
	case "arrowRightUp":
		return ArrowRightUp()
	case "arrowTopLeft":
		return ArrowTopLeft()
	case "arrowTopLeftCircle":
		return ArrowTopLeftCircle()
	case "arrowTopRight":
		return ArrowTopRight()
	case "arrowTopRightCircle":
		return ArrowTopRightCircle()
	case "arrowUp":
		return ArrowUp()
	case "arrowUpBold":
		return ArrowUpBold()
	case "arrowUpCircle":
		return ArrowUpCircle()
	case "arrowUpDown":
		return ArrowUpDown()
	case "arrowUpLeft":
		return ArrowUpLeft()
	case "arrowUpRight":
		return ArrowUpRight()
	case "axe":
		return Axe()
	case "bagPersonal":
		return BagPersonal()
	case "bagPersonalFill":
		return BagPersonalFill()
	case "batteryFifty":
		return BatteryFifty()
	case "batteryOneHundred":
		return BatteryOneHundred()
	case "batterySeventyFive":
		return BatterySeventyFive()
	case "batteryTwentyFive":
		return BatteryTwentyFive()
	case "batteryZero":
		return BatteryZero()
	case "battleAxe":
		return BattleAxe()
	case "book":
		return Book()
	case "bookmark":
		return Bookmark()
	case "borderBottom":
		return BorderBottom()
	case "borderBottomLeft":
		return BorderBottomLeft()
	case "borderBottomLeftRight":
		return BorderBottomLeftRight()
	case "borderBottomRight":
		return BorderBottomRight()
	case "borderInside":
		return BorderInside()
	case "borderLeft":
		return BorderLeft()
	case "borderLeftRight":
		return BorderLeftRight()
	case "borderNone":
		return BorderNone()
	case "borderOutside":
		return BorderOutside()
	case "borderRight":
		return BorderRight()
	case "borderTop":
		return BorderTop()
	case "borderTopBottom":
		return BorderTopBottom()
	case "borderTopLeft":
		return BorderTopLeft()
	case "borderTopLeftBottom":
		return BorderTopLeftBottom()
	case "borderTopLeftRight":
		return BorderTopLeftRight()
	case "borderTopRight":
		return BorderTopRight()
	case "borderTopRightBottom":
		return BorderTopRightBottom()
	case "bow":
		return Bow()
	case "bowArrow":
		return BowArrow()
	case "box":
		return Box()
	case "boxLightDashedDownLeft":
		return BoxLightDashedDownLeft()
	case "boxLightDashedDownRight":
		return BoxLightDashedDownRight()
	case "boxLightDashedUpLeft":
		return BoxLightDashedUpLeft()
	case "boxLightDashedUpRight":
		return BoxLightDashedUpRight()
	case "boxLightDoubleHorizontal":
		return BoxLightDoubleHorizontal()
	case "boxLightDoubleRoundDownLeft":
		return BoxLightDoubleRoundDownLeft()
	case "boxLightDoubleRoundDownRight":
		return BoxLightDoubleRoundDownRight()
	case "boxLightDoubleRoundUpLeft":
		return BoxLightDoubleRoundUpLeft()
	case "boxLightDoubleRoundUpRight":
		return BoxLightDoubleRoundUpRight()
	case "boxLightDoubleVertical":
		return BoxLightDoubleVertical()
	case "boxLightDownLeft":
		return BoxLightDownLeft()
	case "boxLightDownLeftCircle":
		return BoxLightDownLeftCircle()
	case "boxLightDownRight":
		return BoxLightDownRight()
	case "boxLightDownRightCircle":
		return BoxLightDownRightCircle()
	case "boxLightHorizontal":
		return BoxLightHorizontal()
	case "boxLightRoundDownLeft":
		return BoxLightRoundDownLeft()
	case "boxLightRoundDownRight":
		return BoxLightRoundDownRight()
	case "boxLightRoundUpLeft":
		return BoxLightRoundUpLeft()
	case "boxLightRoundUpRight":
		return BoxLightRoundUpRight()
	case "boxLightUpLeft":
		return BoxLightUpLeft()
	case "boxLightUpLeftCircle":
		return BoxLightUpLeftCircle()
	case "boxLightUpRight":
		return BoxLightUpRight()
	case "boxLightUpRightCircle":
		return BoxLightUpRightCircle()
	case "boxLightVertical":
		return BoxLightVertical()
	case "boxOuterLightAll":
		return BoxOuterLightAll()
	case "boxOuterLightDashedAll":
		return BoxOuterLightDashedAll()
	case "boxOuterLightDashedDown":
		return BoxOuterLightDashedDown()
	case "boxOuterLightDashedDownLeft":
		return BoxOuterLightDashedDownLeft()
	case "boxOuterLightDashedDownLeftRight":
		return BoxOuterLightDashedDownLeftRight()
	case "boxOuterLightDashedDownRight":
		return BoxOuterLightDashedDownRight()
	case "boxOuterLightDashedLeft":
		return BoxOuterLightDashedLeft()
	case "boxOuterLightDashedLeftRight":
		return BoxOuterLightDashedLeftRight()
	case "boxOuterLightDashedRight":
		return BoxOuterLightDashedRight()
	case "boxOuterLightDashedUp":
		return BoxOuterLightDashedUp()
	case "boxOuterLightDashedUpDown":
		return BoxOuterLightDashedUpDown()
	case "boxOuterLightDashedUpDownLeft":
		return BoxOuterLightDashedUpDownLeft()
	case "boxOuterLightDashedUpDownRight":
		return BoxOuterLightDashedUpDownRight()
	case "boxOuterLightDashedUpLeft":
		return BoxOuterLightDashedUpLeft()
	case "boxOuterLightDashedUpLeftRight":
		return BoxOuterLightDashedUpLeftRight()
	case "boxOuterLightDashedUpRight":
		return BoxOuterLightDashedUpRight()
	case "boxOuterLightDown":
		return BoxOuterLightDown()
	case "boxOuterLightDownLeft":
		return BoxOuterLightDownLeft()
	case "boxOuterLightDownLeftRight":
		return BoxOuterLightDownLeftRight()
	case "boxOuterLightDownRight":
		return BoxOuterLightDownRight()
	case "boxOuterLightLeft":
		return BoxOuterLightLeft()
	case "boxOuterLightLeftRight":
		return BoxOuterLightLeftRight()
	case "boxOuterLightRight":
		return BoxOuterLightRight()
	case "boxOuterLightRoundDownLeft":
		return BoxOuterLightRoundDownLeft()
	case "boxOuterLightRoundDownRight":
		return BoxOuterLightRoundDownRight()
	case "boxOuterLightRoundUpLeft":
		return BoxOuterLightRoundUpLeft()
	case "boxOuterLightRoundUpRight":
		return BoxOuterLightRoundUpRight()
	case "boxOuterLightUp":
		return BoxOuterLightUp()
	case "boxOuterLightUpDown":
		return BoxOuterLightUpDown()
	case "boxOuterLightUpDownLeft":
		return BoxOuterLightUpDownLeft()
	case "boxOuterLightUpDownRight":
		return BoxOuterLightUpDownRight()
	case "boxOuterLightUpLeft":
		return BoxOuterLightUpLeft()
	case "boxOuterLightUpLeftRight":
		return BoxOuterLightUpLeftRight()
	case "boxOuterLightUpRight":
		return BoxOuterLightUpRight()
	case "briefcase":
		return Briefcase()
	case "bug":
		return Bug()
	case "bugFill":
		return BugFill()
	case "calculator":
		return Calculator()
	case "calendar":
		return Calendar()
	case "calendarMonth":
		return CalendarMonth()
	case "cancel":
		return Cancel()
	case "card":
		return Card()
	case "cardText":
		return CardText()
	case "cart":
		return Cart()
	case "chartBar":
		return ChartBar()
	case "chat":
		return Chat()
	case "check":
		return Check()
	case "checkboxBlank":
		return CheckboxBlank()
	case "checkboxCross":
		return CheckboxCross()
	case "checkboxMarked":
		return CheckboxMarked()
	case "chevronDown":
		return ChevronDown()
	case "chevronDownCircle":
		return ChevronDownCircle()
	case "chevronLeft":
		return ChevronLeft()
	case "chevronLeftCircle":
		return ChevronLeftCircle()
	case "chevronRight":
		return ChevronRight()
	case "chevronRightCircle":
		return ChevronRightCircle()
	case "chevronUp":
		return ChevronUp()
	case "chevronUpCircle":
		return ChevronUpCircle()
	case "circle":
		return Circle()
	case "clipboard":
		return Clipboard()
	case "clock":
		return Clock()
	case "coffee":
		return Coffee()
	case "comment":
		return Comment()
	case "commentText":
		return CommentText()
	case "compass":
		return Compass()
	case "compassEastArrow":
		return CompassEastArrow()
	case "compassNorthArrow":
		return CompassNorthArrow()
	case "compassNorthEast":
		return CompassNorthEast()
	case "compassNorthWest":
		return CompassNorthWest()
	case "compassSouthArrow":
		return CompassSouthArrow()
	case "compassSouthEast":
		return CompassSouthEast()
	case "compassSouthWest":
		return CompassSouthWest()
	case "compassWestArrow":
		return CompassWestArrow()
	case "creditCard":
		return CreditCard()
	case "crown":
		return Crown()
	case "cubeUnfolded":
		return CubeUnfolded()
	case "database":
		return Database()
	case "device":
		return Device()
	case "diamond":
		return Diamond()
	case "division":
		return Division()
	case "door":
		return Door()
	case "doorBox":
		return DoorBox()
	case "doorOpen":
		return DoorOpen()
	case "download":
		return Download()
	case "email":
		return Email()
	case "file":
		return File()
	case "fire":
		return Fire()
	case "flask":
		return Flask()
	case "flaskEmpty":
		return FlaskEmpty()
	case "flaskRoundBottom":
		return FlaskRoundBottom()
	case "flaskRoundBottomEmpty":
		return FlaskRoundBottomEmpty()
	case "floppyDisk":
		return FloppyDisk()
	case "folder":
		return Folder()
	case "folderOpen":
		return FolderOpen()
	case "gamepadCenter":
		return GamepadCenter()
	case "gamepadDown":
		return GamepadDown()
	case "gamepadDownLeft":
		return GamepadDownLeft()
	case "gamepadDownRight":
		return GamepadDownRight()
	case "gamepadEmpty":
		return GamepadEmpty()
	case "gamepadLeft":
		return GamepadLeft()
	case "gamepadRight":
		return GamepadRight()
	case "gamepadUp":
		return GamepadUp()
	case "gamepadUpLeft":
		return GamepadUpLeft()
	case "gamepadUpRight":
		return GamepadUpRight()
	case "heart":
		return Heart()
	case "image":
		return Image()
	case "label":
		return Label()
	case "labelVariant":
		return LabelVariant()
	case "led":
		return Led()
	case "lightbulb":
		return Lightbulb()
	case "lock":
		return Lock()
	case "lockOpen":
		return LockOpen()
	case "login":
		return Login()
	case "logout":
		return Logout()
	case "map":
		return Map()
	case "menuBottomLeft":
		return MenuBottomLeft()
	case "menuBottomRight":
		return MenuBottomRight()
	case "menuDown":
		return MenuDown()
	case "menuDownFill":
		return MenuDownFill()
	case "menuLeft":
		return MenuLeft()
	case "menuLeftFill":
		return MenuLeftFill()
	case "menuLeftRight":
		return MenuLeftRight()
	case "menuRight":
		return MenuRight()
	case "menuRightFill":
		return MenuRightFill()
	case "menuTopLeft":
		return MenuTopLeft()
	case "menuTopRight":
		return MenuTopRight()
	case "menuUp":
		return MenuUp()
	case "menuUpDown":
		return MenuUpDown()
	case "menuUpFill":
		return MenuUpFill()
	case "message":
		return Message()
	case "messageProcessing":
		return MessageProcessing()
	case "messageText":
		return MessageText()
	case "microphone":
		return Microphone()
	case "minus":
		return Minus()
	case "minusBox":
		return MinusBox()
	case "minusBoxFill":
		return MinusBoxFill()
	case "minusCircle":
		return MinusCircle()
	case "minusCircleFill":
		return MinusCircleFill()
	case "monitor":
		return Monitor()
	case "monitorImage":
		return MonitorImage()
	case "multiply":
		return Multiply()
	case "musicNote":
		return MusicNote()
	case "necklace":
		return Necklace()
	case "note":
		return Note()
	case "notebook":
		return Notebook()
	case "notification":
		return Notification()
	case "octagon":
		return Octagon()
	case "octagonAlert":
		return OctagonAlert()
	case "pause":
		return Pause()
	case "pencil":
		return Pencil()
	case "pickaxe":
		return Pickaxe()
	case "pictogrammers":
		return Pictogrammers()
	case "play":
		return Play()
	case "plus":
		return Plus()
	case "plusBox":
		return PlusBox()
	case "plusBoxFill":
		return PlusBoxFill()
	case "plusCircle":
		return PlusCircle()
	case "plusCircleFill":
		return PlusCircleFill()
	case "radiobox":
		return Radiobox()
	case "radioboxMarked":
		return RadioboxMarked()
	case "removeCircle":
		return RemoveCircle()
	case "rotateClockwise":
		return RotateClockwise()
	case "rotateCounterclockwise":
		return RotateCounterclockwise()
	case "script":
		return Script()
	case "shield":
		return Shield()
	case "skull":
		return Skull()
	case "speaker":
		return Speaker()
	case "stop":
		return Stop()
	case "sword":
		return Sword()
	case "tag":
		return Tag()
	case "tagText":
		return TagText()
	case "target":
		return Target()
	case "terminal":
		return Terminal()
	case "textBox":
		return TextBox()
	case "textImage":
		return TextImage()
	case "toggleSwitchOff":
		return ToggleSwitchOff()
	case "toggleSwitchOn":
		return ToggleSwitchOn()
	case "toolbox":
		return Toolbox()
	case "tooltipAbove":
		return TooltipAbove()
	case "tooltipAboveAlert":
		return TooltipAboveAlert()
	case "tooltipAboveText":
		return TooltipAboveText()
	case "tooltipBelow":
		return TooltipBelow()
	case "tooltipBelowAlert":
		return TooltipBelowAlert()
	case "tooltipBelowText":
		return TooltipBelowText()
	case "tooltipEnd":
		return TooltipEnd()
	case "tooltipEndAlert":
		return TooltipEndAlert()
	case "tooltipEndText":
		return TooltipEndText()
	case "tooltipStart":
		return TooltipStart()
	case "tooltipStartAlert":
		return TooltipStartAlert()
	case "tooltipStartText":
		return TooltipStartText()
	case "trash":
		return Trash()
	case "upload":
		return Upload()
	case "volumeHigh":
		return VolumeHigh()
	case "volumeLow":
		return VolumeLow()
	case "volumeMedium":
		return VolumeMedium()
	case "volumeMute":
		return VolumeMute()
	case "wall":
		return Wall()
	case "wallFill":
		return WallFill()
	case "water":
		return Water()
	case "well":
		return Well()
	default:
		panic("Unknown icon name: " + name)
	}
}

func Account(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M9 3h4v1h1v1h1v4h-1v1h-1v1H9v-1H8V9H7V5h1V4h1V3m1 5v1h2V8h1V6h-1V5h-2v1H9v2h1m-3 4h8v1h2v1h1v1h1v4H3v-4h1v-1h1v-1h2v-1m-1 4H5v1h12v-1h-1v-1h-2v-1H8v1H6v1Z"/>`),
		g.Group(children),
	)
}

func AccountBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m0 14h1v-1h2v-1h8v1h2v1h1V5h-1V4H5v1H4v11m12 2v-1h-2v-1H8v1H6v1h10M9 5h4v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V7h1V6h1V5m3 3V7h-2v1H9v2h1v1h2v-1h1V8h-1Z"/>`),
		g.Group(children),
	)
}

func Alert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 20H2v-1H1v-4h1v-2h1v-2h1V9h1V7h1V5h1V3h1V2h6v1h1v2h1v2h1v2h1v2h1v2h1v2h1v4h-1v1M9 6H8v2H7v2H6v2H5v2H4v2H3v2h16v-2h-1v-2h-1v-2h-1v-2h-1V8h-1V6h-1V4H9v2m1 1h2v6h-2V7m0 7h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}

func AlertBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 12h-2V6h2m0 10h-2v-2h2m6 6H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1m-1-1v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		g.Group(children),
	)
}

func AlertBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-6-8V6h-2v6Zm0 4v-2h-2v2Z"/>`),
		g.Group(children),
	)
}

func AlertCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 12h-2V6h2Zm0 4h-2v-2h2Zm3 5H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-1-2v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1Z"/>`),
		g.Group(children),
	)
}

func AlertCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-3-9V6h-2v6Zm0 4v-2h-2v2Z"/>`),
		g.Group(children),
	)
}

func AlertRhombus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 12h-2V6h2Zm0 4h-2v-2h2Zm0 5h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h1V1h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm0-3v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-2v1H9v1H8v1H7v1H6v1H5v1H4v2h1v1h1v1h1v1h1v1h1v1h1v1Z"/>`),
		g.Group(children),
	)
}

func AlertRhombusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 21h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h1V1h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1Zm0-9V6h-2v6Zm0 4v-2h-2v2Z"/>`),
		g.Group(children),
	)
}

func AlphaA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-1 3v1h1v9h-2v-4h-2v4H8V7h1V6h4m-1 2h-2v2h2V8Z"/>`),
		g.Group(children),
	)
}

func AlphaAFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 8h2v2h-2V8m5-7v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-2 5H9v1H8v9h2v-4h2v4h2V7h-1V6Z"/>`),
		g.Group(children),
	)
}

func AlphaB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v3h-1v2h1v3h-1v1H8V6m2 2v2h2V8h-2m2 4h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}

func AlphaBFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h5v-1h1v-3h-1v-2h1V7h-1V6H8m2 2h2v2h-2V8m2 4v2h-2v-2h2Z"/>`),
		g.Group(children),
	)
}

func AlphaC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h4v1h1v2h-2V8h-2v6h2v-1h2v2h-1v1H9v-1H8V7h1V6Z"/>`),
		g.Group(children),
	)
}

func AlphaCFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h4v-1h1v-2h-2v1h-2V8h2v1h2V7h-1V6H9Z"/>`),
		g.Group(children),
	)
}

func AlphaD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v1h1v6h-1v1h-1v1H8V6m2 2v6h2v-1h1V9h-1V8h-2Z"/>`),
		g.Group(children),
	)
}

func AlphaDFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h5v-1h1v-1h1V8h-1V7h-1V6H8m2 2h2v1h1v4h-1v1h-2V8Z"/>`),
		g.Group(children),
	)
}

func AlphaE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v2h-4v2h4v2h-4v2h4v2H8V6Z"/>`),
		g.Group(children),
	)
}

func AlphaEFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h6v-2h-4v-2h4v-2h-4V8h4V6H8Z"/>`),
		g.Group(children),
	)
}

func AlphaF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v2h-4v2h3v2h-3v4H8V6Z"/>`),
		g.Group(children),
	)
}

func AlphaFFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h3v-2h-3V8h4V6H8Z"/>`),
		g.Group(children),
	)
}

func AlphaG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h5v2h-4v6h2v-2h-1v-2h3v5h-1v1H9v-1H8V7h1V6Z"/>`),
		g.Group(children),
	)
}

func AlphaGFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h4v-1h1v-5h-3v2h1v2h-2V8h4V6H9Z"/>`),
		g.Group(children),
	)
}

func AlphaH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v4h2V6h2v10h-2v-4h-2v4H8V6Z"/>`),
		g.Group(children),
	)
}

func AlphaHFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h2v4h2V6h-2v4h-2V6H8Z"/>`),
		g.Group(children),
	)
}

func AlphaI(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-1 3v2h-1v6h1v2H9v-2h1V8H9V6h4Z"/>`),
		g.Group(children),
	)
}

func AlphaIFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-2 5H9v2h1v6H9v2h4v-2h-1V8h1V6Z"/>`),
		g.Group(children),
	)
}

func AlphaJ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-2 3h2v9h-1v1H9v-1H8v-2h2v1h2V6Z"/>`),
		g.Group(children),
	)
}

func AlphaJFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-3 5v8h-2v-1H8v2h1v1h4v-1h1V6h-2Z"/>`),
		g.Group(children),
	)
}

func AlphaK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v2h1V7h1V6h2v2h-1v1h-1v1h-1v1h1v1h1v1h1v3h-2v-2h-1v-1h-1v3H8V6Z"/>`),
		g.Group(children),
	)
}

func AlphaKFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-3h1v1h1v2h2v-3h-1v-1h-1v-1h-1v-1h1V9h1V8h1V6h-2v1h-1v1h-1V6H8Z"/>`),
		g.Group(children),
	)
}

func AlphaL(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v8h4v2H8V6Z"/>`),
		g.Group(children),
	)
}

func AlphaLFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h6v-2h-4V6H8Z"/>`),
		g.Group(children),
	)
}

func AlphaM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M7 6h8v1h1v9h-2V8h-2v7h-2V8H8v8H6V7h1V6Z"/>`),
		g.Group(children),
	)
}

func AlphaMFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M7 6v1H6v9h2V8h2v7h2V8h2v8h2V7h-1V6H7Z"/>`),
		g.Group(children),
	)
}

func AlphaN(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v2h1v2h1V6h2v10h-2v-2h-1v-2h-1v4H8V6Z"/>`),
		g.Group(children),
	)
}

func AlphaNFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h1v2h1v2h2V6h-2v4h-1V8h-1V6H8Z"/>`),
		g.Group(children),
	)
}

func AlphaO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h4v1h1v8h-1v1H9v-1H8V7h1V6m1 2v6h2V8h-2Z"/>`),
		g.Group(children),
	)
}

func AlphaOFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h4v-1h1V7h-1V6H9m1 2h2v6h-2V8Z"/>`),
		g.Group(children),
	)
}

func AlphaP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v4h-1v1h-3v4H8V6m2 2v2h2V8h-2Z"/>`),
		g.Group(children),
	)
}

func AlphaPFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h3v-1h1V7h-1V6H8m2 2h2v2h-2V8Z"/>`),
		g.Group(children),
	)
}

func AlphaQ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h4v1h1v8h1v1h-1v1h-1v-1h-1v-1h-1v1H9v-1H8V7h1V6m1 2v6h1v-1h1V8h-2Z"/>`),
		g.Group(children),
	)
}

func AlphaQFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v8h1v1h2v-1h1v1h1v1h1v-1h1v-1h-1V7h-1V6H9m1 2h2v5h-1v1h-1V8Z"/>`),
		g.Group(children),
	)
}

func AlphaR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h5v1h1v4h-1v2h1v3h-2v-2h-1v-2h-1v4H8V6m2 2v2h2V8h-2Z"/>`),
		g.Group(children),
	)
}

func AlphaRFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v10h2v-4h1v2h1v2h2v-3h-1v-2h1V7h-1V6H8m2 2h2v2h-2V8Z"/>`),
		g.Group(children),
	)
}

func AlphaS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M9 6h5v2h-4v2h3v1h1v4h-1v1H8v-2h4v-2H9v-1H8V7h1V6Z"/>`),
		g.Group(children),
	)
}

func AlphaSFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M9 6v1H8v4h1v1h3v2H8v2h5v-1h1v-4h-1v-1h-3V8h4V6H9Z"/>`),
		g.Group(children),
	)
}

func AlphaT(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v2h-2v8h-2V8H8V6Z"/>`),
		g.Group(children),
	)
}

func AlphaTFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v2h2v8h2V8h2V6H8Z"/>`),
		g.Group(children),
	)
}

func AlphaU(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h2v8h2V6h2v9h-1v1H9v-1H8V6Z"/>`),
		g.Group(children),
	)
}

func AlphaUFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v9h1v1h4v-1h1V6h-2v8h-2V6H8Z"/>`),
		g.Group(children),
	)
}

func AlphaV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M7 6h2v3h1v2h2V9h1V6h2v4h-1v2h-1v2h-1v2h-2v-2H9v-2H8v-2H7V6Z"/>`),
		g.Group(children),
	)
}

func AlphaVFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M7 6v4h1v2h1v2h1v2h2v-2h1v-2h1v-2h1V6h-2v3h-1v2h-2V9H9V6H7Z"/>`),
		g.Group(children),
	)
}

func AlphaW(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M6 6h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1H8v-1H7v-2H6V6Z"/>`),
		g.Group(children),
	)
}

func AlphaWFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M6 6v7h1v2h1v1h2v-1h2v1h2v-1h1v-2h1V6h-2v6h-1v1h-1V8h-2v5H9v-1H8V6H6Z"/>`),
		g.Group(children),
	)
}

func AlphaX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-1 13v-2h-1v-1h-2v1H9v2H7v-3h1v-1h1v-2H8V9H7V6h2v2h1v1h2V8h1V6h2v3h-1v1h-1v2h1v1h1v3h-2Z"/>`),
		g.Group(children),
	)
}

func AlphaXFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-2 15h2v-3h-1v-1h-1v-2h1V9h1V6h-2v2h-1v1h-2V8H9V6H7v3h1v1h1v2H8v1H7v3h2v-2h1v-1h2v1h1v2Z"/>`),
		g.Group(children),
	)
}

func AlphaY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M7 6h2v2h1v2h2V8h1V6h2v3h-1v2h-1v2h-1v3h-2v-3H9v-2H8V9H7V6Z"/>`),
		g.Group(children),
	)
}

func AlphaYFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M7 6v3h1v2h1v2h1v3h2v-3h1v-2h1V9h1V6h-2v2h-1v2h-2V8H9V6H7Z"/>`),
		g.Group(children),
	)
}

func AlphaZ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3M8 6h6v4h-1v1h-1v1h-1v1h-1v1h4v2H8v-4h1v-1h1v-1h1V9h1V8H8V6Z"/>`),
		g.Group(children),
	)
}

func AlphaZFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8M8 6v2h4v1h-1v1h-1v1H9v1H8v4h6v-2h-4v-1h1v-1h1v-1h1v-1h1V6H8Z"/>`),
		g.Group(children),
	)
}

func Application(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h16v1h1v16h-1ZM18 6V4H4v2Zm0 12V8H4v10Z"/>`),
		g.Group(children),
	)
}

func ApplicationCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M11 16H9v-1H8v-4h1v-1h2v2h-1v2h1m4 2h-2v-2h1v-2h-1v-2h2v1h1v4h-1m4 5H3v-1H2V3h1V2h16v1h1v16h-1M18 6V4H4v2m14 12V8H4v10Z"/>`),
		g.Group(children),
	)
}

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h18v6h-1v12H3V8H2V2m15 16V8H5v10h12M4 4v2h14V4H4m3 6h8v2H7v-2Z"/>`),
		g.Group(children),
	)
}

func Arrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 2h-3v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H3v1H2v1H1v1h1v1h1v1h1v1h1v-1h1v-1h1v-3h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h2"/>`),
		g.Group(children),
	)
}

func ArrowBottomLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 17H5V8h2v5h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1v1h1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1h5v2Z"/>`),
		g.Group(children),
	)
}

func ArrowBottomLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 15V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1m4 1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1m9-1H7V8h2v3h1v-1h1V9h1V8h1V7h1v1h1v1h-1v1h-1v1h-1v1h-1v1h3v2Z"/>`),
		g.Group(children),
	)
}

func ArrowBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 8v9H8v-2h5v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7V8H6V7h1V6h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1V8h2Z"/>`),
		g.Group(children),
	)
}

func ArrowBottomRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1m1-4h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1m-1-9v7H8v-2h3v-1h-1v-1H9v-1H8V9H7V8h1V7h1v1h1v1h1v1h1v1h1V8h2Z"/>`),
		g.Group(children),
	)
}

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 17h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V4h2v9h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`),
		g.Group(children),
	)
}

func ArrowDownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 10v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h5V2h8v8h5m-4 2h-3V4H9v8H6v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1Z"/>`),
		g.Group(children),
	)
}

func ArrowDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 10v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V6h2v6h1v-1h1v-1h2m5-3v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1Z"/>`),
		g.Group(children),
	)
}

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 12v2h1v1h1v1h1v1h1v1h2v-2H8v-1H7v-1h4v-1h2v-1h1v-2h1V3h-2v6h-1v2h-2v1H7v-1h1v-1h1V8H7v1H6v1H5v1H4v1"/>`),
		g.Group(children),
	)
}

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 12v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1h-4v-1H9v-1H8v-2H7V3h2v6h1v2h2v1h3v-1h-1v-1h-1V8h2v1h1v1h1v1h1v1"/>`),
		g.Group(children),
	)
}

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 12v-2h1V9h1V8h1V7h1V6h2v2h-1v1H9v1h9v2H9v1h1v1h1v2H9v-1H8v-1H7v-1H6v-1"/>`),
		g.Group(children),
	)
}

func ArrowLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 20h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h2v5h8v8h-8v5m-2-4v-3h8V9h-8V6H9v1H8v1H7v1H6v1H5v2h1v1h1v1h1v1h1v1h1Z"/>`),
		g.Group(children),
	)
}

func ArrowLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 16h-2v-1H9v-1H8v-1H7v-1H6v-2h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h6v2h-6v1h1v1h1v2m3 5H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1m1-4h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1Z"/>`),
		g.Group(children),
	)
}

func ArrowLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 19H8v-1H7v-1H6v-1H5v-1H4v-2h2v1h1v1h1v-4h1V9h1V8h2V7h7v2h-6v1h-2v2h-1v3h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1"/>`),
		g.Group(children),
	)
}

func ArrowLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 10v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H7v1h1v1h1v2H7v-1H6v-1H5v-1H4v-1H3v-2h1V9h1V8h1V7h1V6h2v2H8v1H7v1h8V9h-1V8h-1V6h2v1h1v1h1v1h1v1"/>`),
		g.Group(children),
	)
}

func ArrowLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 3H8v1H7v1H6v1H5v1H4v2h2V8h1V7h1v4h1v2h1v1h2v1h7v-2h-6v-1h-2v-2h-1V7h1v1h1v1h2V7h-1V6h-1V5h-1V4h-1"/>`),
		g.Group(children),
	)
}

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 10v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H4v-2h9V9h-1V8h-1V6h2v1h1v1h1v1h1v1"/>`),
		g.Group(children),
	)
}

func ArrowRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-5H2V7h8V2m2 4v3H4v4h8v3h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6h-1Z"/>`),
		g.Group(children),
	)
}

func ArrowRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1M6 5H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1m4 1h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H6v-2h6V9h-1V8h-1V6Z"/>`),
		g.Group(children),
	)
}

func ArrowRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 19h2v-1h1v-1h1v-1h1v-1h1v-2h-2v1h-1v1h-1v-4h-1V9h-1V8h-2V7H3v2h6v1h2v2h1v3h-1v-1h-1v-1H8v2h1v1h1v1h1v1h1"/>`),
		g.Group(children),
	)
}

func ArrowRightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 3h2v1h1v1h1v1h1v1h1v2h-2V8h-1V7h-1v4h-1v2h-1v1h-2v1H3v-2h6v-1h2v-2h1V7h-1v1h-1v1H8V7h1V6h1V5h1V4h1"/>`),
		g.Group(children),
	)
}

func ArrowTopLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 14V5h9v2H9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7v5H5Z"/>`),
		g.Group(children),
	)
}

func ArrowTopLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1M6 5H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1m1 9V7h7v2h-3v1h1v1h1v1h1v1h1v1h-1v1h-1v-1h-1v-1h-1v-1h-1v-1H9v3H7Z"/>`),
		g.Group(children),
	)
}

func ArrowTopRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 5h9v9h-2V9h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v-1H6v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7H8V5Z"/>`),
		g.Group(children),
	)
}

func ArrowTopRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1M8 7h7v7h-2v-3h-1v1h-1v1h-1v1H9v1H8v-1H7v-1h1v-1h1v-1h1v-1h1V9H8V7Z"/>`),
		g.Group(children),
	)
}

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 5h2v1h1v1h1v1h1v1h1v2h-2v-1h-1V9h-1v9h-2V9H9v1H8v1H6V9h1V8h1V7h1V6h1"/>`),
		g.Group(children),
	)
}

func ArrowUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 12v-2h1V9h1V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-5v8H7v-8H2m4-2h3v8h4v-8h3V9h-1V8h-1V7h-1V6h-1V5h-2v1H9v1H8v1H7v1H6v1Z"/>`),
		g.Group(children),
	)
}

func ArrowUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 12v-2h1V9h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v6h-2v-6H9v1H8v1H6m-5 3V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1m4 1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1Z"/>`),
		g.Group(children),
	)
}

func ArrowUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 3h2v1h1v1h1v1h1v1h1v2h-2V8h-1V7h-1v8h1v-1h1v-1h2v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-2h2v1h1v1h1V7H9v1H8v1H6V7h1V6h1V5h1V4h1"/>`),
		g.Group(children),
	)
}

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 10V8h1V7h1V6h1V5h1V4h2v2H8v1H7v1h4v1h2v1h1v2h1v7h-2v-6h-1v-2h-2v-1H7v1h1v1h1v2H7v-1H6v-1H5v-1H4v-1"/>`),
		g.Group(children),
	)
}

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 10V8h-1V7h-1V6h-1V5h-1V4h-2v2h1v1h1v1h-4v1H9v1H8v2H7v7h2v-6h1v-2h2v-1h3v1h-1v1h-1v2h2v-1h1v-1h1v-1h1v-1"/>`),
		g.Group(children),
	)
}

func Axe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M11 3h2v1h1v1h1v1h2v1h2v1h1v1h1v2h-1v2h-1v1h-1v1h-2v1h-1v-1h-1v-1h-1v-2h-1v-1h-1v-1h-1V9H9V8H8V6h1V5h1V4h1m-1 6v1h1v2h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v-1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1"/>`),
		g.Group(children),
	)
}

func BagPersonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 15H9v2H7v-2H5v4h12v-4m0-6h-1V8h-1V7H7v1H6v1H5v4h12V9m-4 2H9v-1h1V9h2v1h1v1M3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3V8m6-5v2h4V3H9Z"/>`),
		g.Group(children),
	)
}

func BagPersonalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 8h1V6h2V5h1V2h1V1h6v1h1v3h1v1h2v2h1v12h-1v1H4v-1H3V8m6-5v2h4V3H9m8 11H5v1h2v2h1v-2h9v-1m-5-3h1V9h-1V8h-2v1H9v2h1v1h2v-1Z"/>`),
		g.Group(children),
	)
}

func BatteryFifty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7Z"/>`),
		g.Group(children),
	)
}

func BatteryOneHundred(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7m-6 1h2v6h-2V8m3 0h2v6h-2V8Z"/>`),
		g.Group(children),
	)
}

func BatterySeventyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 8h2v6H5V8m3 0h2v6H8V8m10-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7m-6 1h2v6h-2V8Z"/>`),
		g.Group(children),
	)
}

func BatteryTwentyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 8v6H5V8h2m11-3v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5h15m-1 2H4v8h13V7Z"/>`),
		g.Group(children),
	)
}

func BatteryZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 5h15v1h1v3h1v4h-1v3h-1v1H3v-1H2V6h1V5m1 2v8h13V7H4Z"/>`),
		g.Group(children),
	)
}

func BattleAxe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1h-4v1h-1v1H9v1H8v4h4v1h-1v1h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1H2v1H1v1h1v1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v4h4v-1h1v-1h1v-1h1V7h-4V5h-2"/>`),
		g.Group(children),
	)
}

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 2h1V1h14v1h1v18h-1v1H4v-1H3V2m8 7h-1V8H9v1H8v1H7V3H5v16h12V3h-5v7h-1V9Z"/>`),
		g.Group(children),
	)
}

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 2h12v1h1v17h-2v-1h-2v-1h-2v-1h-2v1H8v1H6v1H4V3h1V2m1 2v13h1v-1h2v-1h1v-1h2v1h1v1h2v1h1V4H6Z"/>`),
		g.Group(children),
	)
}

func BorderBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 12h-2v-2h2v2M4 12H2v-2h2v2m8-8h-2V2h2v2m4 0h-2V2h2v2m4 4h-2V6h2v2m0-4h-2V2h2v2m0 12h-2v-2h2v2M4 16H2v-2h2v2M8 4H6V2h2v2M4 4H2V2h2v2m0 4H2V6h2v2m16 12H2v-2h18v2Z"/>`),
		g.Group(children),
	)
}

func BorderBottomLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 4V2h2v2h-2m8 8v-2h2v2h-2m0 4v-2h2v2h-2M6 4V2h2v2H6m12 4V6h2v2h-2m0-4V2h2v2h-2m-4 0V2h2v2h-2M2 20V2h2v16h16v2H2Z"/>`),
		g.Group(children),
	)
}

func BorderBottomLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 4V2h2v2h-2M6 4V2h2v2H6m8 0V2h2v2h-2m4-2h2v18H2V2h2v16h14V2Z"/>`),
		g.Group(children),
	)
}

func BorderBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 12H2v-2h2v2m8-8h-2V2h2v2m4 0h-2V2h2v2M4 16H2v-2h2v2M8 4H6V2h2v2M4 4H2V2h2v2m0 4H2V6h2v2m16 12H2v-2h16V2h2v18Z"/>`),
		g.Group(children),
	)
}

func BorderInside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 10h8V2h2v8h8v2h-8v8h-2v-8H2v-2m4 8h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 2h2v2H2V2m0 4h2v2H2V6m4-4h2v2H6V2m8 0h2v2h-2V2m4 0h2v2h-2V2m0 4h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}

func BorderLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 20v-2h2v2h-2m0-16V2h2v2h-2m8 8v-2h2v2h-2m0 4v-2h2v2h-2m-4 4v-2h2v2h-2m4 0v-2h2v2h-2M6 20v-2h2v2H6M6 4V2h2v2H6m12 4V6h2v2h-2m0-4V2h2v2h-2m-4 0V2h2v2h-2M2 20V2h2v18H2Z"/>`),
		g.Group(children),
	)
}

func BorderLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 2v2h-2V2h2m0 16v2h-2v-2h2M8 2v2H6V2h2m8 0v2h-2V2h2m0 16v2h-2v-2h2m-8 0v2H6v-2h2M4 2v18H2V2h2m14 0h2v18h-2V2Z"/>`),
		g.Group(children),
	)
}

func BorderNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2m-8-8h2v2h-2V2m0 16h2v2h-2v-2m-4 0h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 2h2v2H2V2m0 4h2v2H2V6m4-4h2v2H6V2m8 0h2v2h-2V2m4 0h2v2h-2V2m0 4h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}

func BorderOutside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h18v18H2V2m2 2v14h14V4H4m6 2h2v2h-2V6m0 4h2v2h-2v-2m-4 0h2v2H6v-2m8 0h2v2h-2v-2m-4 4h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 2v2h-2V2h2m0 16v2h-2v-2h2m-8-8v2H2v-2h2m0-4v2H2V6h2m4-4v2H6V2h2M4 2v2H2V2h2m12 0v2h-2V2h2m0 16v2h-2v-2h2M4 14v2H2v-2h2m0 4v2H2v-2h2m4 0v2H6v-2h2M20 2v18h-2V2h2Z"/>`),
		g.Group(children),
	)
}

func BorderTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2m-8 8h2v2h-2v-2m-4 0h2v2H6v-2m-4-4h2v2H2v-2m0 4h2v2H2v-2M2 6h2v2H2V6m16 0h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2M2 2h18v2H2V2Z"/>`),
		g.Group(children),
	)
}

func BorderTopBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 10h2v2H2v-2m16 0h2v2h-2v-2M2 14h2v2H2v-2m0-8h2v2H2V6m16 0h2v2h-2V6m0 8h2v2h-2v-2M2 18h18v2H2v-2M2 4V2h18v2H2Z"/>`),
		g.Group(children),
	)
}

func BorderTopLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 10h2v2h-2v-2m-8 8h2v2h-2v-2m-4 0h2v2H6v-2M18 6h2v2h-2V6m-4 12h2v2h-2v-2m4 0h2v2h-2v-2m0-4h2v2h-2v-2M2 2h18v2H4v16H2V2Z"/>`),
		g.Group(children),
	)
}

func BorderTopLeftBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 10h2v2h-2v-2m0-4h2v2h-2V6m0 8h2v2h-2v-2m2 4v2H2V2h18v2H4v14h16Z"/>`),
		g.Group(children),
	)
}

func BorderTopLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 18v2h-2v-2h2m4 0v2h-2v-2h2m-8 0v2H6v-2h2m-4 2H2V2h18v18h-2V4H4v16Z"/>`),
		g.Group(children),
	)
}

func BorderTopRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 18v2h-2v-2h2m-8-8v2H2v-2h2m0-4v2H2V6h2m12 12v2h-2v-2h2M4 14v2H2v-2h2m0 4v2H2v-2h2m4 0v2H6v-2h2M20 2v18h-2V4H2V2h18Z"/>`),
		g.Group(children),
	)
}

func BorderTopRightBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 12H2v-2h2v2m0 4H2v-2h2v2m0-8H2V6h2v2M2 4V2h18v18H2v-2h16V4H2Z"/>`),
		g.Group(children),
	)
}

func Bow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 3h10v1h2v1h2v1h1v1h1v2h1v2h1v10h-2v-2h-2v-1h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6V9H5V8H4V7H3V5H1m15 13h1v-6h-1v-2h-1V8h-1V7h-2V6h-2V5H4v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1"/>`),
		g.Group(children),
	)
}

func BowArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 3h10v1h2v1h3V4h1V2h3v3h-2v1h-1v3h1v2h1v10h-2v-2h-2v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v1H8v1H7v3H6v1H5v1H4v-1H3v-1H2v-1H1v-1h1v-1h1v-1h3v-1h1v-1h1v-2H7v-1H6V9H5V8H4V7H3V5H1V3m15 15h1v-6h-1v-2h-1V9h-1v1h-1v1h-1v1h-1v1h1v1h1v1h1v1h1v1h1v1M12 7V6h-2V5H4v1h1v1h1v1h1v1h1v1h1v1h1v-1h1V9h1V8h1V7h-1Z"/>`),
		g.Group(children),
	)
}

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m13 3V4H5v1H4v12h1v1h12v-1h1V5h-1Z"/>`),
		g.Group(children),
	)
}

func BoxLightDashedDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 12h-2v-2h2v2m-4 0h-3v-2h3v2m-5 0h-3V9h2v1h1v2M12 0v2h-2V0h2m0 4v3h-2V4h2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDashedDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 0v2h-2V0h2m0 4v3h-2V4h2m0 5v3H9v-2h1V9h2M.002 10H2v2H.002v-2M4 10h3v2H4v-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDashedUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 22v-2h2v2h-2m0-4v-3h2v3h-2m0-5v-3h3v2h-1v1h-2m12-1h-2v-2h2v2m-4 0h-3v-2h3v2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDashedUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 10h2v2H0v-2m4 0h3v2H4v-2m5 0h3v3h-2v-1H9v-2m1 12v-2h2v2h-2m0-4v-3h2v3h-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDoubleHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 12h22v2H0v-2m0-4h22v2H0V8Z"/>`),
		g.Group(children),
	)
}

func BoxLightDoubleRoundDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 0v3h1v2h1v2h1v1h1v1h1v1h2v1h2v1h3v2h-4v-1h-2v-1h-2v-1h-1v-1h-1V9h-1V8h-1V6H9V4H8V0h2m12 10h-2V9h-3V8h-1V7h-1V6h-1V5h-1V2h-1V0h2v1h1v3h1v1h1v1h1v1h3v1h1v2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDoubleRoundDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 12h3v-1h2v-1h2V9h1V8h1V7h1V5h1V3h1V0h2v4h-1v2h-1v2h-1v1h-1v1H9v1H8v1H6v1H4v1H0v-2M10 0v2H9v3H8v1H7v1H6v1H5v1H2v1H0V8h1V7h3V6h1V5h1V4h1V1h1V0h2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDoubleRoundUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 10h-3v1h-2v1h-2v1h-1v1h-1v1h-1v2h-1v2h-1v3H8v-4h1v-2h1v-2h1v-1h1v-1h1v-1h1v-1h2V9h2V8h4v2M12 22v-2h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h2v2h-1v1h-3v1h-1v1h-1v1h-1v3h-1v1h-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDoubleRoundUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 22v-3h-1v-2h-1v-2H9v-1H8v-1H7v-1H5v-1H3v-1H0V8h4v1h2v1h2v1h1v1h1v1h1v1h1v2h1v2h1v4h-2M0 12h2v1h3v1h1v1h1v1h1v1h1v3h1v2H8v-1H7v-3H6v-1H5v-1H4v-1H1v-1H0v-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDoubleVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 22V0h2v22h-2m-4 0V0h2v22H8Z"/>`),
		g.Group(children),
	)
}

func BoxLightDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 12H10V0h2v10h10v2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDownLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 15H9v-1H8v-1H7V9h1V8h1V7h1V0h2v7h1v1h1v1h1v1h7v2h-7v1h-1v1h-1v1m-4-3h1v1h2v-1h1v-2h-1V9h-2v1H9v2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 0v12H0v-2h10V0h2Z"/>`),
		g.Group(children),
	)
}

func BoxLightDownRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 9v4h-1v1h-1v1H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7h1V0h2v7h1v1h1v1h1m-3 4v-1h1v-2h-1V9h-2v1H9v2h1v1h2Z"/>`),
		g.Group(children),
	)
}

func BoxLightHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 10h22v2H0v-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightRoundDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 12h-4v-1h-3v-1h-1V9h-1V8h-1V7h-1V4h-1V0h2v3h1v3h1v1h1v1h1v1h3v1h3v2Z"/>`),
		g.Group(children),
	)
}

func BoxLightRoundDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 0v4h-1v3h-1v1H9v1H8v1H7v1H4v1H0v-2h3V9h3V8h1V7h1V6h1V3h1V0h2Z"/>`),
		g.Group(children),
	)
}

func BoxLightRoundUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 22v-4h1v-3h1v-1h1v-1h1v-1h1v-1h3v-1h4v2h-3v1h-3v1h-1v1h-1v1h-1v3h-1v3h-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightRoundUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 10h4v1h3v1h1v1h1v1h1v1h1v3h1v4h-2v-3H9v-3H8v-1H7v-1H6v-1H3v-1H0v-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 22V10h12v2H12v10h-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightUpLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 13V9h1V8h1V7h4v1h1v1h1v1h7v2h-7v1h-1v1h-1v1h-1v7h-2v-7H9v-1H8v-1H7m3-4v1H9v2h1v1h2v-1h1v-2h-1V9h-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 10h12v12h-2V12H0v-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightUpRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M9 7h4v1h1v1h1v4h-1v1h-1v1h-1v7h-2v-7H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7m4 3h-1V9h-2v1H9v2h1v1h2v-1h1v-2Z"/>`),
		g.Group(children),
	)
}

func BoxLightVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 0v22h-2V0h2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 0h22v22H0V0m2 2v18h18V2H2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v2h2v2H0v-4h2M6 0h4v2H6V0m6 0h4v2h-4V0m6 0h4v4h-2V2h-2V0m0 22v-2h2v-2h2v4h-4m-2 0h-4v-2h4v2m-6 0H6v-2h4v2M20 6h2v4h-2V6m0 6h2v4h-2v-4Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 20h4v2H1v-2m6 0h3v2H7v-2m5 0h4v2h-4v-2m6 0h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 18h2v2h2v2H0v-4m6 2h4v2H6v-2m6 0h4v2h-4v-2m6 0h3v2h-3v-2M0 16v-4h2v4H0m0-6V7h2v3H0m0-5V1h2v4H0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedDownLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 18h2v2h2v2H0v-4m6 2h4v2H6v-2m6 0h4v2h-4v-2m6 0h2v-2h2v4h-4v-2M0 16v-4h2v4H0m0-6V7h2v3H0m0-5V1h2v4H0m22-4v4h-2V1h2m0 6v3h-2V7h2m0 5v4h-2v-4h2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 22v-2h2v-2h2v4h-4m2-6v-4h2v4h-2m0-6V6h2v4h-2m0-6V1h2v3h-2m-4 18h-4v-2h4v2m-6 0H7v-2h3v2m-5 0H1v-2h4v2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0m0-6V1h2v3H0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 21v-4h2v4h-2m0-6v-3h2v3h-2m0-5V6h2v4h-2m0-6V1h2v3h-2M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0m0-6V1h2v3H0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 21v-4h2v4h-2m0-6v-3h2v3h-2m0-5V6h2v4h-2m0-6V1h2v3h-2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 0h4v2H1V0m6 0h3v2H7V0m5 0h4v2h-4V0m6 0h3v2h-3V0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M21 2h-4V0h4v2m-6 0h-3V0h3v2m-5 0H6V0h4v2M4 2H1V0h3v2m17 20h-4v-2h4v2m-6 0h-3v-2h3v2m-5 0H6v-2h4v2m-6 0H1v-2h3v2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedUpDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v2h2v2H0v-4h2M6 0h4v2H6V0m6 0h3v2h-3V0m5 0h4v2h-4V0m4 22h-4v-2h4v2m-6 0h-3v-2h3v2m-5 0H6v-2h4v2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedUpDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 22v-2h2v-2h2v4h-4m2-6v-4h2v4h-2m0-6V6h2v4h-2m0-6V2h-2V0h4v4h-2m-4 18h-4v-2h4v2m-6 0H7v-2h3v2m-5 0H1v-2h4v2M1 0h4v2H1V0m6 0h3v2H7V0m5 0h4v2h-4V0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 0v2H2v2H0V0h4M2 6v4H0V6h2m0 6v4H0v-4h2m0 6v3H0v-3h2M6 0h4v2H6V0m6 0h3v2h-3V0m5 0h4v2h-4V0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedUpLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 4h-2V2h-2V0h4v4m-6-2h-4V0h4v2m-6 0H6V0h4v2M4 2H2v2H0V0h4v2m18 4v4h-2V6h2m0 6v3h-2v-3h2m0 5v4h-2v-4h2M0 21v-4h2v4H0m0-6v-3h2v3H0m0-5V6h2v4H0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDashedUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 4h-2V2h-2V0h4v4m-6-2h-4V0h4v2m-6 0H6V0h4v2M4 2H1V0h3v2m18 4v4h-2V6h2m0 6v3h-2v-3h2m0 5v4h-2v-4h2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 20h22v2H0v-2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 0h2v20h20v2H0V0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDownLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 22H0V0h2v20h18V0h2v22Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 22v-2h20V0h2v22H0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 22V0h2v22H0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 22V0h2v22H0m20 0V0h2v22h-2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 22V0h2v22h-2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightRoundDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 0h2v5h1v4h1v2h1v2h1v1h1v1h1v1h1v1h2v1h2v1h4v1h5v2h-6v-1h-4v-1h-2v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-2H2v-2H1V6H0V0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightRoundDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 22v-2h5v-1h4v-1h2v-1h2v-1h1v-1h1v-1h1v-1h1v-2h1V9h1V5h1V0h2v6h-1v4h-1v2h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-2v1h-2v1H6v1H0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightRoundUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 0v2h-5v1h-4v1h-2v1H9v1H8v1H7v1H6v1H5v2H4v2H3v4H2v5H.004v-6H1v-4h1v-2h1V8h1V7h1V6h1V5h1V4h1V3h2V2h2V1h4V0h6Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightRoundUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 22h-2v-5h-1v-4h-1v-2h-1V9h-1V8h-1V7h-1V6h-1V5h-2V4H9V3H5V2H.004V0H6v1h4v1h2v1h2v1h1v1h1v1h1v1h1v1h1v2h1v2h1v4h1v6Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 0h22v2H0V0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 0h22v2H0V0m0 20h22v2H0v-2Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightUpDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 22V0h22v2H2v18h20v2H0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightUpDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 0v22H0v-2h20V2H0V0h22Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 0v2H2v20H0V0h22Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightUpLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 0h22v22h-2V2H2v20H0V0Z"/>`),
		g.Group(children),
	)
}

func BoxOuterLightUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 22h-2V2H.002V0H22v22Z"/>`),
		g.Group(children),
	)
}

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 6h5V3h1V2h6v1h1v3h5v1h1v12h-1v1H2v-1H1V7h1V6m7 0h4V4H9v2m10 2H3v10h16V8Z"/>`),
		g.Group(children),
	)
}

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 7h3V6h1V5h1V4H7V3H6V2h1V1h1v1h1v1h1v1h2V3h1V2h1V1h1v1h1v1h-1v1h-1v1h1v1h1v1h3v2h-2v2h2v2h-2v2h2v2h-3v1h-1v1h-1v1H8v-1H7v-1H6v-1H3v-2h2v-2H3v-2h2V9H3V7m10 11v-1h1v-1h1V8h-1V7h-1V6H9v1H8v1H7v8h1v1h1v1h4m-4-5h4v2H9v-2m0-4h4v2H9V9Z"/>`),
		g.Group(children),
	)
}

func BugFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 7h3V6h1V5h1V4H7V3H6V2h1V1h1v1h1v1h1v1h2V3h1V2h1V1h1v1h1v1h-1v1h-1v1h1v1h1v1h3v2h-2v2h2v2h-2v2h2v2h-3v1h-1v1h-1v1H8v-1H7v-1H6v-1H3v-2h2v-2H3v-2h2V9H3V7m6 6v2h4v-2H9m0-4v2h4V9H9Z"/>`),
		g.Group(children),
	)
}

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 21H4v-1H3V3h1V2h14v1h1v17h-1ZM16 7V4H6v3Zm-8 4V9H6v2Zm4 0V9h-2v2Zm4 0V9h-2v2Zm-8 4v-2H6v2Zm4 0v-2h-2v2Zm4 0v-2h-2v2Zm-8 4v-2H6v2Zm4 0v-2h-2v2Zm4 0v-2h-2v2Z"/>`),
		g.Group(children),
	)
}

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h2V0h2v2h8V0h2v2h2v1h1v16h-1v1M4 4v2h14V4H4m0 4v10h14V8H4m8 4h4v4h-4v-4Z"/>`),
		g.Group(children),
	)
}

func CalendarMonth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h2V0h2v2h8V0h2v2h2v1h1v16h-1v1M4 4v2h14V4H4m0 4v10h14V8H4m10 6h2v2h-2v-2m-4 0h2v2h-2v-2m-4 0h2v2H6v-2m0-4h2v2H6v-2m4 0h2v2h-2v-2m4 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}

func Cancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-2h1V8h-1V6h-1m-3 10v-1h-1v-1h-1v-1h-1v-1h-1v-1H9v-1H8V9H7V8H6V7H5V6H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h-1v-1h-1Z"/>`),
		g.Group(children),
	)
}

func Card(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 3h18v1h1v14h-1v1H2v-1H1V4h1V3m1 2v12h16V5H3Z"/>`),
		g.Group(children),
	)
}

func CardText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 8v2H5V8h12M5 12h10v2H5v-2M2 3h18v1h1v14h-1v1H2v-1H1V4h1V3m1 2v12h16V5H3Z"/>`),
		g.Group(children),
	)
}

func Cart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 14v2H6v-1H5v-4H4V8H3V3H1V1h4v3h16v4h-1v3h-1v1H7v2h12M5 7h1v3h12V7h1V6H5v1m2 10h2v1h1v2H9v1H7v-1H6v-2h1v-1m8 0h2v1h1v2h-1v1h-2v-1h-1v-2h1v-1Z"/>`),
		g.Group(children),
	)
}

func ChartBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h2v16h16v2H2V2m4 14V8h4v8H6m5 0V4h4v12h-4m5 0v-5h4v5h-4Z"/>`),
		g.Group(children),
	)
}

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 4h9v1h2v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-2v1H8v1H4v1H1v-2h2v-1h1v-2H3v-1H2V8h1V7h1V6h1V5h2V4m10 4V7h-2V6H8v1H6v1H5v1H4v4h1v1h1v1h2v1h7v-1h2v-1h1v-1h1V9h-1V8h-1Z"/>`),
		g.Group(children),
	)
}

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 11h2v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-2Z"/>`),
		g.Group(children),
	)
}

func CheckboxBlank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 4h1V3h14v1h1v14h-1v1H4v-1H3V4m2 13h12V5H5v12Z"/>`),
		g.Group(children),
	)
}

func CheckboxCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 12h1v1h1v2h-2v-1h-1v-1h-2v1H9v1H7v-2h1v-1h1v-2H8V9H7V7h2v1h1v1h2V8h1V7h2v2h-1v1h-1v2m5 7H4v-1H3V4h1V3h14v1h1v14h-1v1M5 5v12h12V5H5Z"/>`),
		g.Group(children),
	)
}

func CheckboxMarked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 4h1V3h14v1h-1v1H5v12h12v-6h1v-1h1v8h-1v1H4v-1H3V4m3 5h2v1h1v1h1v1h2v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6V9Z"/>`),
		g.Group(children),
	)
}

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 10h1V9h1V7h-2v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8V9H7V8H6V7H4v2h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1"/>`),
		g.Group(children),
	)
}

func ChevronDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 9v2h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6V9h2v1h1v1h1v1h2v-1h1v-1h1V9h2m-1-8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`),
		g.Group(children),
	)
}

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 16v1h1v1h2v-2h-1v-1h-1v-1h-1v-1h-1v-1h-1v-2h1V9h1V8h1V7h1V6h1V4h-2v1h-1v1h-1v1h-1v1H9v1H8v1H7v2h1v1h1v1h1v1h1v1"/>`),
		g.Group(children),
	)
}

func ChevronLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 16h-2v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h2v2h-1v1h-1v1h-1v2h1v1h1v1h1v2m2-15v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`),
		g.Group(children),
	)
}

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 6V5H9V4H7v2h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1H9v1H8v1H7v2h2v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6"/>`),
		g.Group(children),
	)
}

func ChevronRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M9 16h2v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H9v2h1v1h1v1h1v2h-1v1h-1v1H9v2m6-15v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`),
		g.Group(children),
	)
}

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 12H5v1H4v2h2v-1h1v-1h1v-1h1v-1h1v-1h2v1h1v1h1v1h1v1h1v1h2v-2h-1v-1h-1v-1h-1v-1h-1V9h-1V8h-1V7h-2v1H9v1H8v1H7v1H6"/>`),
		g.Group(children),
	)
}

func ChevronUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 13v-2h1v-1h1V9h1V8h1V7h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6m9-12v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`),
		g.Group(children),
	)
}

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`),
		g.Group(children),
	)
}

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 5h1V4h4V2h2V1h4v1h2v2h4v1h1v15h-1v1H3v-1H2V5m8-2v2h2V3h-2m8 3h-2v2H6V6H4v13h14V6Z"/>`),
		g.Group(children),
	)
}

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 5h2v6h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-1V5m5-4v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3Z"/>`),
		g.Group(children),
	)
}

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 20v-2h16v2H1M2 3h17v1h1v6h-1v1h-3v3h-1v1h-1v1H4v-1H3v-1H2V3m14 2v4h2V5h-2M4 5v8h1v1h8v-1h1V5H4Z"/>`),
		g.Group(children),
	)
}

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h18v1h1v14h-1v1h-8v1h-1v1h-1v1H6v-3H2v-1H1V3h1V2m1 2v12h5v3h1v-1h1v-1h1v-1h8V4H3Z"/>`),
		g.Group(children),
	)
}

func CommentText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h18v1h1v14h-1v1h-8v1h-1v1h-1v1H6v-3H2v-1H1V3h1V2m1 2v12h5v3h1v-1h1v-1h1v-1h8V4H3m2 3h12v2H5V7m0 4h10v2H5v-2Z"/>`),
		g.Group(children),
	)
}

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 1v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8m-1 2H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3m-4 6V8h2V7h2V6h2v2h-1v2h-1v2h-1v1h-1v1h-2v1H8v1H6v-2h1v-2h1v-2h1V9h1m2 1h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}

func CompassEastArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 10v2H5v1H4v1H3v1H2v1H0v-2h1v-1h1v-1h1v-2H2V9H1V8H0V6h2v1h1v1h1v1h1v1m5-4h6v2h-4v2h4v2h-4v2h4v2h-6"/>`),
		g.Group(children),
	)
}

func CompassNorthArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 4h2v1.5h1V8h1V4h2v10h-2v-2h-1v-2h-1v4H8m2 3v-1h2v1h1v1h1v1h1v1h1v2h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6v-2h1v-1h1v-1h1v-1"/>`),
		g.Group(children),
	)
}

func CompassNorthEast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 6h2v1.5h1V10h1V6h2v10H8v-2H7v-2H6v4H4m8-10h6v2h-4v2h4v2h-4v2h4v2h-6"/>`),
		g.Group(children),
	)
}

func CompassNorthWest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 6h2v1.5h1V10h1V6h2v10H6v-2H5v-2H4v4H2m8-10h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1h-2v-1h-1v-2h-1"/>`),
		g.Group(children),
	)
}

func CompassSouthArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 6h2V5h1V4h1V3h1V2h1V0h-2v1h-1v1h-1v1h-2V2H9V1H8V0H6v2h1v1h1v1h1v1h1M9 8h5v2h-4v2h3v1h1v4h-1v1H8v-2h4v-2H9v-1H8V9h1"/>`),
		g.Group(children),
	)
}

func CompassSouthEast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 6h5v2H6v2h3v1h1v4H9v1H4v-2h4v-2H5v-1H4V7h1m7-1h6v2h-4v2h4v2h-4v2h4v2h-6"/>`),
		g.Group(children),
	)
}

func CompassSouthWest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 6h5v2H4v2h3v1h1v4H7v1H2v-2h4v-2H3v-1H2V7h1m7-1h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1h-2v-1h-1v-2h-1"/>`),
		g.Group(children),
	)
}

func CompassWestArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 6h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1H8v1H6v-1H5v-2H4m12-3v2h1v1h1v1h1v1h1v1h2v-2h-1v-1h-1v-1h-1v-2h1V9h1V8h1V6h-2v1h-1v1h-1v1h-1v1"/>`),
		g.Group(children),
	)
}

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 4h18v1h1v12h-1v1H2v-1H1V5h1V4m1 2v2h16V6H3m0 10h16v-5H3v5Z"/>`),
		g.Group(children),
	)
}

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 17h18v2H2v-2M4 6v1h1v1h1V7h1V6h1V5h1V4h1V3h2v1h1v1h1v1h1v1h1v1h1V7h1V6h1V5h1v11H2V5h1v1h1m3 8h11v-4h-3V9h-1V8h-1V7h-1V6h-2v1H9v1H8v1H7v1H4v4h3Z"/>`),
		g.Group(children),
	)
}

func CubeUnfolded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 3v5h10v7h-5v5h-7v-5H0V8h5V3h7m-2 2H7v3h3V5m-3 5v3h3v-3H7m-2 0H2v3h3v-3m12 0v3h3v-3h-3m-2 5h-3v3h3v-3m-3-5v3h3v-3h-3Z"/>`),
		g.Group(children),
	)
}

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 2h8v1h2v1h1v1h1v12h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3V5h1V4h1V3h2V2m1 14v-1H6v-1H5v2h1v1h2v1h6v-1h2v-1h1v-2h-1v1h-2v1H8m0-5v-1H6V9H5v3h2v1h2v1h4v-1h2v-1h2V9h-1v1h-2v1H8m1-3v1h4V8h2V7h2V6h-1V5h-2V4H8v1H6v1H5v1h2v1h2Z"/>`),
		g.Group(children),
	)
}

func Device(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h.94v18H20v1H2v-1h-.94V2H2V1m1 2v16h16V3H3m1 1h14v8H4V4m1 10h3v3H5v-3m7 1h2v2h-2v-2m3-1h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 2h10v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2V6h1V5h1V4h1V3h1V2m9 3V4h-1v2h1v1h2V6h-1V5h-1m-3 1V4h-2v2H9v1h4V6h-1M8 6V4H7v1H6v1H5v1h2V6h1m-4 5h1v1h1v1h1v1h1v-2H7V9H4v2m6 1v4h2v-4h1V9H9v3h1m4 0v2h1v-1h1v-1h1v-1h1V9h-3v3h-1Z"/>`),
		g.Group(children),
	)
}

func Division(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 8h-2V7H9V5h1V4h2v1h1v2h-1Zm5 4H5v-2h12Zm-5 6h-2v-1H9v-2h1v-1h2v1h1v2h-1Z"/>`),
		g.Group(children),
	)
}

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 10h2v2h-2v-2m4-8v1h1v15h2v2H3v-2h2V3h1V2h10m-1 2H7v14h8V4Z"/>`),
		g.Group(children),
	)
}

func DoorBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 14h-2v-2h2Zm3 4h1v-1h1V5h-1V4H5v1H4v12h1v1h1V6h10Zm2 2H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-4-2V8H8v10Z"/>`),
		g.Group(children),
	)
}

func DoorOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 10v2H9v-2h1M6 2h10v1h1v15h2v2H3v-2h2V3h1V2m1 2v14h4V4H7m6 0v1h1V4h-1m1 1v1h1V5h-1m0 1h-1v1h1V6m0 1v1h1V7h-1m0 1h-1v1h1V8m0 1v1h1V9h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1Z"/>`),
		g.Group(children),
	)
}

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 17v2H4v-2h14M14 2v6h4v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V8h4V2h6m-2 2h-2v6H9v1h1v1h2v-1h1v-1h-1V4Z"/>`),
		g.Group(children),
	)
}

func Email(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 5h1V4h18v1h1v13h-1v1H2v-1H1V5m2 12h16V9h-1v1h-2v1h-2v1h-2v1h-2v-1H8v-1H6v-1H4V9H3v8M19 6H3v1h2v1h2v1h2v1h4V9h2V8h2V7h2V6Z"/>`),
		g.Group(children),
	)
}

func File(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 1v1h1v1h1v1h1v1h1v1h1v1h1v13h-1v1H4v-1H3V2h1V1h9m0 3h-1v4h4V7h-1V6h-1V5h-1V4M5 3v16h12v-9h-6V9h-1V3H5Z"/>`),
		g.Group(children),
	)
}

func Fire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 20H7v-1H6v-1H5v-1H4v-5h1v-2h1V9h1V8h1v1h1v2h1V9h1V5h-1V4H9V3H8V2h3v1h2v1h1v1h1v1h1v1h1v2h1v7h-1v2h-1v1h-2m-2-1v-1h2v-1h1v-2h1v-4h-1V8h-1V7h-1v4h-1v2h-1v1h-1v1H9v-1H8v-3H7v1H6v4h1v1h1v1Z"/>`),
		g.Group(children),
	)
}

func Flask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M11 12h2v2h-2v-2m3-11v1h1v3h-1v2h1v2h1v2h1v2h1v1h1v2h1v4h-1v1H3v-1H2v-4h1v-2h1v-1h1v-2h1V9h1V7h1V5H7V2h1V1h6m-2 2h-2v5H9v2H8v2H7v1H6v2H5v2h1v-1h1v-1h1v-1h1v1h1v1h1v1h1v1h2v-1h1v-2h1v-1h-1v-2h-1v-2h-1V8h-1V3Z"/>`),
		g.Group(children),
	)
}

func FlaskEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 1h6v1h1v3h-1v2h1v2h1v2h1v2h1v1h1v2h1v4h-1v1H3v-1H2v-4h1v-2h1v-1h1v-2h1V9h1V7h1V5H7V2h1V1m2 2v5H9v2H8v2H7v1H6v2H5v2H4v2h14v-2h-1v-2h-1v-1h-1v-2h-1v-2h-1V8h-1V3h-2Z"/>`),
		g.Group(children),
	)
}

func FlaskRoundBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M11 11h2v2h-2v-2m2-10v2h1v5h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-6h1v-1h1V9h1V8h1V3h1V1h4m-1 4h-2v4H9v1H8v1H7v1H6v2h1v-1h2v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h-1v-1h-1v-1h-1V9h-1V5Z"/>`),
		g.Group(children),
	)
}

func FlaskRoundBottomEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M9 1h4v2h1v5h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1H8v-1H7v-1H6v-1H5v-1H4v-6h1v-1h1V9h1V8h1V3h1V1m1 4v4H9v1H8v1H7v1H6v4h1v1h1v1h1v1h4v-1h1v-1h1v-1h1v-4h-1v-1h-1v-1h-1V9h-1V5h-2Z"/>`),
		g.Group(children),
	)
}

func FloppyDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 3h1V2h13v1h1v1h1v1h1v1h1v13h-1v1H3v-1H2V3m16 4h-1V6h-1V5h-1v4H6V4H4v14h2v-5h10v5h2V7m-7-3v3h2V4h-2m3 14v-3H8v3h6Z"/>`),
		g.Group(children),
	)
}

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 3h7v1h1v1h10v1h1v12h-1v1H2v-1H1V4h1V3m1 4v10h16V7H3Z"/>`),
		g.Group(children),
	)
}

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 4h1V3h7v1h1v1h10v1h1v12h-1v1H2v-1H1V4m2 5h16V7H9V6H8V5H3v4m0 8h16v-6H3v6Z"/>`),
		g.Group(children),
	)
}

func GamepadCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 7H9v1H8v4h1v1h4v-1h1V9h-1V8Z"/>`),
		g.Group(children),
	)
}

func GamepadDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 14H9v4h4v-4Z"/>`),
		g.Group(children),
	)
}

func GamepadDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6M7 9H3v4h4V9m2 6v4h4v-4H9Z"/>`),
		g.Group(children),
	)
}

func GamepadDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m5 8h-4v4h4V9M9 15v4h4v-4H9Z"/>`),
		g.Group(children),
	)
}

func GamepadEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 2H9v6H3v4h6v6h4v-6h6V9h-6V3Z"/>`),
		g.Group(children),
	)
}

func GamepadLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6M7 9H3v4h4V9Z"/>`),
		g.Group(children),
	)
}

func GamepadRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m5 8h-4v4h4V9Z"/>`),
		g.Group(children),
	)
}

func GamepadUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m-1 2H9v4h4V3Z"/>`),
		g.Group(children),
	)
}

func GamepadUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6M7 9H3v4h4V9m2-6v4h4V3H9Z"/>`),
		g.Group(children),
	)
}

func GamepadUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 1v1h1v5h5v1h1v6h-1v1h-5v5h-1v1H8v-1H7v-5H2v-1H1V8h1V7h5V2h1V1h6m5 8h-4v4h4V9M9 3v4h4V3H9Z"/>`),
		g.Group(children),
	)
}

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 20h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2H1V5h1V4h1V3h1V2h5v1h1v1h2V3h1V2h5v1h1v1h1v1h1v5h-1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1m-7-9v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V6h-1V5h-1V4h-3v1h-1v1h-1v1h-2V6H9V5H8V4H5v1H4v1H3v3h1v2h1Z"/>`),
		g.Group(children),
	)
}

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 4h1V3h18v1h1v14h-1v1H2v-1H1V4m2 10h1v-1h1v-1h1v-1h1v-1h1V9h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2V5H3v9m11 3v-1h-1v-1h-1v-1h-1v-1h-1v-1H8v1H7v1H6v1H5v1H4v1h10m-1-9h1V7h2v1h1v2h-1v1h-2v-1h-1V8Z"/>`),
		g.Group(children),
	)
}

func Label(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 4h13v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-1H1V5h1V4m14 9h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H3v10h11v-1h1v-1h1v-1Z"/>`),
		g.Group(children),
	)
}

func LabelVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 4v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-2h1v-1h1v-1h1v-1h1v-1h1v-2H6V9H5V8H4V7H3V6H2V4h13m-1 12v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H6v1h1v1h1v1h1v1h1v2H9v1H8v1H7v1H6v1h8Z"/>`),
		g.Group(children),
	)
}

func Led(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 21v-6H4v-2h2V6h1V4h1V3h1V2h4v1h1v1h1v2h1v7h2v2h-4v6h-2v-6h-2v6H8m4-16V4h-2v1H9v2H8v6h1v-1h4v1h1V7h-1V5h-1Z"/>`),
		g.Group(children),
	)
}

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 19h6v2H8v-2m0-1v-4H7v-1H6v-1H5v-1H4V5h1V4h1V3h1V2h1V1h6v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v4H8m5-6h1v-1h1v-1h1V6h-1V5h-1V4h-1V3H9v1H8v1H7v1H6v4h1v1h1v1h1v1h1v3h2v-3h1v-1Z"/>`),
		g.Group(children),
	)
}

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 12h2v1h1v2h-1v1h-2v-1H9v-2h1v-1M8 2h6v1h1v1h1v4h1v1h1v10h-1v1H5v-1H4V9h1V8h1V4h1V3h1V2m1 2v1H8v3h6V5h-1V4H9m7 6H6v8h10v-8Z"/>`),
		g.Group(children),
	)
}

func LockOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 13h2v1h1v2h-1v1h-2v-1H9v-2h1v-1m4-11v1h1v1h1v5h1v1h1v10h-1v1H5v-1H4V10h1V9h9V5h-1V4H9v1H8v2H6V4h1V3h1V2h6m2 9H6v8h10v-8Z"/>`),
		g.Group(children),
	)
}

func Login(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 1h12v1h1v18h-1v1H5v-1H4v-6h2v5h10V3H6v5H4V2h1V1m3 5h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1H8v-2h1v-1h1v-1H2v-2h8V9H9V8H8V6Z"/>`),
		g.Group(children),
	)
}

func Logout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 1v1h1v4h-1V5h-1V3H6v16h10v-2h1v-1h1v4h-1v1H5v-1H4V2h1V1h12m-4 5h2v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1H8v-2h7V9h-1V8h-1V6Z"/>`),
		g.Group(children),
	)
}

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 4h2V3h2V2h4v1h3v1h2V3h2V2h3v16h-2v1h-2v1h-4v-1H9v-1H7v1H5v1H2V4m2 2v11h2v-1h1V5H5v1H4m8-1H9v11h1v1h3V6h-1V5m4 1h-1v11h2v-1h1V5h-2v1Z"/>`),
		g.Group(children),
	)
}

func MenuBottomLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 4h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2H7V4m2 4v5h5v-1h-1v-1h-1v-1h-1V9h-1V8H9Z"/>`),
		g.Group(children),
	)
}

func MenuBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 15v-2h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v11H4m4-2h5V8h-1v1h-1v1h-1v1H9v1H8v1Z"/>`),
		g.Group(children),
	)
}

func MenuDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 8h14v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4V8m4 2v1h1v1h1v1h2v-1h1v-1h1v-1H8Z"/>`),
		g.Group(children),
	)
}

func MenuDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 9V8H5v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1V9"/>`),
		g.Group(children),
	)
}

func MenuLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 4v14h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-2h1V9h1V8h1V7h1V6h1V5h1V4h2m-2 4h-1v1h-1v1h-1v2h1v1h1v1h1V8Z"/>`),
		g.Group(children),
	)
}

func MenuLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 17h1V5h-1v1h-1v1h-1v1h-1v1H9v1H8v2h1v1h1v1h1v1h1v1h1"/>`),
		g.Group(children),
	)
}

func MenuLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 4h2v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2V4m-2 0v14H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-2h1V9h1V8h1V7h1V6h1V5h1V4h2m4 4v6h1v-1h1v-1h1v-2h-1V9h-1V8h-1M8 8H7v1H6v1H5v2h1v1h1v1h1V8Z"/>`),
		g.Group(children),
	)
}

func MenuRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 18V4h2v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7m2-4h1v-1h1v-1h1v-2h-1V9h-1V8H9v6Z"/>`),
		g.Group(children),
	)
}

func MenuRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M9 5H8v12h1v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H9"/>`),
		g.Group(children),
	)
}

func MenuTopLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 7v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7V7h11m-4 2H9v5h1v-1h1v-1h1v-1h1v-1h1V9Z"/>`),
		g.Group(children),
	)
}

func MenuTopRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 18h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V7h11v11m-2-4V9H8v1h1v1h1v1h1v1h1v1h1Z"/>`),
		g.Group(children),
	)
}

func MenuUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 14H4v-2h1v-1h1v-1h1V9h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v1h1v1h1v2m-4-2v-1h-1v-1h-1V9h-2v1H9v1H8v1h6Z"/>`),
		g.Group(children),
	)
}

func MenuUpDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 10V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v2H4m0 2h14v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-2m4-4h6V7h-1V6h-1V5h-2v1H9v1H8v1m0 6v1h1v1h1v1h2v-1h1v-1h1v-1H8Z"/>`),
		g.Group(children),
	)
}

func MenuUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 13v1h12v-1h-1v-1h-1v-1h-1v-1h-1V9h-1V8h-2v1H9v1H8v1H7v1H6v1"/>`),
		g.Group(children),
	)
}

func Message(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m1 2v13h1v-1h15V3H3Z"/>`),
		g.Group(children),
	)
}

func MessageProcessing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m2 14h15V3H3v13h1v-1m2-7h2v2H6V8m4 0h2v2h-2V8m4 0h2v2h-2V8Z"/>`),
		g.Group(children),
	)
}

func MessageText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m1 2v13h1v-1h15V3H3m2 3h12v2H5V6m0 4h9v2H5v-2Z"/>`),
		g.Group(children),
	)
}

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 21v-3H8v-1H6v-1H5v-1H4v-2H3V9h2v3h1v2h1v1h2v1h4v-1h2v-1h1v-2h1V9h2v4h-1v2h-1v1h-1v1h-2v1h-2v3h-2m-2-8v-1H7V3h1V2h1V1h4v1h1v1h1v9h-1v1h-1v1H9v-1H8m1-2h1v1h2v-1h1V4h-1V3h-2v1H9v7Z"/>`),
		g.Group(children),
	)
}

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 12H5v-2h12Z"/>`),
		g.Group(children),
	)
}

func MinusBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 12H6v-2h10Zm2 8H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-1-2v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		g.Group(children),
	)
}

func MinusBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-2-8v-2H6v2Z"/>`),
		g.Group(children),
	)
}

func MinusCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1m-1 4v2H6v-2h10Z"/>`),
		g.Group(children),
	)
}

func MinusCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2m1-8v-2H6v2Z"/>`),
		g.Group(children),
	)
}

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h18v1h1v12h-1v1h-7v2h2v2H7v-2h2v-2H2v-1H1V3h1V2m1 2v10h16V4H3Z"/>`),
		g.Group(children),
	)
}

func MonitorImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 6h2v1h1v2h-1v1h-2V9h-1V7h1V6M2 2h18v1h1v12h-1v1h-7v2h2v2H7v-2h2v-2H2v-1H1V3h1V2m1 2v5h1V8h1V7h1V6h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h4V4H3m7 7H9v-1H8V9H6v1H5v1H4v1H3v2h9v-1h-1v-1h-1v-1Z"/>`),
		g.Group(children),
	)
}

func Multiply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 16h-2v-1h-1v-1h-1v-1h-2v1H9v1H8v1H6v-2h1v-1h1v-1h1v-2H8V9H7V8H6V6h2v1h1v1h1v1h2V8h1V7h1V6h2v2h-1v1h-1v1h-1v2h1v1h1v1h1Z"/>`),
		g.Group(children),
	)
}

func MusicNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M11 2h7v5h-5v11h-1v1h-1v1H7v-1H6v-1H5v-4h1v-1h1v-1h4V2m0 13h-1v-1H8v1H7v2h1v1h2v-1h1v-2Z"/>`),
		g.Group(children),
	)
}

func Necklace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M9 17h1v-1h2v1h1v2h-1v1h-2v-1H9v-2m1-2v-1H9v-1h-.91v-1H7v-2H6V8H5V6H4V3h1V2h12v1h1v3h-1v2h-1v2h-1v2h-1v1h-1v1h-1v1h-2M7 5v2h1v2h1v2h1.09v1H12v-1h1V9h1V7h1V5h1V4H6v1h1Z"/>`),
		g.Group(children),
	)
}

func Note(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 3v1h1v1h1v1h1v1h1v1h1v1h1v9h-1v1H2v-1H1V4h1V3h13m0 3h-1v4h4V9h-1V8h-1V7h-1V6M3 5v12h16v-5h-6v-1h-1V5H3Z"/>`),
		g.Group(children),
	)
}

func Notebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 2v18h-1v1H4v-1H3v-2H1v-2h2v-4H1v-2h2V6H1V4h2V2h1V1h14v1h1m-5 7h-1V8h-1v1h-1v1h-1V3H7v16h10V3h-2v7h-1V9M3 4v2h2V4H3m2 6H3v2h2v-2m0 6H3v2h2v-2Z"/>`),
		g.Group(children),
	)
}

func Notification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 20v1h-2v-1H8v-2h6v2h-2M2 15h1v-1h1V6h1V4h1V3h2V2h2V1h2v1h2v1h2v1h1v2h1v8h1v1h1v2H2v-2m4 0h10V7h-1V5h-2V4H9v1H7v2H6v8Z"/>`),
		g.Group(children),
	)
}

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1V6m13 11h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1h6v-1h1v-1Z"/>`),
		g.Group(children),
	)
}

func OctagonAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 6h1V5h1V4h1V3h1V2h1V1h8v1h1v1h1v1h1v1h1v1h1v1h1v8h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V7h1V6m13 11h1v-1h1v-1h1v-1h1V8h-1V7h-1V6h-1V5h-1V4h-1V3H8v1H7v1H6v1H5v1H4v1H3v6h1v1h1v1h1v1h1v1h1v1h6v-1h1v-1M10 6h2v7h-2V6m0 8h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 4h3v14H6V4m7 14V4h3v14h-3Z"/>`),
		g.Group(children),
	)
}

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 2h1v1h1v1h1v1h1v1h-1v1h-1v1h-1V7h-1V6h-1V5h-1V4h1V3h1m-4 3h2v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H2v-4h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1V9h1V8h1V7h1"/>`),
		g.Group(children),
	)
}

func Pickaxe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 2h3v1h2v1h2v1h2v2h1v2h1v2h1v3h-2v-2h-1v-1h-1v-1h-2V9h-1V8h-1V6h-1V5h-1V4H8m3 5h1v1h1v1h-1v1h-1v1h-1v1H9v1H8v1H7v1H6v1H5v1H4v1H3v1H2v-1H1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1v-1h1"/>`),
		g.Group(children),
	)
}

func Pictogrammers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 0h14v1h1v1h1v18h-1v1h-1v1H4v-1H3v-1H2V2h1V1h1V0m0 18v1h1v1h12v-1h1v-1H4M17 2H5v1H4v12h1v1h12v-1h1V3h-1V2m-4 2v1h1v1h1v2h-1v1h-1v1h-3v4H8V4h5m-1 2h-2v2h2V6Z"/>`),
		g.Group(children),
	)
}

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 5v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1H9v1H7V4h2v1h1m2 5h-1V9h-1V8H9v6h1v-1h1v-1h1v-2Z"/>`),
		g.Group(children),
	)
}

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 17h-2v-5H5v-2h5V5h2v5h5v2h-5Z"/>`),
		g.Group(children),
	)
}

func PlusBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 16h-2v-4H6v-2h4V6h2v4h4v2h-4Zm6 4H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-1-2v-1h1V5h-1V4H5v1H4v12h1v1Z"/>`),
		g.Group(children),
	)
}

func PlusBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 20H4v-1H3v-1H2V4h1V3h1V2h14v1h1v1h1v14h-1v1h-1Zm-6-4v-4h4v-2h-4V6h-2v4H6v2h4v4Z"/>`),
		g.Group(children),
	)
}

func PlusCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1m-7 0h2v4h4v2h-4v4h-2v-4H6v-2h4V6Z"/>`),
		g.Group(children),
	)
}

func PlusCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 21H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-3-5v-4h4v-2h-4V6h-2v4H6v2h4v4Z"/>`),
		g.Group(children),
	)
}

func Radiobox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 2h6v1h2v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-2v1H8v-1H6v-1H5v-1H4v-1H3v-2H2V8h1V6h1V5h1V4h1V3h2V2m1 2v1H7v1H6v1H5v2H4v4h1v2h1v1h1v1h2v1h4v-1h2v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-2V4H9Z"/>`),
		g.Group(children),
	)
}

func RadioboxMarked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 2h6v1h2v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-2v1H8v-1H6v-1H5v-1H4v-1H3v-2H2V8h1V6h1V5h1V4h1V3h2V2m1 2v1H7v1H6v1H5v2H4v4h1v2h1v1h1v1h2v1h4v-1h2v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-2V4H9m4 3v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V9h1V8h1V7h4Z"/>`),
		g.Group(children),
	)
}

func RemoveCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M21 7v8h-1v2h-1v1h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1m-4-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1m-3 1v1h1v1h-1v1h-1v2h1v1h1v1h-1v1h-1v-1h-1v-1h-2v1H9v1H8v-1H7v-1h1v-1h1v-2H8V9H7V8h1V7h1v1h1v1h2V8h1V7h1Z"/>`),
		g.Group(children),
	)
}

func RotateClockwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M22 11v1h-1v1h-1v1h-1v1h-2v-1h-1v-1h-1v-1h-1v-1h3V9h-1V7h-1V6h-2V5H9v1H7v1H6v2H5v4h1v2h1v1h2v1h4v-1h3v2h-2v1H8v-1H6v-1H5v-1H4v-2H3V8h1V6h1V5h1V4h2V3h6v1h2v1h1v1h1v2h1v3h3Z"/>`),
		g.Group(children),
	)
}

func RotateCounterclockwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M0 11v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1H5V9h1V7h1V6h2V5h4v1h2v1h1v2h1v4h-1v2h-1v1h-2v1H9v-1H6v2h2v1h6v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v3H0Z"/>`),
		g.Group(children),
	)
}

func Script(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 1H5v1H4v13h2V3h9v16h-2v-1h-1v-1H1v3h1v1h14v-1h1V3h2v2h2V2h-1"/>`),
		g.Group(children),
	)
}

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 4h2V3h2V2h2V1h4v1h2v1h2v1h2v10h-1v2h-1v2h-1v1h-1v1h-2v1H9v-1H7v-1H6v-1H5v-2H4v-2H3V4m7-1v1H8v1H6v1H5v7h1v2h1v2h1v1h2v1h2v-1h2v-1h1v-2h1v-2h1V6h-1V5h-2V4h-2V3h-2Z"/>`),
		g.Group(children),
	)
}

func Skull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 2h2V1h6v1h2v1h1v1h1v1h1v2h1v7h-1v2h-1v4h-1v1H5v-1H4v-4H3v-2H2V8h1V5h1V4h1V3h1V2m9 3V4h-2V3H9v1H7v1H6v1H5v3H4v4h1v2h1v4h2v-2h2v2h2v-2h2v2h2v-4h1v-2h1V8h-1V6h-1V5h-1M7 8h3v3H7V8m5 3V8h3v3h-3m-2 2h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 1h14v1h1v18h-1v1H4v-1H3V2h1V1m1 2v16h12V3H5m4 2h1V4h2v1h1v2h-1v1h-2V7H9V5m0 13v-1H8v-1H7v-4h1v-1h1v-1h4v1h1v1h1v4h-1v1h-1v1H9m1-5H9v2h1v1h2v-1h1v-2h-1v-1h-2v1Z"/>`),
		g.Group(children),
	)
}

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 5v1h1v10h-1v1H6v-1H5V6h1V5h10M7 7v8h8V7H7Z"/>`),
		g.Group(children),
	)
}

func Sword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 3v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-1h2v2h-1v1h-1v1h1v1h1v1h1v2h-1v1h-2v-1h-1v-1h-1v-1h-1v1h-1v1h-2v-2h1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V8H3V7H2V2h5v1h1M7 5H6V4H4v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h-1v-1h-1v-1h-1V9h-1V8H9V7H8V6H7V5Z"/>`),
		g.Group(children),
	)
}

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 2h1V1h9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V2m2 8h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1v-1h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3H3v7m2-6h2v1h1v2H7v1H5V7H4V5h1V4Z"/>`),
		g.Group(children),
	)
}

func TagText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 2h1V1h9v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2v-1H1V2m2 8h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1v-1h1v-2h-1v-1h-1V9h-1V8h-1V7h-1V6h-1V5h-1V4h-1V3H3v7m11 1h1v2h-1v-1h-1v-1h-1v-1h-1V9h-1V7h1v1h1v1h1v1h1v1m-4 1h1v1h1v2h-1v-1h-1v-1H9v-1H8v-2h1v1h1v1M5 4h2v1h1v2H7v1H5V7H4V5h1V4Z"/>`),
		g.Group(children),
	)
}

func Target(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 13h-2v-1H9v-2h1V9h2v1h1v2h-1Zm2 4H8v-1H7v-1H6v-1H5V8h1V7h1V6h1V5h6v1h1v1h1v1h1v6h-1v1h-1v1h-1Zm1 4H7v-1H5v-1H4v-1H3v-1H2v-2H1V7h1V5h1V4h1V3h1V2h2V1h8v1h2v1h1v1h1v1h1v2h1v8h-1v2h-1v1h-1v1h-1v1h-2Zm-2-6v-1h1v-1h1V9h-1V8h-1V7H9v1H8v1H7v4h1v1h1v1Zm1 4v-1h2v-1h1v-1h1v-2h1V8h-1V6h-1V5h-1V4h-2V3H8v1H6v1H5v1H4v2H3v6h1v2h1v1h1v1h2v1Z"/>`),
		g.Group(children),
	)
}

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 3v16h-1v1H3v-1H2V3h1V2h16v1h1m-2 3H4v12h14V6M9 9v1h1v1h1v2h-1v1H9v1H8v1H6v-2h1v-1h1v-2H7v-1H6V8h2v1h1m2 7v-2h5v2h-5Z"/>`),
		g.Group(children),
	)
}

func TextBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m13 3V4H5v1H4v12h1v1h12v-1h1V5h-1M6 8h10v2H6V8m0 4h7v2H6v-2Z"/>`),
		g.Group(children),
	)
}

func TextImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 4h10v10H2V4m2 2v6h6V6H4m10-2h6v2h-6V4m0 4h6v2h-6V8m0 4h6v2h-6v-2M2 16h16v2H2v-2Z"/>`),
		g.Group(children),
	)
}

func ToggleSwitchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 8h4v1h1v4H9v1H5v-1H4V9h1V8m14-3v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1V5h16m-1 2H4v1H3v6h1v1h14v-1h1V8h-1V7Z"/>`),
		g.Group(children),
	)
}

func ToggleSwitchOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 5h16v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1V5m10 3v1h-1v4h1v1h4v-1h1V9h-1V8h-4Z"/>`),
		g.Group(children),
	)
}

func Toolbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 6h5V3h1V2h6v1h1v3h5v1h1v12h-1v1H2v-1H1V7h1V6m7 0h4V4H9v2m10 2H3v4h3v-2h3v2h4v-2h3v2h3V8M3 18h16v-4h-3v2h-3v-2H9v2H6v-2H3v4Z"/>`),
		g.Group(children),
	)
}

func TooltipAbove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3Z"/>`),
		g.Group(children),
	)
}

func TooltipAboveAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 5h2v5h-2V5m0 6h2v2h-2v-2M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3Z"/>`),
		g.Group(children),
	)
}

func TooltipAboveText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1h-5v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H2v-1H1V2h1V1m1 2v12h5v1h1v1h1v1h2v-1h1v-1h1v-1h5V3H3m2 3h12v2H5V6m0 4h10v2H5v-2Z"/>`),
		g.Group(children),
	)
}

func TooltipBelow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 21h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3Z"/>`),
		g.Group(children),
	)
}

func TooltipBelowAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 17h2v-2h-2v2m0-3h2V9h-2v5m-8 7h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3Z"/>`),
		g.Group(children),
	)
}

func TooltipBelowText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 21h18v-1h1V6h-1V5h-5V4h-1V3h-1V2h-1V1h-2v1H9v1H8v1H7v1H2v1H1v14h1v1m1-2V7h5V6h1V5h1V4h2v1h1v1h1v1h5v12H3m2-3h10v-2H5v2m0-4h12v-2H5v2Z"/>`),
		g.Group(children),
	)
}

func TooltipEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M21 20V2h-1V1H6v1H5v5H4v1H3v1H2v1H1v2h1v1h1v1h1v1h1v5h1v1h14v-1h1m-2-1H7v-5H6v-1H5v-1H4v-2h1V9h1V8h1V3h12v16Z"/>`),
		g.Group(children),
	)
}

func TooltipEndAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 15h-2v-2h2v2m0-3h-2V7h2v5m7-10v18h-1v1H6v-1H5v-5H4v-1H3v-1H2v-1H1v-2h1V9h1V8h1V7h1V2h1V1h14v1h1m-2 1H7v5H6v1H5v1H4v2h1v1h1v1h1v5h12V3Z"/>`),
		g.Group(children),
	)
}

func TooltipEndText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M17 8V6H9v2h8m-2 8v-2H9v2h6m2-4v-2H9v2h8m4 8V2h-1V1H6v1H5v5H4v1H3v1H2v1H1v2h1v1h1v1h1v1h1v5h1v1h14v-1h1m-2-1H7v-5H6v-1H5v-1H4v-2h1V9h1V8h1V3h12v16Z"/>`),
		g.Group(children),
	)
}

func TooltipStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 20V2h1V1h14v1h1v5h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v5h-1v1H2v-1H1m2-1h12v-5h1v-1h1v-1h1v-2h-1V9h-1V8h-1V3H3v16Z"/>`),
		g.Group(children),
	)
}

func TooltipStartAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 15h2v-2H8v2m0-3h2V7H8v5M1 2v18h1v1h14v-1h1v-5h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V2h-1V1H2v1H1m2 1h12v5h1v1h1v1h1v2h-1v1h-1v1h-1v5H3V3Z"/>`),
		g.Group(children),
	)
}

func TooltipStartText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 8V6h8v2H5m0 8v-2h6v2H5m0-4v-2h8v2H5m-4 8V2h1V1h14v1h1v5h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v5h-1v1H2v-1H1m2-1h12v-5h1v-1h1v-1h1v-2h-1V9h-1V8h-1V3H3v16Z"/>`),
		g.Group(children),
	)
}

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 7v9H8V7h2m2 0h2v9h-2V7M8 2h6v1h5v2h-1v14h-1v1H5v-1H4V5H3V3h5V2M6 5v13h10V5H6Z"/>`),
		g.Group(children),
	)
}

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M18 17v2H4v-2h14M8 15V9H4V8h1V7h1V6h1V5h1V4h1V3h1V2h2v1h1v1h1v1h1v1h1v1h1v1h1v1h-4v6H8m2-2h2V7h1V6h-1V5h-2v1H9v1h1v6Z"/>`),
		g.Group(children),
	)
}

func VolumeHigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M13 2h2v1h1v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-1v1h-2v-2h1v-1h1v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-1V4h-1V2m1 5v1h1v2h1v2h-1v2h-1v1h-1V7h1M2 8h4V7h1V6h1V5h1V4h1V3h1v16h-1v-1H9v-1H8v-1H7v-1H6v-1H2V8m2 2v2h3v1h1v1h1V8H8v1H7v1H4Z"/>`),
		g.Group(children),
	)
}

func VolumeLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 8h4V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1h-1v-1h-1v-1H6V8m2 2v2h3v1h1v1h1V8h-1v1h-1v1H8Z"/>`),
		g.Group(children),
	)
}

func VolumeMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 7v1h1v2h1v2h-1v2h-1v1h-1V7h1M8 8V7h1V6h1V5h1V4h1V3h1v16h-1v-1h-1v-1h-1v-1H9v-1H8v-1H4V8h4m-2 2v2h3v1h1v1h1V8h-1v1H9v1H6Z"/>`),
		g.Group(children),
	)
}

func VolumeMute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 7h2v1h1v1h2V8h1V7h2v2h-1v1h-1v2h1v1h1v2h-2v-1h-1v-1h-2v1h-1v1h-2v-2h1v-1h1v-2h-1V9h-1V7M6 8V7h1V6h1V5h1V4h1V3h1v16h-1v-1H9v-1H8v-1H7v-1H6v-1H2V8h4m1 2H4v2h3v1h1v1h1V8H8v1H7v1Z"/>`),
		g.Group(children),
	)
}

func Wall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 20H4v-5H1V7h3V2h15v5h2v8h-2ZM12 7V4H6v3Zm5 0V4h-3v3Zm-9 6V9H3v4Zm11 0V9h-9v4Zm-8 5v-3H6v3Zm6 0v-3h-4v3Z"/>`),
		g.Group(children),
	)
}

func WallFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 3h6v4H5m8-4h5v4h-5m-8 8v4h6v-4m2 0h5v4h-5M8 9H3v4h5m2-4h10v4H10"/>`),
		g.Group(children),
	)
}

func Water(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M14 21H8v-1H6v-1H5v-1H4v-2H3v-3h1v-2h1V9h1V7h1V6h1V4h1V3h1V1h2v2h1v1h1v2h1v1h1v2h1v2h1v2h1v3h-1v2h-1v1h-1v1h-2Z"/>`),
		g.Group(children),
	)
}

func Well(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 21H3v-5H1v-2h20v2h-2v5M5 16v3h12v-3H5M2 7V5h1V4h1V3h1V2h1V1h10v1h1v1h1v1h1v1h1v2h-2v6h-2V7h-4v2h2v4H8V9h2V7H6v6H4V7H2m5-4v1H6v1h10V4h-1V3H7Z"/>`),
		g.Group(children),
	)
}
