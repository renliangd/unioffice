//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package memstore implements tempStorage interface
// by using memory as a storage
package memstore ;import (_f "encoding/hex";_dg "errors";_de "fmt";_fb "github.com/unidoc/unioffice/common/tempstorage";_c "io";_b "io/ioutil";_cf "math/rand";_d "sync";);

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_da *memFile )Write (p []byte )(int ,error ){_da ._df ._fef =append (_da ._df ._fef ,p ...);_da ._df ._feg +=int64 (len (p ));return len (p ),nil ;};type memStorage struct{_ac _d .Map };func _eb (_eg string )string {_bc ,_ :=_efc (6);return _eg +_bc };

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage (){_be :=memStorage {_ac :_d .Map {}};_fb .SetAsStorage (&_be )};

// RemoveAll removes all files according to the dir argument prefix
func (_bf *memStorage )RemoveAll (dir string )error {_bf ._ac .Range (func (_bfd ,_gfd interface{})bool {_bf ._ac .Delete (_bfd );return true });return nil ;};

// Name returns the filename of the underlying memDataCell
func (_a *memFile )Name ()string {return _a ._df ._dfb };

// TempDir creates a name for a new temp directory using a pattern argument
func (_eed *memStorage )TempDir (pattern string )(string ,error ){return _eb (pattern ),nil };type memFile struct{_df *memDataCell ;_gf int64 ;};func _efc (_efd int )(string ,error ){_gfb :=make ([]byte ,_efd );if _ ,_gd :=_cf .Read (_gfb );_gd !=nil {return "",_gd ;};return _f .EncodeToString (_gfb ),nil ;};

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_bg *memFile )Read (p []byte )(int ,error ){_ga :=_bg ._gf ;_fe :=_bg ._df ._feg ;_cb :=int64 (len (p ));if _cb > _fe {_cb =_fe ;p =p [:_cb ];};if _ga >=_fe {return 0,_c .EOF ;};_e :=_ga +_cb ;if _e >=_fe {_e =_fe ;};_gfe :=copy (p ,_bg ._df ._fef [_ga :_e ]);_bg ._gf =_e ;return _gfe ,nil ;};

// Close is not applicable in this implementation
func (_ef *memFile )Close ()error {return nil };

// Add reads a file from a disk and adds it to the storage
func (_fd *memStorage )Add (path string )error {_ge ,_agb :=_b .ReadFile (path );if _agb !=nil {return _agb ;};_fd ._ac .Store (path ,&memDataCell {_dfb :path ,_fef :_ge });return nil ;};

// TempFile creates a new empty file in the storage and returns it
func (_cc *memStorage )TempFile (dir ,pattern string )(_fb .File ,error ){_gc :=dir +"\u002f"+_eb (pattern );_fege :=&memDataCell {_dfb :_gc ,_fef :[]byte {}};_beg :=&memFile {_df :_fege };_cc ._ac .Store (_gc ,_fege );return _beg ,nil ;};

// Open returns tempstorage File object by name
func (_cd *memStorage )Open (path string )(_fb .File ,error ){_ee ,_ag :=_cd ._ac .Load (path );if !_ag {return nil ,_dg .New (_de .Sprintf ("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",path ));};return &memFile {_df :_ee .(*memDataCell )},nil ;};type memDataCell struct{_dfb string ;_fef []byte ;_feg int64 ;};