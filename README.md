Iup Go Wrapper
==============

iup.go is a [Go][1] wrapper around the [Iup][2] GUI toolkit. The project
was started on April 27, 2011.

Installing the Iup library
--------------------------

Iup.go deals with three projects from [Tecgraf][4]:

* Iup - Cross platform native control GUI library
* Cd - Cross platform canvas drawing library
* Im - Cross platform image library

To use iup.go you must install all three libraries. Download the appropriate archive files
from the [Iup Download Tips][3] page. You should then place the .a (library) files in
your lib/ directory and .h (header) files in your include/ directory. If compiling on 
Windows, .dll (dynamically linked library) files in your %PATH%.

Building Iup.go
---------------

    $ git clone https://github.com/jcowgar/iup.go.git
    $ cd iup.go
    $ ./all.bash
    
Special notes for Windows
-------------------------

You should download the 'dllw4' variant of Iup. This is the dynamically linked version of
the libraries for MinGW. For easy compliation you should have already installed a working
version of MinGW and MSYS. Building will be the same as any other platform but started from
the MSYS Bash shell.

Wrapper Status
==============

System
------

* **DONE** IupOpen
* **DONE** IupClose
* iuplua_open
* **DONE** IupVersion
* **DONE** IupLoad
* **DONE** IupLoadBuffer
* **DONE** IupSetLanguage
* **DONE** IupGetLanguage

Attributes
----------

* **DONE** IupStoreAttribute
* **DONE** IupStoreAttributeId
* IupSetAttribute
* IupSetAttributeId
* IupSetfAttribute
* IupSetfAttributeId
* IupSetfAttributeId2
* **DONE** IupSetAttributes
* IupResetAttribute
* IupSetAtt
* **DONE** IupSetAttributeHandle
* IupGetAttributeHandle
* **DONE** IupGetAttribute
* IupGetAttributeId
* IupGetAllAttributes
* IupGetAttributes
* **DONE** IupGetFloat
* **DONE** IupGetInt
* IupStoreGlobal
* IupSetGlobal
* IupGetGlobal

Events
------

* **DONE** IupMainLoop
* **DONE** IupMainLoopLevel
* **DONE** IupLoopStep
* **DONE** IupLoopStepWait
* **DONE** IupExitLoop
* **DONE** IupFlush
* IupGetCallback
* IupSetCallback
* IupSetCallbacks
* IupGetActionName
* IupGetFunction
* IupSetFunction
* IupRecordInput
* IupPlayInput

Layout/Construction
-------------------

* IupCreate
* IupCreatev
* IupCreatep
* **DONE** IupDestroy
* IupMap
* IupUnmap
* IupGetAllClasses
* IupGetClassName
* IupGetClassType
* IupClassMatch
* IupGetClassAttributes
* IupGetClassCallbacks
* IupSaveClassAttributes
* IupCopyClassAttributes
* IupSetClassDefaultAttribute

Layout/Composition
------------------

* IupFill
* **DONE** IupHbox
* **DONE** IupHboxv
* **DONE** IupVbox
* **DONE** IupVboxv
* IupZbox
* IupZboxv
* IupRadio
* IupNormalizer
* IupNormalizerv
* IupCbox
* IupCboxv
* IupSbox
* IupSplit

Layout/Hierarchy
----------------

* IupAppend
* IupDetach
* IupInsert
* IupReparent
* IupGetParent
* IupGetChild
* IupGetChildPos
* IupGetChildCount
* IupGetNextChild
* IupGetBrother
* IupGetDialog
* IupGetDialogChild

Layout/Utilities
----------------

* IupRefresh
* IupRefreshChildren
* IupUpdate
* IupUpdateChildren
* IupRedraw
* IupConvertXYToPos

Dialogs/Reference
-----------------

* **DONE** IupDialog
* **DONE** IupPopup
* **DONE** IupShow
* IupShowXY
* IupHide

Dialogs/Predefined
------------------

* **DONE** IupFileDlg
* IupMessageDlg
* IupColorDlg
* IupFontDlg
* IupAlarm
* IupGetFile
* IupGetColor
* IupGetParam
* IupGetText
* IupListDialog
* **DONE** IupMessage
* IupLayoutDialog
* IupElementPropertiesDialog

Controls/Standard
-----------------

* **DONE** IupButton
* **DONE** IupCanvas
* **DONE** IupFrame
* **DONE** IupLabel
* **DONE** IupList
* **DONE** IupProgressBar
* **DONE** IupSpin
* **DONE** IupTabs
* **DONE** IupTabsv
* **DONE** IupText
** IupTextConvertLinColToPos
** IupTextConvertPosToLinCol
* **DONE** IupToggle
* **DONE** IupTree
* **DONE** IupVal

Controls/Additional
-------------------

* **DONE** IupCells
* **DONE** IupColorbar
* **DONE** IupColorBrowser
* **DONE** IupDial
* **DONE** IupMatrix
** IupMatSetAttribute
** IupMatStoreAttribute
** IupMatGetAttribute
** IupMatGetInt
** IupMatGetFloat
** IupMatSetfAttribute
* **DONE** IupGLCanvas
* **DONE** IupPPlot
** IupPlotBegin
** IupPlotAdd
** IupPlotAddStr
** IupPlotEnd
** IupPlotInsert
** IupPlotInsertStr
** IupPlotInsertPoints
** IupPlotInsertStrPoints
** IupPlotAddPoints
** IupPlotAddStrPoints
** IupPlotTransform
** IupPlotPaintTo
* IupOleControl
* IupWebBrowser

Resources/Fonts
---------------

* IupMapFont
* IupUnMapFont

Resources/Images
----------------

* IupImage
* IupImageRGB
* IupImageRGBA
* IupImageLibOpen
* IupLoadImage
* IupSaveImage
* IupGetNativeHandleImage
* IupGetImageNativeHandle
* IupSaveImageAsText

Resources/Keyboard
------------------

* IupNextField
* IupPreviousField
* IupGetFocus
* IupSetFocus

Resources/Menus
---------------

* **DONE** IupItem
* **DONE** IupMenu
* **DONE** IupSeparator
* **DONE** IupSubmenu

Resources/Names
---------------

* IupSetHandle
* IupGetHandle
* IupGetName
* IupGetAllNames
* IupGetAllDialogs

Resources/Miscellaneous
-----------------------

* IupClipboard
* IupTimer
* IupTuioClient
* IupUser
* **DONE** IupHelp
* iupMaskSet
* iupMaskMatSet
* iupMaskSetInt
* iupMaskSetFloat
* iupMaskMatSetInt
* iupMaskMatSetFloat
* iupMaskRemove
* iupMaskMatRemove
* iupMaskCheck
* iupMaskMatCheck
* iupMaskGet
* iupMaskGetFloat
* iupMaskGetInt
* iupMaskMatGet
* iupMaskMatGetFloat
* iupMaskMatGetDouble
* iupMaskMatGetInt

[1]: http://golang.org                                       "Go"
[2]: http://www.tecgraf.puc-rio.br/iup/                      "Iup"
[3]: http://www.tecgraf.puc-rio.br/iup/en/download_tips.html "Iup Download Tips"
[4]: http://www.tecgraf.puc-rio.br/                          "Tecgraf"
