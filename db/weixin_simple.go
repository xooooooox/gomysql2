package db

import (
	"database/sql"
	"fmt"
	"github.com/xooooooox/gomysql"
	"strings"
)

func (s *ImsAccount) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAccountInsertSql, s.Uniacid, s.Hash, s.Type, s.Isconnect, s.Isdeleted, s.Endtime, s.SendAccountExpireStatus, s.SendApiExpireStatus)
}

func (s *ImsAccount) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAccountDeleteSql, s.Acid)
}

func (s *ImsAccount) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_account` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccount) Get() (err error) {
	var tmp *ImsAccount
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountScan(rows)
		return
	}, ImsAccountSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccount) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAccountInsertSql, s.Uniacid, s.Hash, s.Type, s.Isconnect, s.Isdeleted, s.Endtime, s.SendAccountExpireStatus, s.SendApiExpireStatus)
}

func (s *ImsAccount) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAccountDeleteSql, s.Acid)
}

func (s *ImsAccount) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_account` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccount) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAccount
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountScan(rows)
		return
	}, ImsAccountSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountAliapp) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAccountAliappInsertSql, s.Uniacid, s.Level, s.Name, s.Description, s.Key)
}

func (s *ImsAccountAliapp) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAccountAliappDeleteSql, s.Acid)
}

func (s *ImsAccountAliapp) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_account_aliapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountAliapp) Get() (err error) {
	var tmp *ImsAccountAliapp
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountAliappScan(rows)
		return
	}, ImsAccountAliappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountAliapp) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAccountAliappInsertSql, s.Uniacid, s.Level, s.Name, s.Description, s.Key)
}

func (s *ImsAccountAliapp) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAccountAliappDeleteSql, s.Acid)
}

func (s *ImsAccountAliapp) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_account_aliapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountAliapp) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAccountAliapp
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountAliappScan(rows)
		return
	}, ImsAccountAliappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountBaiduapp) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAccountBaiduappInsertSql, s.Uniacid, s.Name, s.Appid, s.Key, s.Secret, s.Description)
}

func (s *ImsAccountBaiduapp) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAccountBaiduappDeleteSql, s.Acid)
}

func (s *ImsAccountBaiduapp) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_account_baiduapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountBaiduapp) Get() (err error) {
	var tmp *ImsAccountBaiduapp
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountBaiduappScan(rows)
		return
	}, ImsAccountBaiduappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountBaiduapp) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAccountBaiduappInsertSql, s.Uniacid, s.Name, s.Appid, s.Key, s.Secret, s.Description)
}

func (s *ImsAccountBaiduapp) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAccountBaiduappDeleteSql, s.Acid)
}

func (s *ImsAccountBaiduapp) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_account_baiduapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountBaiduapp) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAccountBaiduapp
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountBaiduappScan(rows)
		return
	}, ImsAccountBaiduappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountPhoneapp) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAccountPhoneappInsertSql, s.Uniacid, s.Name)
}

func (s *ImsAccountPhoneapp) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAccountPhoneappDeleteSql, s.Acid)
}

func (s *ImsAccountPhoneapp) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_account_phoneapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountPhoneapp) Get() (err error) {
	var tmp *ImsAccountPhoneapp
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountPhoneappScan(rows)
		return
	}, ImsAccountPhoneappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountPhoneapp) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAccountPhoneappInsertSql, s.Uniacid, s.Name)
}

func (s *ImsAccountPhoneapp) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAccountPhoneappDeleteSql, s.Acid)
}

func (s *ImsAccountPhoneapp) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_account_phoneapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountPhoneapp) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAccountPhoneapp
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountPhoneappScan(rows)
		return
	}, ImsAccountPhoneappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountToutiaoapp) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAccountToutiaoappInsertSql, s.Uniacid, s.Name, s.Appid, s.Key, s.Secret, s.Description)
}

func (s *ImsAccountToutiaoapp) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAccountToutiaoappDeleteSql, s.Acid)
}

func (s *ImsAccountToutiaoapp) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_account_toutiaoapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountToutiaoapp) Get() (err error) {
	var tmp *ImsAccountToutiaoapp
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountToutiaoappScan(rows)
		return
	}, ImsAccountToutiaoappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountToutiaoapp) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAccountToutiaoappInsertSql, s.Uniacid, s.Name, s.Appid, s.Key, s.Secret, s.Description)
}

func (s *ImsAccountToutiaoapp) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAccountToutiaoappDeleteSql, s.Acid)
}

func (s *ImsAccountToutiaoapp) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_account_toutiaoapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountToutiaoapp) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAccountToutiaoapp
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountToutiaoappScan(rows)
		return
	}, ImsAccountToutiaoappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountWebapp) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAccountWebappInsertSql, s.Uniacid, s.Name)
}

func (s *ImsAccountWebapp) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAccountWebappDeleteSql, s.Acid)
}

func (s *ImsAccountWebapp) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_account_webapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountWebapp) Get() (err error) {
	var tmp *ImsAccountWebapp
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountWebappScan(rows)
		return
	}, ImsAccountWebappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountWebapp) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAccountWebappInsertSql, s.Uniacid, s.Name)
}

func (s *ImsAccountWebapp) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAccountWebappDeleteSql, s.Acid)
}

func (s *ImsAccountWebapp) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_account_webapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountWebapp) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAccountWebapp
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountWebappScan(rows)
		return
	}, ImsAccountWebappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountWechats) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAccountWechatsInsertSql, s.Uniacid, s.Token, s.Encodingaeskey, s.Level, s.Name, s.Account, s.Original, s.Signature, s.Country, s.Province, s.City, s.Username, s.Password, s.Lastupdate, s.Key, s.Secret, s.Styleid, s.Subscribeurl, s.AuthRefreshToken)
}

func (s *ImsAccountWechats) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAccountWechatsDeleteSql, s.Acid)
}

func (s *ImsAccountWechats) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_account_wechats` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountWechats) Get() (err error) {
	var tmp *ImsAccountWechats
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountWechatsScan(rows)
		return
	}, ImsAccountWechatsSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountWechats) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAccountWechatsInsertSql, s.Uniacid, s.Token, s.Encodingaeskey, s.Level, s.Name, s.Account, s.Original, s.Signature, s.Country, s.Province, s.City, s.Username, s.Password, s.Lastupdate, s.Key, s.Secret, s.Styleid, s.Subscribeurl, s.AuthRefreshToken)
}

func (s *ImsAccountWechats) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAccountWechatsDeleteSql, s.Acid)
}

func (s *ImsAccountWechats) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_account_wechats` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountWechats) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAccountWechats
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountWechatsScan(rows)
		return
	}, ImsAccountWechatsSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountWxapp) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAccountWxappInsertSql, s.Uniacid, s.Token, s.Encodingaeskey, s.Level, s.Account, s.Original, s.Key, s.Secret, s.Name, s.Appdomain, s.AuthRefreshToken)
}

func (s *ImsAccountWxapp) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAccountWxappDeleteSql, s.Acid)
}

func (s *ImsAccountWxapp) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_account_wxapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountWxapp) Get() (err error) {
	var tmp *ImsAccountWxapp
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountWxappScan(rows)
		return
	}, ImsAccountWxappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountWxapp) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAccountWxappInsertSql, s.Uniacid, s.Token, s.Encodingaeskey, s.Level, s.Account, s.Original, s.Key, s.Secret, s.Name, s.Appdomain, s.AuthRefreshToken)
}

func (s *ImsAccountWxapp) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAccountWxappDeleteSql, s.Acid)
}

func (s *ImsAccountWxapp) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_account_wxapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountWxapp) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAccountWxapp
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountWxappScan(rows)
		return
	}, ImsAccountWxappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountXzapp) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAccountXzappInsertSql, s.Uniacid, s.Name, s.Original, s.Lastupdate, s.Styleid, s.Createtime, s.Token, s.Encodingaeskey, s.XzappId, s.Level, s.Key, s.Secret)
}

func (s *ImsAccountXzapp) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAccountXzappDeleteSql, s.Acid)
}

func (s *ImsAccountXzapp) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_account_xzapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountXzapp) Get() (err error) {
	var tmp *ImsAccountXzapp
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountXzappScan(rows)
		return
	}, ImsAccountXzappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAccountXzapp) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAccountXzappInsertSql, s.Uniacid, s.Name, s.Original, s.Lastupdate, s.Styleid, s.Createtime, s.Token, s.Encodingaeskey, s.XzappId, s.Level, s.Key, s.Secret)
}

func (s *ImsAccountXzapp) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAccountXzappDeleteSql, s.Acid)
}

func (s *ImsAccountXzapp) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Acid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_account_xzapp` SET %s WHERE ( `acid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAccountXzapp) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAccountXzapp
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAccountXzappScan(rows)
		return
	}, ImsAccountXzappSelectSql, s.Acid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsActivityClerks) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsActivityClerksInsertSql, s.Uniacid, s.Uid, s.Storeid, s.Name, s.Password, s.Mobile, s.Openid, s.Nickname)
}

func (s *ImsActivityClerks) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsActivityClerksDeleteSql, s.Id)
}

func (s *ImsActivityClerks) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_activity_clerks` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsActivityClerks) Get() (err error) {
	var tmp *ImsActivityClerks
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsActivityClerksScan(rows)
		return
	}, ImsActivityClerksSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsActivityClerks) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsActivityClerksInsertSql, s.Uniacid, s.Uid, s.Storeid, s.Name, s.Password, s.Mobile, s.Openid, s.Nickname)
}

func (s *ImsActivityClerks) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsActivityClerksDeleteSql, s.Id)
}

func (s *ImsActivityClerks) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_activity_clerks` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsActivityClerks) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsActivityClerks
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsActivityClerksScan(rows)
		return
	}, ImsActivityClerksSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsActivityClerkMenu) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsActivityClerkMenuInsertSql, s.Uniacid, s.Displayorder, s.Pid, s.GroupName, s.Title, s.Icon, s.Url, s.Type, s.Permission, s.System)
}

func (s *ImsActivityClerkMenu) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsActivityClerkMenuDeleteSql, s.Id)
}

func (s *ImsActivityClerkMenu) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_activity_clerk_menu` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsActivityClerkMenu) Get() (err error) {
	var tmp *ImsActivityClerkMenu
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsActivityClerkMenuScan(rows)
		return
	}, ImsActivityClerkMenuSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsActivityClerkMenu) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsActivityClerkMenuInsertSql, s.Uniacid, s.Displayorder, s.Pid, s.GroupName, s.Title, s.Icon, s.Url, s.Type, s.Permission, s.System)
}

func (s *ImsActivityClerkMenu) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsActivityClerkMenuDeleteSql, s.Id)
}

func (s *ImsActivityClerkMenu) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_activity_clerk_menu` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsActivityClerkMenu) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsActivityClerkMenu
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsActivityClerkMenuScan(rows)
		return
	}, ImsActivityClerkMenuSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAddress) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAddressInsertSql, s.Pid, s.Name, s.Level)
}

func (s *ImsAddress) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAddressDeleteSql, s.Id)
}

func (s *ImsAddress) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_address` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAddress) Get() (err error) {
	var tmp *ImsAddress
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAddressScan(rows)
		return
	}, ImsAddressSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAddress) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAddressInsertSql, s.Pid, s.Name, s.Level)
}

func (s *ImsAddress) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAddressDeleteSql, s.Id)
}

func (s *ImsAddress) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_address` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAddress) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAddress
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAddressScan(rows)
		return
	}, ImsAddressSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleCategory) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsArticleCategoryInsertSql, s.Title, s.Displayorder, s.Type)
}

func (s *ImsArticleCategory) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsArticleCategoryDeleteSql, s.Id)
}

func (s *ImsArticleCategory) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_article_category` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleCategory) Get() (err error) {
	var tmp *ImsArticleCategory
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleCategoryScan(rows)
		return
	}, ImsArticleCategorySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleCategory) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsArticleCategoryInsertSql, s.Title, s.Displayorder, s.Type)
}

func (s *ImsArticleCategory) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsArticleCategoryDeleteSql, s.Id)
}

func (s *ImsArticleCategory) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_article_category` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleCategory) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsArticleCategory
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleCategoryScan(rows)
		return
	}, ImsArticleCategorySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleComment) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsArticleCommentInsertSql, s.Articleid, s.Parentid, s.Uid, s.Content, s.IsLike, s.IsReply, s.LikeNum, s.Createtime)
}

func (s *ImsArticleComment) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsArticleCommentDeleteSql, s.Id)
}

func (s *ImsArticleComment) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_article_comment` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleComment) Get() (err error) {
	var tmp *ImsArticleComment
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleCommentScan(rows)
		return
	}, ImsArticleCommentSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleComment) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsArticleCommentInsertSql, s.Articleid, s.Parentid, s.Uid, s.Content, s.IsLike, s.IsReply, s.LikeNum, s.Createtime)
}

func (s *ImsArticleComment) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsArticleCommentDeleteSql, s.Id)
}

func (s *ImsArticleComment) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_article_comment` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleComment) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsArticleComment
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleCommentScan(rows)
		return
	}, ImsArticleCommentSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleNews) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsArticleNewsInsertSql, s.Cateid, s.Title, s.Content, s.Thumb, s.Source, s.Author, s.Displayorder, s.IsDisplay, s.IsShowHome, s.Createtime, s.Click)
}

func (s *ImsArticleNews) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsArticleNewsDeleteSql, s.Id)
}

func (s *ImsArticleNews) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_article_news` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleNews) Get() (err error) {
	var tmp *ImsArticleNews
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleNewsScan(rows)
		return
	}, ImsArticleNewsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleNews) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsArticleNewsInsertSql, s.Cateid, s.Title, s.Content, s.Thumb, s.Source, s.Author, s.Displayorder, s.IsDisplay, s.IsShowHome, s.Createtime, s.Click)
}

func (s *ImsArticleNews) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsArticleNewsDeleteSql, s.Id)
}

func (s *ImsArticleNews) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_article_news` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleNews) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsArticleNews
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleNewsScan(rows)
		return
	}, ImsArticleNewsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleNotice) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsArticleNoticeInsertSql, s.Cateid, s.Title, s.Content, s.Displayorder, s.IsDisplay, s.IsShowHome, s.Createtime, s.Click, s.Style, s.Group)
}

func (s *ImsArticleNotice) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsArticleNoticeDeleteSql, s.Id)
}

func (s *ImsArticleNotice) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_article_notice` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleNotice) Get() (err error) {
	var tmp *ImsArticleNotice
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleNoticeScan(rows)
		return
	}, ImsArticleNoticeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleNotice) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsArticleNoticeInsertSql, s.Cateid, s.Title, s.Content, s.Displayorder, s.IsDisplay, s.IsShowHome, s.Createtime, s.Click, s.Style, s.Group)
}

func (s *ImsArticleNotice) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsArticleNoticeDeleteSql, s.Id)
}

func (s *ImsArticleNotice) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_article_notice` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleNotice) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsArticleNotice
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleNoticeScan(rows)
		return
	}, ImsArticleNoticeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleUnreadNotice) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsArticleUnreadNoticeInsertSql, s.NoticeId, s.Uid, s.IsNew)
}

func (s *ImsArticleUnreadNotice) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsArticleUnreadNoticeDeleteSql, s.Id)
}

func (s *ImsArticleUnreadNotice) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_article_unread_notice` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleUnreadNotice) Get() (err error) {
	var tmp *ImsArticleUnreadNotice
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleUnreadNoticeScan(rows)
		return
	}, ImsArticleUnreadNoticeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsArticleUnreadNotice) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsArticleUnreadNoticeInsertSql, s.NoticeId, s.Uid, s.IsNew)
}

func (s *ImsArticleUnreadNotice) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsArticleUnreadNoticeDeleteSql, s.Id)
}

func (s *ImsArticleUnreadNotice) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_article_unread_notice` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsArticleUnreadNotice) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsArticleUnreadNotice
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsArticleUnreadNoticeScan(rows)
		return
	}, ImsArticleUnreadNoticeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAttachmentGroup) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsAttachmentGroupInsertSql, s.Pid, s.Name, s.Uniacid, s.Uid, s.Type)
}

func (s *ImsAttachmentGroup) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsAttachmentGroupDeleteSql, s.Id)
}

func (s *ImsAttachmentGroup) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_attachment_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAttachmentGroup) Get() (err error) {
	var tmp *ImsAttachmentGroup
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAttachmentGroupScan(rows)
		return
	}, ImsAttachmentGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsAttachmentGroup) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsAttachmentGroupInsertSql, s.Pid, s.Name, s.Uniacid, s.Uid, s.Type)
}

func (s *ImsAttachmentGroup) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsAttachmentGroupDeleteSql, s.Id)
}

func (s *ImsAttachmentGroup) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_attachment_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsAttachmentGroup) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsAttachmentGroup
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsAttachmentGroupScan(rows)
		return
	}, ImsAttachmentGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsBasicReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsBasicReplyInsertSql, s.Rid, s.Content)
}

func (s *ImsBasicReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsBasicReplyDeleteSql, s.Id)
}

func (s *ImsBasicReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_basic_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsBasicReply) Get() (err error) {
	var tmp *ImsBasicReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsBasicReplyScan(rows)
		return
	}, ImsBasicReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsBasicReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsBasicReplyInsertSql, s.Rid, s.Content)
}

func (s *ImsBasicReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsBasicReplyDeleteSql, s.Id)
}

func (s *ImsBasicReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_basic_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsBasicReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsBasicReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsBasicReplyScan(rows)
		return
	}, ImsBasicReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsBusiness) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsBusinessInsertSql, s.Weid, s.Title, s.Thumb, s.Content, s.Phone, s.Qq, s.Province, s.City, s.Dist, s.Address, s.Lng, s.Lat, s.Industry1, s.Industry2, s.Createtime)
}

func (s *ImsBusiness) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsBusinessDeleteSql, s.Id)
}

func (s *ImsBusiness) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_business` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsBusiness) Get() (err error) {
	var tmp *ImsBusiness
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsBusinessScan(rows)
		return
	}, ImsBusinessSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsBusiness) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsBusinessInsertSql, s.Weid, s.Title, s.Thumb, s.Content, s.Phone, s.Qq, s.Province, s.City, s.Dist, s.Address, s.Lng, s.Lat, s.Industry1, s.Industry2, s.Createtime)
}

func (s *ImsBusiness) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsBusinessDeleteSql, s.Id)
}

func (s *ImsBusiness) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_business` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsBusiness) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsBusiness
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsBusinessScan(rows)
		return
	}, ImsBusinessSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreAttachment) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreAttachmentInsertSql, s.Uniacid, s.Uid, s.Filename, s.Attachment, s.Type, s.Createtime, s.ModuleUploadDir, s.GroupId, s.Displayorder)
}

func (s *ImsCoreAttachment) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreAttachmentDeleteSql, s.Id)
}

func (s *ImsCoreAttachment) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_attachment` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreAttachment) Get() (err error) {
	var tmp *ImsCoreAttachment
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreAttachmentScan(rows)
		return
	}, ImsCoreAttachmentSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreAttachment) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreAttachmentInsertSql, s.Uniacid, s.Uid, s.Filename, s.Attachment, s.Type, s.Createtime, s.ModuleUploadDir, s.GroupId, s.Displayorder)
}

func (s *ImsCoreAttachment) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreAttachmentDeleteSql, s.Id)
}

func (s *ImsCoreAttachment) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_attachment` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreAttachment) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreAttachment
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreAttachmentScan(rows)
		return
	}, ImsCoreAttachmentSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreCache) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreCacheInsertSql, s.Value)
}

func (s *ImsCoreCache) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreCacheDeleteSql, s.Key)
}

func (s *ImsCoreCache) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Key
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_cache` SET %s WHERE ( `key` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreCache) Get() (err error) {
	var tmp *ImsCoreCache
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreCacheScan(rows)
		return
	}, ImsCoreCacheSelectSql, s.Key)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreCache) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreCacheInsertSql, s.Value)
}

func (s *ImsCoreCache) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreCacheDeleteSql, s.Key)
}

func (s *ImsCoreCache) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Key
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_cache` SET %s WHERE ( `key` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreCache) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreCache
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreCacheScan(rows)
		return
	}, ImsCoreCacheSelectSql, s.Key)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreCron) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreCronInsertSql, s.Cloudid, s.Module, s.Uniacid, s.Type, s.Name, s.Filename, s.Lastruntime, s.Nextruntime, s.Weekday, s.Day, s.Hour, s.Minute, s.Extra, s.Status, s.Createtime)
}

func (s *ImsCoreCron) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreCronDeleteSql, s.Id)
}

func (s *ImsCoreCron) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_cron` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreCron) Get() (err error) {
	var tmp *ImsCoreCron
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreCronScan(rows)
		return
	}, ImsCoreCronSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreCron) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreCronInsertSql, s.Cloudid, s.Module, s.Uniacid, s.Type, s.Name, s.Filename, s.Lastruntime, s.Nextruntime, s.Weekday, s.Day, s.Hour, s.Minute, s.Extra, s.Status, s.Createtime)
}

func (s *ImsCoreCron) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreCronDeleteSql, s.Id)
}

func (s *ImsCoreCron) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_cron` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreCron) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreCron
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreCronScan(rows)
		return
	}, ImsCoreCronSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreCronRecord) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreCronRecordInsertSql, s.Uniacid, s.Module, s.Type, s.Tid, s.Note, s.Tag, s.Createtime)
}

func (s *ImsCoreCronRecord) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreCronRecordDeleteSql, s.Id)
}

func (s *ImsCoreCronRecord) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_cron_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreCronRecord) Get() (err error) {
	var tmp *ImsCoreCronRecord
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreCronRecordScan(rows)
		return
	}, ImsCoreCronRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreCronRecord) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreCronRecordInsertSql, s.Uniacid, s.Module, s.Type, s.Tid, s.Note, s.Tag, s.Createtime)
}

func (s *ImsCoreCronRecord) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreCronRecordDeleteSql, s.Id)
}

func (s *ImsCoreCronRecord) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_cron_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreCronRecord) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreCronRecord
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreCronRecordScan(rows)
		return
	}, ImsCoreCronRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreJob) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreJobInsertSql, s.Type, s.Uniacid, s.Payload, s.Status, s.Title, s.Handled, s.Total, s.Createtime, s.Updatetime, s.Endtime, s.Uid, s.Isdeleted)
}

func (s *ImsCoreJob) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreJobDeleteSql, s.Id)
}

func (s *ImsCoreJob) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_job` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreJob) Get() (err error) {
	var tmp *ImsCoreJob
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreJobScan(rows)
		return
	}, ImsCoreJobSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreJob) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreJobInsertSql, s.Type, s.Uniacid, s.Payload, s.Status, s.Title, s.Handled, s.Total, s.Createtime, s.Updatetime, s.Endtime, s.Uid, s.Isdeleted)
}

func (s *ImsCoreJob) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreJobDeleteSql, s.Id)
}

func (s *ImsCoreJob) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_job` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreJob) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreJob
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreJobScan(rows)
		return
	}, ImsCoreJobSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreMenu) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreMenuInsertSql, s.Pid, s.Title, s.Name, s.Url, s.AppendTitle, s.AppendUrl, s.Displayorder, s.Type, s.IsDisplay, s.IsSystem, s.PermissionName, s.GroupName, s.Icon)
}

func (s *ImsCoreMenu) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreMenuDeleteSql, s.Id)
}

func (s *ImsCoreMenu) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_menu` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreMenu) Get() (err error) {
	var tmp *ImsCoreMenu
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreMenuScan(rows)
		return
	}, ImsCoreMenuSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreMenu) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreMenuInsertSql, s.Pid, s.Title, s.Name, s.Url, s.AppendTitle, s.AppendUrl, s.Displayorder, s.Type, s.IsDisplay, s.IsSystem, s.PermissionName, s.GroupName, s.Icon)
}

func (s *ImsCoreMenu) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreMenuDeleteSql, s.Id)
}

func (s *ImsCoreMenu) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_menu` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreMenu) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreMenu
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreMenuScan(rows)
		return
	}, ImsCoreMenuSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreMenuShortcut) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreMenuShortcutInsertSql, s.Uid, s.Uniacid, s.Modulename, s.Displayorder, s.Position, s.Updatetime)
}

func (s *ImsCoreMenuShortcut) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreMenuShortcutDeleteSql, s.Id)
}

func (s *ImsCoreMenuShortcut) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_menu_shortcut` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreMenuShortcut) Get() (err error) {
	var tmp *ImsCoreMenuShortcut
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreMenuShortcutScan(rows)
		return
	}, ImsCoreMenuShortcutSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreMenuShortcut) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreMenuShortcutInsertSql, s.Uid, s.Uniacid, s.Modulename, s.Displayorder, s.Position, s.Updatetime)
}

func (s *ImsCoreMenuShortcut) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreMenuShortcutDeleteSql, s.Id)
}

func (s *ImsCoreMenuShortcut) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_menu_shortcut` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreMenuShortcut) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreMenuShortcut
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreMenuShortcutScan(rows)
		return
	}, ImsCoreMenuShortcutSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCorePaylog) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCorePaylogInsertSql, s.Type, s.Uniacid, s.Acid, s.Openid, s.Uniontid, s.Tid, s.Fee, s.Status, s.Module, s.Tag, s.IsUsecard, s.CardType, s.CardId, s.CardFee, s.EncryptCode, s.IsWish)
}

func (s *ImsCorePaylog) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCorePaylogDeleteSql, s.Plid)
}

func (s *ImsCorePaylog) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Plid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_paylog` SET %s WHERE ( `plid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCorePaylog) Get() (err error) {
	var tmp *ImsCorePaylog
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCorePaylogScan(rows)
		return
	}, ImsCorePaylogSelectSql, s.Plid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCorePaylog) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCorePaylogInsertSql, s.Type, s.Uniacid, s.Acid, s.Openid, s.Uniontid, s.Tid, s.Fee, s.Status, s.Module, s.Tag, s.IsUsecard, s.CardType, s.CardId, s.CardFee, s.EncryptCode, s.IsWish)
}

func (s *ImsCorePaylog) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCorePaylogDeleteSql, s.Plid)
}

func (s *ImsCorePaylog) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Plid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_paylog` SET %s WHERE ( `plid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCorePaylog) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCorePaylog
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCorePaylogScan(rows)
		return
	}, ImsCorePaylogSelectSql, s.Plid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCorePerformance) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCorePerformanceInsertSql, s.Type, s.Runtime, s.Runurl, s.Runsql, s.Createtime)
}

func (s *ImsCorePerformance) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCorePerformanceDeleteSql, s.Id)
}

func (s *ImsCorePerformance) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_performance` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCorePerformance) Get() (err error) {
	var tmp *ImsCorePerformance
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCorePerformanceScan(rows)
		return
	}, ImsCorePerformanceSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCorePerformance) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCorePerformanceInsertSql, s.Type, s.Runtime, s.Runurl, s.Runsql, s.Createtime)
}

func (s *ImsCorePerformance) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCorePerformanceDeleteSql, s.Id)
}

func (s *ImsCorePerformance) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_performance` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCorePerformance) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCorePerformance
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCorePerformanceScan(rows)
		return
	}, ImsCorePerformanceSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreQueue) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreQueueInsertSql, s.Uniacid, s.Acid, s.Message, s.Params, s.Keyword, s.Response, s.Module, s.Type, s.Dateline)
}

func (s *ImsCoreQueue) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreQueueDeleteSql, s.Qid)
}

func (s *ImsCoreQueue) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Qid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_queue` SET %s WHERE ( `qid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreQueue) Get() (err error) {
	var tmp *ImsCoreQueue
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreQueueScan(rows)
		return
	}, ImsCoreQueueSelectSql, s.Qid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreQueue) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreQueueInsertSql, s.Uniacid, s.Acid, s.Message, s.Params, s.Keyword, s.Response, s.Module, s.Type, s.Dateline)
}

func (s *ImsCoreQueue) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreQueueDeleteSql, s.Qid)
}

func (s *ImsCoreQueue) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Qid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_queue` SET %s WHERE ( `qid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreQueue) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreQueue
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreQueueScan(rows)
		return
	}, ImsCoreQueueSelectSql, s.Qid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreRefundlog) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreRefundlogInsertSql, s.Uniacid, s.RefundUniontid, s.Reason, s.Uniontid, s.Fee, s.Status, s.IsWish)
}

func (s *ImsCoreRefundlog) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreRefundlogDeleteSql, s.Id)
}

func (s *ImsCoreRefundlog) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_refundlog` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreRefundlog) Get() (err error) {
	var tmp *ImsCoreRefundlog
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreRefundlogScan(rows)
		return
	}, ImsCoreRefundlogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreRefundlog) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreRefundlogInsertSql, s.Uniacid, s.RefundUniontid, s.Reason, s.Uniontid, s.Fee, s.Status, s.IsWish)
}

func (s *ImsCoreRefundlog) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreRefundlogDeleteSql, s.Id)
}

func (s *ImsCoreRefundlog) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_refundlog` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreRefundlog) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreRefundlog
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreRefundlogScan(rows)
		return
	}, ImsCoreRefundlogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreResource) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreResourceInsertSql, s.Uniacid, s.MediaId, s.Trunk, s.Type, s.Dateline)
}

func (s *ImsCoreResource) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreResourceDeleteSql, s.Mid)
}

func (s *ImsCoreResource) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Mid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_resource` SET %s WHERE ( `mid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreResource) Get() (err error) {
	var tmp *ImsCoreResource
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreResourceScan(rows)
		return
	}, ImsCoreResourceSelectSql, s.Mid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreResource) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreResourceInsertSql, s.Uniacid, s.MediaId, s.Trunk, s.Type, s.Dateline)
}

func (s *ImsCoreResource) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreResourceDeleteSql, s.Mid)
}

func (s *ImsCoreResource) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Mid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_resource` SET %s WHERE ( `mid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreResource) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreResource
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreResourceScan(rows)
		return
	}, ImsCoreResourceSelectSql, s.Mid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreSendsmsLog) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreSendsmsLogInsertSql, s.Uniacid, s.Mobile, s.Content, s.Result, s.Createtime)
}

func (s *ImsCoreSendsmsLog) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreSendsmsLogDeleteSql, s.Id)
}

func (s *ImsCoreSendsmsLog) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_sendsms_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreSendsmsLog) Get() (err error) {
	var tmp *ImsCoreSendsmsLog
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreSendsmsLogScan(rows)
		return
	}, ImsCoreSendsmsLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreSendsmsLog) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreSendsmsLogInsertSql, s.Uniacid, s.Mobile, s.Content, s.Result, s.Createtime)
}

func (s *ImsCoreSendsmsLog) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreSendsmsLogDeleteSql, s.Id)
}

func (s *ImsCoreSendsmsLog) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_sendsms_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreSendsmsLog) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreSendsmsLog
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreSendsmsLogScan(rows)
		return
	}, ImsCoreSendsmsLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreSessions) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreSessionsInsertSql, s.Uniacid, s.Openid, s.Data, s.Expiretime)
}

func (s *ImsCoreSessions) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreSessionsDeleteSql, s.Sid)
}

func (s *ImsCoreSessions) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Sid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_sessions` SET %s WHERE ( `sid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreSessions) Get() (err error) {
	var tmp *ImsCoreSessions
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreSessionsScan(rows)
		return
	}, ImsCoreSessionsSelectSql, s.Sid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreSessions) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreSessionsInsertSql, s.Uniacid, s.Openid, s.Data, s.Expiretime)
}

func (s *ImsCoreSessions) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreSessionsDeleteSql, s.Sid)
}

func (s *ImsCoreSessions) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Sid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_sessions` SET %s WHERE ( `sid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreSessions) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreSessions
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreSessionsScan(rows)
		return
	}, ImsCoreSessionsSelectSql, s.Sid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreSettings) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoreSettingsInsertSql, s.Value)
}

func (s *ImsCoreSettings) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoreSettingsDeleteSql, s.Key)
}

func (s *ImsCoreSettings) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Key
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_core_settings` SET %s WHERE ( `key` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreSettings) Get() (err error) {
	var tmp *ImsCoreSettings
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreSettingsScan(rows)
		return
	}, ImsCoreSettingsSelectSql, s.Key)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoreSettings) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoreSettingsInsertSql, s.Value)
}

func (s *ImsCoreSettings) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoreSettingsDeleteSql, s.Key)
}

func (s *ImsCoreSettings) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Key
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_core_settings` SET %s WHERE ( `key` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoreSettings) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoreSettings
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoreSettingsScan(rows)
		return
	}, ImsCoreSettingsSelectSql, s.Key)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCouponLocation) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCouponLocationInsertSql, s.Uniacid, s.Acid, s.Sid, s.LocationId, s.BusinessName, s.BranchName, s.Category, s.Province, s.City, s.District, s.Address, s.Longitude, s.Latitude, s.Telephone, s.PhotoList, s.AvgPrice, s.OpenTime, s.Recommend, s.Special, s.Introduction, s.OffsetType, s.Status, s.Message)
}

func (s *ImsCouponLocation) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCouponLocationDeleteSql, s.Id)
}

func (s *ImsCouponLocation) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_coupon_location` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCouponLocation) Get() (err error) {
	var tmp *ImsCouponLocation
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCouponLocationScan(rows)
		return
	}, ImsCouponLocationSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCouponLocation) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCouponLocationInsertSql, s.Uniacid, s.Acid, s.Sid, s.LocationId, s.BusinessName, s.BranchName, s.Category, s.Province, s.City, s.District, s.Address, s.Longitude, s.Latitude, s.Telephone, s.PhotoList, s.AvgPrice, s.OpenTime, s.Recommend, s.Special, s.Introduction, s.OffsetType, s.Status, s.Message)
}

func (s *ImsCouponLocation) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCouponLocationDeleteSql, s.Id)
}

func (s *ImsCouponLocation) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_coupon_location` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCouponLocation) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCouponLocation
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCouponLocationScan(rows)
		return
	}, ImsCouponLocationSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoverReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCoverReplyInsertSql, s.Uniacid, s.Multiid, s.Rid, s.Module, s.Do, s.Title, s.Description, s.Thumb, s.Url)
}

func (s *ImsCoverReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCoverReplyDeleteSql, s.Id)
}

func (s *ImsCoverReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_cover_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoverReply) Get() (err error) {
	var tmp *ImsCoverReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoverReplyScan(rows)
		return
	}, ImsCoverReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCoverReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCoverReplyInsertSql, s.Uniacid, s.Multiid, s.Rid, s.Module, s.Do, s.Title, s.Description, s.Thumb, s.Url)
}

func (s *ImsCoverReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCoverReplyDeleteSql, s.Id)
}

func (s *ImsCoverReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_cover_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCoverReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCoverReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCoverReplyScan(rows)
		return
	}, ImsCoverReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCustomReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsCustomReplyInsertSql, s.Rid, s.Start1, s.End1, s.Start2, s.End2)
}

func (s *ImsCustomReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsCustomReplyDeleteSql, s.Id)
}

func (s *ImsCustomReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_custom_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCustomReply) Get() (err error) {
	var tmp *ImsCustomReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCustomReplyScan(rows)
		return
	}, ImsCustomReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsCustomReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsCustomReplyInsertSql, s.Rid, s.Start1, s.End1, s.Start2, s.End2)
}

func (s *ImsCustomReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsCustomReplyDeleteSql, s.Id)
}

func (s *ImsCustomReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_custom_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsCustomReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsCustomReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsCustomReplyScan(rows)
		return
	}, ImsCustomReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsImagesReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsImagesReplyInsertSql, s.Rid, s.Title, s.Description, s.Mediaid, s.Createtime)
}

func (s *ImsImagesReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsImagesReplyDeleteSql, s.Id)
}

func (s *ImsImagesReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_images_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsImagesReply) Get() (err error) {
	var tmp *ImsImagesReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsImagesReplyScan(rows)
		return
	}, ImsImagesReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsImagesReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsImagesReplyInsertSql, s.Rid, s.Title, s.Description, s.Mediaid, s.Createtime)
}

func (s *ImsImagesReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsImagesReplyDeleteSql, s.Id)
}

func (s *ImsImagesReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_images_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsImagesReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsImagesReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsImagesReplyScan(rows)
		return
	}, ImsImagesReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcCashRecord) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcCashRecordInsertSql, s.Uniacid, s.Uid, s.ClerkId, s.StoreId, s.ClerkType, s.Fee, s.FinalFee, s.Credit1, s.Credit1Fee, s.Credit2, s.Cash, s.ReturnCash, s.FinalCash, s.Remark, s.Createtime, s.TradeType)
}

func (s *ImsMcCashRecord) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcCashRecordDeleteSql, s.Id)
}

func (s *ImsMcCashRecord) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_cash_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcCashRecord) Get() (err error) {
	var tmp *ImsMcCashRecord
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcCashRecordScan(rows)
		return
	}, ImsMcCashRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcCashRecord) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcCashRecordInsertSql, s.Uniacid, s.Uid, s.ClerkId, s.StoreId, s.ClerkType, s.Fee, s.FinalFee, s.Credit1, s.Credit1Fee, s.Credit2, s.Cash, s.ReturnCash, s.FinalCash, s.Remark, s.Createtime, s.TradeType)
}

func (s *ImsMcCashRecord) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcCashRecordDeleteSql, s.Id)
}

func (s *ImsMcCashRecord) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_cash_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcCashRecord) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcCashRecord
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcCashRecordScan(rows)
		return
	}, ImsMcCashRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcChatsRecord) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcChatsRecordInsertSql, s.Uniacid, s.Acid, s.Flag, s.Openid, s.Msgtype, s.Content, s.Createtime)
}

func (s *ImsMcChatsRecord) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcChatsRecordDeleteSql, s.Id)
}

func (s *ImsMcChatsRecord) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_chats_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcChatsRecord) Get() (err error) {
	var tmp *ImsMcChatsRecord
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcChatsRecordScan(rows)
		return
	}, ImsMcChatsRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcChatsRecord) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcChatsRecordInsertSql, s.Uniacid, s.Acid, s.Flag, s.Openid, s.Msgtype, s.Content, s.Createtime)
}

func (s *ImsMcChatsRecord) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcChatsRecordDeleteSql, s.Id)
}

func (s *ImsMcChatsRecord) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_chats_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcChatsRecord) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcChatsRecord
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcChatsRecordScan(rows)
		return
	}, ImsMcChatsRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcCreditsRecharge) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcCreditsRechargeInsertSql, s.Uniacid, s.Uid, s.Openid, s.Tid, s.Transid, s.Fee, s.Type, s.Tag, s.Status, s.Createtime, s.Backtype)
}

func (s *ImsMcCreditsRecharge) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcCreditsRechargeDeleteSql, s.Id)
}

func (s *ImsMcCreditsRecharge) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_credits_recharge` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcCreditsRecharge) Get() (err error) {
	var tmp *ImsMcCreditsRecharge
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcCreditsRechargeScan(rows)
		return
	}, ImsMcCreditsRechargeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcCreditsRecharge) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcCreditsRechargeInsertSql, s.Uniacid, s.Uid, s.Openid, s.Tid, s.Transid, s.Fee, s.Type, s.Tag, s.Status, s.Createtime, s.Backtype)
}

func (s *ImsMcCreditsRecharge) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcCreditsRechargeDeleteSql, s.Id)
}

func (s *ImsMcCreditsRecharge) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_credits_recharge` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcCreditsRecharge) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcCreditsRecharge
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcCreditsRechargeScan(rows)
		return
	}, ImsMcCreditsRechargeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcCreditsRecord) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcCreditsRecordInsertSql, s.Uid, s.Uniacid, s.Credittype, s.Num, s.Operator, s.Module, s.ClerkId, s.StoreId, s.ClerkType, s.Createtime, s.Remark, s.RealUniacid)
}

func (s *ImsMcCreditsRecord) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcCreditsRecordDeleteSql, s.Id)
}

func (s *ImsMcCreditsRecord) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_credits_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcCreditsRecord) Get() (err error) {
	var tmp *ImsMcCreditsRecord
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcCreditsRecordScan(rows)
		return
	}, ImsMcCreditsRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcCreditsRecord) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcCreditsRecordInsertSql, s.Uid, s.Uniacid, s.Credittype, s.Num, s.Operator, s.Module, s.ClerkId, s.StoreId, s.ClerkType, s.Createtime, s.Remark, s.RealUniacid)
}

func (s *ImsMcCreditsRecord) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcCreditsRecordDeleteSql, s.Id)
}

func (s *ImsMcCreditsRecord) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_credits_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcCreditsRecord) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcCreditsRecord
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcCreditsRecordScan(rows)
		return
	}, ImsMcCreditsRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcFansGroups) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcFansGroupsInsertSql, s.Uniacid, s.Acid, s.Groups)
}

func (s *ImsMcFansGroups) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcFansGroupsDeleteSql, s.Id)
}

func (s *ImsMcFansGroups) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_fans_groups` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcFansGroups) Get() (err error) {
	var tmp *ImsMcFansGroups
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcFansGroupsScan(rows)
		return
	}, ImsMcFansGroupsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcFansGroups) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcFansGroupsInsertSql, s.Uniacid, s.Acid, s.Groups)
}

func (s *ImsMcFansGroups) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcFansGroupsDeleteSql, s.Id)
}

func (s *ImsMcFansGroups) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_fans_groups` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcFansGroups) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcFansGroups
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcFansGroupsScan(rows)
		return
	}, ImsMcFansGroupsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcFansTag) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcFansTagInsertSql, s.Uniacid, s.Fanid, s.Openid, s.Subscribe, s.Nickname, s.Sex, s.Language, s.City, s.Province, s.Country, s.Headimgurl, s.SubscribeTime, s.Unionid, s.Remark, s.Groupid, s.TagidList, s.SubscribeScene, s.QrSceneStr, s.QrScene)
}

func (s *ImsMcFansTag) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcFansTagDeleteSql, s.Id)
}

func (s *ImsMcFansTag) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_fans_tag` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcFansTag) Get() (err error) {
	var tmp *ImsMcFansTag
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcFansTagScan(rows)
		return
	}, ImsMcFansTagSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcFansTag) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcFansTagInsertSql, s.Uniacid, s.Fanid, s.Openid, s.Subscribe, s.Nickname, s.Sex, s.Language, s.City, s.Province, s.Country, s.Headimgurl, s.SubscribeTime, s.Unionid, s.Remark, s.Groupid, s.TagidList, s.SubscribeScene, s.QrSceneStr, s.QrScene)
}

func (s *ImsMcFansTag) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcFansTagDeleteSql, s.Id)
}

func (s *ImsMcFansTag) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_fans_tag` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcFansTag) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcFansTag
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcFansTagScan(rows)
		return
	}, ImsMcFansTagSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcFansTagMapping) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcFansTagMappingInsertSql, s.Fanid, s.Tagid)
}

func (s *ImsMcFansTagMapping) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcFansTagMappingDeleteSql, s.Id)
}

func (s *ImsMcFansTagMapping) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_fans_tag_mapping` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcFansTagMapping) Get() (err error) {
	var tmp *ImsMcFansTagMapping
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcFansTagMappingScan(rows)
		return
	}, ImsMcFansTagMappingSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcFansTagMapping) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcFansTagMappingInsertSql, s.Fanid, s.Tagid)
}

func (s *ImsMcFansTagMapping) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcFansTagMappingDeleteSql, s.Id)
}

func (s *ImsMcFansTagMapping) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_fans_tag_mapping` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcFansTagMapping) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcFansTagMapping
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcFansTagMappingScan(rows)
		return
	}, ImsMcFansTagMappingSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcGroups) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcGroupsInsertSql, s.Uniacid, s.Title, s.Credit, s.Isdefault)
}

func (s *ImsMcGroups) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcGroupsDeleteSql, s.Groupid)
}

func (s *ImsMcGroups) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Groupid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_groups` SET %s WHERE ( `groupid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcGroups) Get() (err error) {
	var tmp *ImsMcGroups
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcGroupsScan(rows)
		return
	}, ImsMcGroupsSelectSql, s.Groupid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcGroups) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcGroupsInsertSql, s.Uniacid, s.Title, s.Credit, s.Isdefault)
}

func (s *ImsMcGroups) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcGroupsDeleteSql, s.Groupid)
}

func (s *ImsMcGroups) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Groupid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_groups` SET %s WHERE ( `groupid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcGroups) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcGroups
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcGroupsScan(rows)
		return
	}, ImsMcGroupsSelectSql, s.Groupid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcHandsel) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcHandselInsertSql, s.Uniacid, s.Touid, s.Fromuid, s.Module, s.Sign, s.Action, s.CreditValue, s.Createtime)
}

func (s *ImsMcHandsel) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcHandselDeleteSql, s.Id)
}

func (s *ImsMcHandsel) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_handsel` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcHandsel) Get() (err error) {
	var tmp *ImsMcHandsel
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcHandselScan(rows)
		return
	}, ImsMcHandselSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcHandsel) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcHandselInsertSql, s.Uniacid, s.Touid, s.Fromuid, s.Module, s.Sign, s.Action, s.CreditValue, s.Createtime)
}

func (s *ImsMcHandsel) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcHandselDeleteSql, s.Id)
}

func (s *ImsMcHandsel) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_handsel` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcHandsel) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcHandsel
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcHandselScan(rows)
		return
	}, ImsMcHandselSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMappingFans) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcMappingFansInsertSql, s.Acid, s.Uniacid, s.Uid, s.Openid, s.Nickname, s.Groupid, s.Salt, s.Follow, s.Followtime, s.Unfollowtime, s.Tag, s.Updatetime, s.Unionid, s.UserFrom)
}

func (s *ImsMcMappingFans) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcMappingFansDeleteSql, s.Fanid)
}

func (s *ImsMcMappingFans) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Fanid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_mapping_fans` SET %s WHERE ( `fanid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMappingFans) Get() (err error) {
	var tmp *ImsMcMappingFans
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMappingFansScan(rows)
		return
	}, ImsMcMappingFansSelectSql, s.Fanid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMappingFans) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcMappingFansInsertSql, s.Acid, s.Uniacid, s.Uid, s.Openid, s.Nickname, s.Groupid, s.Salt, s.Follow, s.Followtime, s.Unfollowtime, s.Tag, s.Updatetime, s.Unionid, s.UserFrom)
}

func (s *ImsMcMappingFans) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcMappingFansDeleteSql, s.Fanid)
}

func (s *ImsMcMappingFans) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Fanid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_mapping_fans` SET %s WHERE ( `fanid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMappingFans) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcMappingFans
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMappingFansScan(rows)
		return
	}, ImsMcMappingFansSelectSql, s.Fanid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMassRecord) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcMassRecordInsertSql, s.Uniacid, s.Acid, s.Groupname, s.Fansnum, s.Msgtype, s.Content, s.Group, s.AttachId, s.MediaId, s.Type, s.Status, s.CronId, s.Sendtime, s.Finalsendtime, s.Createtime, s.MsgId, s.MsgDataId)
}

func (s *ImsMcMassRecord) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcMassRecordDeleteSql, s.Id)
}

func (s *ImsMcMassRecord) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_mass_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMassRecord) Get() (err error) {
	var tmp *ImsMcMassRecord
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMassRecordScan(rows)
		return
	}, ImsMcMassRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMassRecord) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcMassRecordInsertSql, s.Uniacid, s.Acid, s.Groupname, s.Fansnum, s.Msgtype, s.Content, s.Group, s.AttachId, s.MediaId, s.Type, s.Status, s.CronId, s.Sendtime, s.Finalsendtime, s.Createtime, s.MsgId, s.MsgDataId)
}

func (s *ImsMcMassRecord) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcMassRecordDeleteSql, s.Id)
}

func (s *ImsMcMassRecord) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_mass_record` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMassRecord) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcMassRecord
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMassRecordScan(rows)
		return
	}, ImsMcMassRecordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMembers) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcMembersInsertSql, s.Uniacid, s.Mobile, s.Email, s.Password, s.Salt, s.Groupid, s.Credit1, s.Credit2, s.Credit3, s.Credit4, s.Credit5, s.Credit6, s.Createtime, s.Realname, s.Nickname, s.Avatar, s.Qq, s.Vip, s.Gender, s.Birthyear, s.Birthmonth, s.Birthday, s.Constellation, s.Zodiac, s.Telephone, s.Idcard, s.Studentid, s.Grade, s.Address, s.Zipcode, s.Nationality, s.Resideprovince, s.Residecity, s.Residedist, s.Graduateschool, s.Company, s.Education, s.Occupation, s.Position, s.Revenue, s.Affectivestatus, s.Lookingfor, s.Bloodtype, s.Height, s.Weight, s.Alipay, s.Msn, s.Taobao, s.Site, s.Bio, s.Interest, s.PayPassword, s.UserFrom)
}

func (s *ImsMcMembers) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcMembersDeleteSql, s.Uid)
}

func (s *ImsMcMembers) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Uid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_members` SET %s WHERE ( `uid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMembers) Get() (err error) {
	var tmp *ImsMcMembers
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMembersScan(rows)
		return
	}, ImsMcMembersSelectSql, s.Uid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMembers) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcMembersInsertSql, s.Uniacid, s.Mobile, s.Email, s.Password, s.Salt, s.Groupid, s.Credit1, s.Credit2, s.Credit3, s.Credit4, s.Credit5, s.Credit6, s.Createtime, s.Realname, s.Nickname, s.Avatar, s.Qq, s.Vip, s.Gender, s.Birthyear, s.Birthmonth, s.Birthday, s.Constellation, s.Zodiac, s.Telephone, s.Idcard, s.Studentid, s.Grade, s.Address, s.Zipcode, s.Nationality, s.Resideprovince, s.Residecity, s.Residedist, s.Graduateschool, s.Company, s.Education, s.Occupation, s.Position, s.Revenue, s.Affectivestatus, s.Lookingfor, s.Bloodtype, s.Height, s.Weight, s.Alipay, s.Msn, s.Taobao, s.Site, s.Bio, s.Interest, s.PayPassword, s.UserFrom)
}

func (s *ImsMcMembers) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcMembersDeleteSql, s.Uid)
}

func (s *ImsMcMembers) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Uid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_members` SET %s WHERE ( `uid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMembers) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcMembers
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMembersScan(rows)
		return
	}, ImsMcMembersSelectSql, s.Uid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMemberAddress) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcMemberAddressInsertSql, s.Uniacid, s.Uid, s.Username, s.Mobile, s.Zipcode, s.Province, s.City, s.District, s.Address, s.Isdefault)
}

func (s *ImsMcMemberAddress) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcMemberAddressDeleteSql, s.Id)
}

func (s *ImsMcMemberAddress) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_member_address` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMemberAddress) Get() (err error) {
	var tmp *ImsMcMemberAddress
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMemberAddressScan(rows)
		return
	}, ImsMcMemberAddressSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMemberAddress) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcMemberAddressInsertSql, s.Uniacid, s.Uid, s.Username, s.Mobile, s.Zipcode, s.Province, s.City, s.District, s.Address, s.Isdefault)
}

func (s *ImsMcMemberAddress) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcMemberAddressDeleteSql, s.Id)
}

func (s *ImsMcMemberAddress) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_member_address` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMemberAddress) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcMemberAddress
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMemberAddressScan(rows)
		return
	}, ImsMcMemberAddressSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMemberFields) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcMemberFieldsInsertSql, s.Uniacid, s.Fieldid, s.Title, s.Available, s.Displayorder)
}

func (s *ImsMcMemberFields) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcMemberFieldsDeleteSql, s.Id)
}

func (s *ImsMcMemberFields) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_member_fields` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMemberFields) Get() (err error) {
	var tmp *ImsMcMemberFields
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMemberFieldsScan(rows)
		return
	}, ImsMcMemberFieldsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMemberFields) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcMemberFieldsInsertSql, s.Uniacid, s.Fieldid, s.Title, s.Available, s.Displayorder)
}

func (s *ImsMcMemberFields) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcMemberFieldsDeleteSql, s.Id)
}

func (s *ImsMcMemberFields) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_member_fields` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMemberFields) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcMemberFields
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMemberFieldsScan(rows)
		return
	}, ImsMcMemberFieldsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMemberProperty) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcMemberPropertyInsertSql, s.Uniacid, s.Property)
}

func (s *ImsMcMemberProperty) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcMemberPropertyDeleteSql, s.Id)
}

func (s *ImsMcMemberProperty) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_member_property` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMemberProperty) Get() (err error) {
	var tmp *ImsMcMemberProperty
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMemberPropertyScan(rows)
		return
	}, ImsMcMemberPropertySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcMemberProperty) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcMemberPropertyInsertSql, s.Uniacid, s.Property)
}

func (s *ImsMcMemberProperty) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcMemberPropertyDeleteSql, s.Id)
}

func (s *ImsMcMemberProperty) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_member_property` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcMemberProperty) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcMemberProperty
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcMemberPropertyScan(rows)
		return
	}, ImsMcMemberPropertySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcOauthFans) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMcOauthFansInsertSql, s.OauthOpenid, s.Acid, s.Uid, s.Openid)
}

func (s *ImsMcOauthFans) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMcOauthFansDeleteSql, s.Id)
}

func (s *ImsMcOauthFans) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mc_oauth_fans` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcOauthFans) Get() (err error) {
	var tmp *ImsMcOauthFans
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcOauthFansScan(rows)
		return
	}, ImsMcOauthFansSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMcOauthFans) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMcOauthFansInsertSql, s.OauthOpenid, s.Acid, s.Uid, s.Openid)
}

func (s *ImsMcOauthFans) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMcOauthFansDeleteSql, s.Id)
}

func (s *ImsMcOauthFans) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mc_oauth_fans` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMcOauthFans) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMcOauthFans
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMcOauthFansScan(rows)
		return
	}, ImsMcOauthFansSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMenuEvent) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMenuEventInsertSql, s.Uniacid, s.Keyword, s.Type, s.Picmd5, s.Openid, s.Createtime)
}

func (s *ImsMenuEvent) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMenuEventDeleteSql, s.Id)
}

func (s *ImsMenuEvent) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_menu_event` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMenuEvent) Get() (err error) {
	var tmp *ImsMenuEvent
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMenuEventScan(rows)
		return
	}, ImsMenuEventSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMenuEvent) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMenuEventInsertSql, s.Uniacid, s.Keyword, s.Type, s.Picmd5, s.Openid, s.Createtime)
}

func (s *ImsMenuEvent) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMenuEventDeleteSql, s.Id)
}

func (s *ImsMenuEvent) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_menu_event` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMenuEvent) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMenuEvent
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMenuEventScan(rows)
		return
	}, ImsMenuEventSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMessageNoticeLog) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMessageNoticeLogInsertSql, s.Message, s.IsRead, s.Uid, s.Sign, s.Type, s.Status, s.CreateTime, s.EndTime, s.Url)
}

func (s *ImsMessageNoticeLog) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMessageNoticeLogDeleteSql, s.Id)
}

func (s *ImsMessageNoticeLog) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_message_notice_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMessageNoticeLog) Get() (err error) {
	var tmp *ImsMessageNoticeLog
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMessageNoticeLogScan(rows)
		return
	}, ImsMessageNoticeLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMessageNoticeLog) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMessageNoticeLogInsertSql, s.Message, s.IsRead, s.Uid, s.Sign, s.Type, s.Status, s.CreateTime, s.EndTime, s.Url)
}

func (s *ImsMessageNoticeLog) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMessageNoticeLogDeleteSql, s.Id)
}

func (s *ImsMessageNoticeLog) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_message_notice_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMessageNoticeLog) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMessageNoticeLog
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMessageNoticeLogScan(rows)
		return
	}, ImsMessageNoticeLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMobilenumber) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMobilenumberInsertSql, s.Rid, s.Enabled, s.Dateline)
}

func (s *ImsMobilenumber) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMobilenumberDeleteSql, s.Id)
}

func (s *ImsMobilenumber) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_mobilenumber` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMobilenumber) Get() (err error) {
	var tmp *ImsMobilenumber
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMobilenumberScan(rows)
		return
	}, ImsMobilenumberSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMobilenumber) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMobilenumberInsertSql, s.Rid, s.Enabled, s.Dateline)
}

func (s *ImsMobilenumber) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMobilenumberDeleteSql, s.Id)
}

func (s *ImsMobilenumber) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_mobilenumber` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMobilenumber) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMobilenumber
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMobilenumberScan(rows)
		return
	}, ImsMobilenumberSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModules) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsModulesInsertSql, s.Name, s.ApplicationType, s.Type, s.Title, s.Version, s.Ability, s.Description, s.Author, s.Url, s.Settings, s.Subscribes, s.Handles, s.Isrulefields, s.Issystem, s.Target, s.Iscard, s.Permissions, s.TitleInitial, s.WxappSupport, s.WelcomeSupport, s.OauthType, s.WebappSupport, s.PhoneappSupport, s.AccountSupport, s.XzappSupport, s.AliappSupport, s.Logo, s.BaiduappSupport, s.ToutiaoappSupport, s.From, s.CloudRecord, s.Sections, s.Label)
}

func (s *ImsModules) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsModulesDeleteSql, s.Mid)
}

func (s *ImsModules) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Mid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_modules` SET %s WHERE ( `mid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModules) Get() (err error) {
	var tmp *ImsModules
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesScan(rows)
		return
	}, ImsModulesSelectSql, s.Mid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModules) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsModulesInsertSql, s.Name, s.ApplicationType, s.Type, s.Title, s.Version, s.Ability, s.Description, s.Author, s.Url, s.Settings, s.Subscribes, s.Handles, s.Isrulefields, s.Issystem, s.Target, s.Iscard, s.Permissions, s.TitleInitial, s.WxappSupport, s.WelcomeSupport, s.OauthType, s.WebappSupport, s.PhoneappSupport, s.AccountSupport, s.XzappSupport, s.AliappSupport, s.Logo, s.BaiduappSupport, s.ToutiaoappSupport, s.From, s.CloudRecord, s.Sections, s.Label)
}

func (s *ImsModules) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsModulesDeleteSql, s.Mid)
}

func (s *ImsModules) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Mid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_modules` SET %s WHERE ( `mid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModules) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsModules
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesScan(rows)
		return
	}, ImsModulesSelectSql, s.Mid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesBindings) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsModulesBindingsInsertSql, s.Module, s.Entry, s.Call, s.Title, s.Do, s.State, s.Direct, s.Url, s.Icon, s.Displayorder, s.Multilevel, s.Parent)
}

func (s *ImsModulesBindings) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsModulesBindingsDeleteSql, s.Eid)
}

func (s *ImsModulesBindings) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Eid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_modules_bindings` SET %s WHERE ( `eid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesBindings) Get() (err error) {
	var tmp *ImsModulesBindings
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesBindingsScan(rows)
		return
	}, ImsModulesBindingsSelectSql, s.Eid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesBindings) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsModulesBindingsInsertSql, s.Module, s.Entry, s.Call, s.Title, s.Do, s.State, s.Direct, s.Url, s.Icon, s.Displayorder, s.Multilevel, s.Parent)
}

func (s *ImsModulesBindings) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsModulesBindingsDeleteSql, s.Eid)
}

func (s *ImsModulesBindings) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Eid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_modules_bindings` SET %s WHERE ( `eid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesBindings) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsModulesBindings
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesBindingsScan(rows)
		return
	}, ImsModulesBindingsSelectSql, s.Eid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesCloud) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsModulesCloudInsertSql, s.Name, s.ApplicationType, s.Title, s.TitleInitial, s.Logo, s.Version, s.InstallStatus, s.AccountSupport, s.WxappSupport, s.WebappSupport, s.PhoneappSupport, s.WelcomeSupport, s.MainModuleName, s.MainModuleLogo, s.HasNewVersion, s.HasNewBranch, s.IsBan, s.Lastupdatetime, s.XzappSupport, s.CloudId, s.AliappSupport, s.BaiduappSupport, s.ToutiaoappSupport, s.Buytime, s.ModuleStatus, s.Label)
}

func (s *ImsModulesCloud) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsModulesCloudDeleteSql, s.Id)
}

func (s *ImsModulesCloud) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_modules_cloud` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesCloud) Get() (err error) {
	var tmp *ImsModulesCloud
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesCloudScan(rows)
		return
	}, ImsModulesCloudSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesCloud) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsModulesCloudInsertSql, s.Name, s.ApplicationType, s.Title, s.TitleInitial, s.Logo, s.Version, s.InstallStatus, s.AccountSupport, s.WxappSupport, s.WebappSupport, s.PhoneappSupport, s.WelcomeSupport, s.MainModuleName, s.MainModuleLogo, s.HasNewVersion, s.HasNewBranch, s.IsBan, s.Lastupdatetime, s.XzappSupport, s.CloudId, s.AliappSupport, s.BaiduappSupport, s.ToutiaoappSupport, s.Buytime, s.ModuleStatus, s.Label)
}

func (s *ImsModulesCloud) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsModulesCloudDeleteSql, s.Id)
}

func (s *ImsModulesCloud) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_modules_cloud` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesCloud) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsModulesCloud
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesCloudScan(rows)
		return
	}, ImsModulesCloudSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesIgnore) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsModulesIgnoreInsertSql, s.Name, s.Version)
}

func (s *ImsModulesIgnore) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsModulesIgnoreDeleteSql, s.Id)
}

func (s *ImsModulesIgnore) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_modules_ignore` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesIgnore) Get() (err error) {
	var tmp *ImsModulesIgnore
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesIgnoreScan(rows)
		return
	}, ImsModulesIgnoreSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesIgnore) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsModulesIgnoreInsertSql, s.Name, s.Version)
}

func (s *ImsModulesIgnore) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsModulesIgnoreDeleteSql, s.Id)
}

func (s *ImsModulesIgnore) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_modules_ignore` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesIgnore) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsModulesIgnore
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesIgnoreScan(rows)
		return
	}, ImsModulesIgnoreSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesPlugin) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsModulesPluginInsertSql, s.Name, s.MainModule)
}

func (s *ImsModulesPlugin) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsModulesPluginDeleteSql, s.Id)
}

func (s *ImsModulesPlugin) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_modules_plugin` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesPlugin) Get() (err error) {
	var tmp *ImsModulesPlugin
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesPluginScan(rows)
		return
	}, ImsModulesPluginSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesPlugin) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsModulesPluginInsertSql, s.Name, s.MainModule)
}

func (s *ImsModulesPlugin) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsModulesPluginDeleteSql, s.Id)
}

func (s *ImsModulesPlugin) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_modules_plugin` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesPlugin) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsModulesPlugin
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesPluginScan(rows)
		return
	}, ImsModulesPluginSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesPluginRank) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsModulesPluginRankInsertSql, s.Uniacid, s.Uid, s.Rank, s.PluginName, s.MainModuleName)
}

func (s *ImsModulesPluginRank) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsModulesPluginRankDeleteSql, s.Id)
}

func (s *ImsModulesPluginRank) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_modules_plugin_rank` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesPluginRank) Get() (err error) {
	var tmp *ImsModulesPluginRank
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesPluginRankScan(rows)
		return
	}, ImsModulesPluginRankSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesPluginRank) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsModulesPluginRankInsertSql, s.Uniacid, s.Uid, s.Rank, s.PluginName, s.MainModuleName)
}

func (s *ImsModulesPluginRank) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsModulesPluginRankDeleteSql, s.Id)
}

func (s *ImsModulesPluginRank) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_modules_plugin_rank` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesPluginRank) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsModulesPluginRank
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesPluginRankScan(rows)
		return
	}, ImsModulesPluginRankSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesRank) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsModulesRankInsertSql, s.ModuleName, s.Uid, s.Rank, s.Uniacid)
}

func (s *ImsModulesRank) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsModulesRankDeleteSql, s.Id)
}

func (s *ImsModulesRank) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_modules_rank` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesRank) Get() (err error) {
	var tmp *ImsModulesRank
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesRankScan(rows)
		return
	}, ImsModulesRankSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesRank) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsModulesRankInsertSql, s.ModuleName, s.Uid, s.Rank, s.Uniacid)
}

func (s *ImsModulesRank) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsModulesRankDeleteSql, s.Id)
}

func (s *ImsModulesRank) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_modules_rank` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesRank) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsModulesRank
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesRankScan(rows)
		return
	}, ImsModulesRankSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesRecycle) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsModulesRecycleInsertSql, s.Name, s.Type, s.AccountSupport, s.WxappSupport, s.WelcomeSupport, s.WebappSupport, s.PhoneappSupport, s.XzappSupport, s.AliappSupport, s.BaiduappSupport, s.ToutiaoappSupport)
}

func (s *ImsModulesRecycle) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsModulesRecycleDeleteSql, s.Id)
}

func (s *ImsModulesRecycle) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_modules_recycle` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesRecycle) Get() (err error) {
	var tmp *ImsModulesRecycle
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesRecycleScan(rows)
		return
	}, ImsModulesRecycleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsModulesRecycle) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsModulesRecycleInsertSql, s.Name, s.Type, s.AccountSupport, s.WxappSupport, s.WelcomeSupport, s.WebappSupport, s.PhoneappSupport, s.XzappSupport, s.AliappSupport, s.BaiduappSupport, s.ToutiaoappSupport)
}

func (s *ImsModulesRecycle) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsModulesRecycleDeleteSql, s.Id)
}

func (s *ImsModulesRecycle) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_modules_recycle` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsModulesRecycle) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsModulesRecycle
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsModulesRecycleScan(rows)
		return
	}, ImsModulesRecycleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMusicReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsMusicReplyInsertSql, s.Rid, s.Title, s.Description, s.Url, s.Hqurl)
}

func (s *ImsMusicReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsMusicReplyDeleteSql, s.Id)
}

func (s *ImsMusicReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_music_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMusicReply) Get() (err error) {
	var tmp *ImsMusicReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMusicReplyScan(rows)
		return
	}, ImsMusicReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsMusicReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsMusicReplyInsertSql, s.Rid, s.Title, s.Description, s.Url, s.Hqurl)
}

func (s *ImsMusicReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsMusicReplyDeleteSql, s.Id)
}

func (s *ImsMusicReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_music_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsMusicReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsMusicReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsMusicReplyScan(rows)
		return
	}, ImsMusicReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsNewsReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsNewsReplyInsertSql, s.Rid, s.ParentId, s.Title, s.Author, s.Description, s.Thumb, s.Content, s.Url, s.Displayorder, s.Incontent, s.Createtime, s.MediaId)
}

func (s *ImsNewsReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsNewsReplyDeleteSql, s.Id)
}

func (s *ImsNewsReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_news_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsNewsReply) Get() (err error) {
	var tmp *ImsNewsReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsNewsReplyScan(rows)
		return
	}, ImsNewsReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsNewsReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsNewsReplyInsertSql, s.Rid, s.ParentId, s.Title, s.Author, s.Description, s.Thumb, s.Content, s.Url, s.Displayorder, s.Incontent, s.Createtime, s.MediaId)
}

func (s *ImsNewsReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsNewsReplyDeleteSql, s.Id)
}

func (s *ImsNewsReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_news_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsNewsReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsNewsReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsNewsReplyScan(rows)
		return
	}, ImsNewsReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsPhoneappVersions) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsPhoneappVersionsInsertSql, s.Uniacid, s.Version, s.Description, s.Modules, s.Createtime)
}

func (s *ImsPhoneappVersions) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsPhoneappVersionsDeleteSql, s.Id)
}

func (s *ImsPhoneappVersions) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_phoneapp_versions` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsPhoneappVersions) Get() (err error) {
	var tmp *ImsPhoneappVersions
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsPhoneappVersionsScan(rows)
		return
	}, ImsPhoneappVersionsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsPhoneappVersions) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsPhoneappVersionsInsertSql, s.Uniacid, s.Version, s.Description, s.Modules, s.Createtime)
}

func (s *ImsPhoneappVersions) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsPhoneappVersionsDeleteSql, s.Id)
}

func (s *ImsPhoneappVersions) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_phoneapp_versions` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsPhoneappVersions) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsPhoneappVersions
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsPhoneappVersionsScan(rows)
		return
	}, ImsPhoneappVersionsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsProfileFields) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsProfileFieldsInsertSql, s.Field, s.Available, s.Title, s.Description, s.Displayorder, s.Required, s.Unchangeable, s.Showinregister, s.FieldLength)
}

func (s *ImsProfileFields) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsProfileFieldsDeleteSql, s.Id)
}

func (s *ImsProfileFields) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_profile_fields` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsProfileFields) Get() (err error) {
	var tmp *ImsProfileFields
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsProfileFieldsScan(rows)
		return
	}, ImsProfileFieldsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsProfileFields) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsProfileFieldsInsertSql, s.Field, s.Available, s.Title, s.Description, s.Displayorder, s.Required, s.Unchangeable, s.Showinregister, s.FieldLength)
}

func (s *ImsProfileFields) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsProfileFieldsDeleteSql, s.Id)
}

func (s *ImsProfileFields) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_profile_fields` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsProfileFields) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsProfileFields
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsProfileFieldsScan(rows)
		return
	}, ImsProfileFieldsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsQrcode) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsQrcodeInsertSql, s.Uniacid, s.Acid, s.Type, s.Extra, s.Qrcid, s.SceneStr, s.Name, s.Keyword, s.Model, s.Ticket, s.Url, s.Expire, s.Subnum, s.Createtime, s.Status)
}

func (s *ImsQrcode) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsQrcodeDeleteSql, s.Id)
}

func (s *ImsQrcode) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_qrcode` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsQrcode) Get() (err error) {
	var tmp *ImsQrcode
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsQrcodeScan(rows)
		return
	}, ImsQrcodeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsQrcode) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsQrcodeInsertSql, s.Uniacid, s.Acid, s.Type, s.Extra, s.Qrcid, s.SceneStr, s.Name, s.Keyword, s.Model, s.Ticket, s.Url, s.Expire, s.Subnum, s.Createtime, s.Status)
}

func (s *ImsQrcode) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsQrcodeDeleteSql, s.Id)
}

func (s *ImsQrcode) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_qrcode` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsQrcode) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsQrcode
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsQrcodeScan(rows)
		return
	}, ImsQrcodeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsQrcodeStat) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsQrcodeStatInsertSql, s.Uniacid, s.Acid, s.Qid, s.Openid, s.Type, s.Qrcid, s.SceneStr, s.Name, s.Createtime)
}

func (s *ImsQrcodeStat) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsQrcodeStatDeleteSql, s.Id)
}

func (s *ImsQrcodeStat) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_qrcode_stat` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsQrcodeStat) Get() (err error) {
	var tmp *ImsQrcodeStat
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsQrcodeStatScan(rows)
		return
	}, ImsQrcodeStatSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsQrcodeStat) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsQrcodeStatInsertSql, s.Uniacid, s.Acid, s.Qid, s.Openid, s.Type, s.Qrcid, s.SceneStr, s.Name, s.Createtime)
}

func (s *ImsQrcodeStat) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsQrcodeStatDeleteSql, s.Id)
}

func (s *ImsQrcodeStat) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_qrcode_stat` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsQrcodeStat) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsQrcodeStat
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsQrcodeStatScan(rows)
		return
	}, ImsQrcodeStatSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsRule) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsRuleInsertSql, s.Uniacid, s.Name, s.Module, s.Displayorder, s.Status, s.Containtype)
}

func (s *ImsRule) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsRuleDeleteSql, s.Id)
}

func (s *ImsRule) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_rule` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsRule) Get() (err error) {
	var tmp *ImsRule
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsRuleScan(rows)
		return
	}, ImsRuleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsRule) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsRuleInsertSql, s.Uniacid, s.Name, s.Module, s.Displayorder, s.Status, s.Containtype)
}

func (s *ImsRule) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsRuleDeleteSql, s.Id)
}

func (s *ImsRule) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_rule` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsRule) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsRule
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsRuleScan(rows)
		return
	}, ImsRuleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsRuleKeyword) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsRuleKeywordInsertSql, s.Rid, s.Uniacid, s.Module, s.Content, s.Type, s.Displayorder, s.Status)
}

func (s *ImsRuleKeyword) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsRuleKeywordDeleteSql, s.Id)
}

func (s *ImsRuleKeyword) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_rule_keyword` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsRuleKeyword) Get() (err error) {
	var tmp *ImsRuleKeyword
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsRuleKeywordScan(rows)
		return
	}, ImsRuleKeywordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsRuleKeyword) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsRuleKeywordInsertSql, s.Rid, s.Uniacid, s.Module, s.Content, s.Type, s.Displayorder, s.Status)
}

func (s *ImsRuleKeyword) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsRuleKeywordDeleteSql, s.Id)
}

func (s *ImsRuleKeyword) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_rule_keyword` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsRuleKeyword) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsRuleKeyword
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsRuleKeywordScan(rows)
		return
	}, ImsRuleKeywordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteArticle) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteArticleInsertSql, s.Uniacid, s.Rid, s.Kid, s.Iscommend, s.Ishot, s.Pcate, s.Ccate, s.Template, s.Title, s.Description, s.Content, s.Thumb, s.Incontent, s.Source, s.Author, s.Displayorder, s.Linkurl, s.Createtime, s.Edittime, s.Click, s.Type, s.Credit)
}

func (s *ImsSiteArticle) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteArticleDeleteSql, s.Id)
}

func (s *ImsSiteArticle) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_article` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteArticle) Get() (err error) {
	var tmp *ImsSiteArticle
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteArticleScan(rows)
		return
	}, ImsSiteArticleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteArticle) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteArticleInsertSql, s.Uniacid, s.Rid, s.Kid, s.Iscommend, s.Ishot, s.Pcate, s.Ccate, s.Template, s.Title, s.Description, s.Content, s.Thumb, s.Incontent, s.Source, s.Author, s.Displayorder, s.Linkurl, s.Createtime, s.Edittime, s.Click, s.Type, s.Credit)
}

func (s *ImsSiteArticle) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteArticleDeleteSql, s.Id)
}

func (s *ImsSiteArticle) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_article` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteArticle) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteArticle
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteArticleScan(rows)
		return
	}, ImsSiteArticleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteArticleComment) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteArticleCommentInsertSql, s.Uniacid, s.Articleid, s.Parentid, s.Uid, s.Openid, s.Content, s.IsRead, s.Iscomment, s.Createtime)
}

func (s *ImsSiteArticleComment) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteArticleCommentDeleteSql, s.Id)
}

func (s *ImsSiteArticleComment) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_article_comment` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteArticleComment) Get() (err error) {
	var tmp *ImsSiteArticleComment
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteArticleCommentScan(rows)
		return
	}, ImsSiteArticleCommentSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteArticleComment) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteArticleCommentInsertSql, s.Uniacid, s.Articleid, s.Parentid, s.Uid, s.Openid, s.Content, s.IsRead, s.Iscomment, s.Createtime)
}

func (s *ImsSiteArticleComment) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteArticleCommentDeleteSql, s.Id)
}

func (s *ImsSiteArticleComment) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_article_comment` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteArticleComment) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteArticleComment
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteArticleCommentScan(rows)
		return
	}, ImsSiteArticleCommentSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteCategory) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteCategoryInsertSql, s.Uniacid, s.Nid, s.Name, s.Parentid, s.Displayorder, s.Enabled, s.Icon, s.Description, s.Styleid, s.Linkurl, s.Ishomepage, s.Icontype, s.Css, s.Multiid)
}

func (s *ImsSiteCategory) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteCategoryDeleteSql, s.Id)
}

func (s *ImsSiteCategory) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_category` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteCategory) Get() (err error) {
	var tmp *ImsSiteCategory
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteCategoryScan(rows)
		return
	}, ImsSiteCategorySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteCategory) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteCategoryInsertSql, s.Uniacid, s.Nid, s.Name, s.Parentid, s.Displayorder, s.Enabled, s.Icon, s.Description, s.Styleid, s.Linkurl, s.Ishomepage, s.Icontype, s.Css, s.Multiid)
}

func (s *ImsSiteCategory) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteCategoryDeleteSql, s.Id)
}

func (s *ImsSiteCategory) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_category` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteCategory) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteCategory
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteCategoryScan(rows)
		return
	}, ImsSiteCategorySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteMulti) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteMultiInsertSql, s.Uniacid, s.Title, s.Styleid, s.SiteInfo, s.Status, s.Bindhost)
}

func (s *ImsSiteMulti) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteMultiDeleteSql, s.Id)
}

func (s *ImsSiteMulti) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_multi` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteMulti) Get() (err error) {
	var tmp *ImsSiteMulti
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteMultiScan(rows)
		return
	}, ImsSiteMultiSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteMulti) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteMultiInsertSql, s.Uniacid, s.Title, s.Styleid, s.SiteInfo, s.Status, s.Bindhost)
}

func (s *ImsSiteMulti) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteMultiDeleteSql, s.Id)
}

func (s *ImsSiteMulti) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_multi` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteMulti) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteMulti
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteMultiScan(rows)
		return
	}, ImsSiteMultiSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteNav) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteNavInsertSql, s.Uniacid, s.Multiid, s.Section, s.Module, s.Displayorder, s.Name, s.Description, s.Position, s.Url, s.Icon, s.Css, s.Status, s.Categoryid)
}

func (s *ImsSiteNav) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteNavDeleteSql, s.Id)
}

func (s *ImsSiteNav) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_nav` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteNav) Get() (err error) {
	var tmp *ImsSiteNav
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteNavScan(rows)
		return
	}, ImsSiteNavSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteNav) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteNavInsertSql, s.Uniacid, s.Multiid, s.Section, s.Module, s.Displayorder, s.Name, s.Description, s.Position, s.Url, s.Icon, s.Css, s.Status, s.Categoryid)
}

func (s *ImsSiteNav) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteNavDeleteSql, s.Id)
}

func (s *ImsSiteNav) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_nav` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteNav) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteNav
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteNavScan(rows)
		return
	}, ImsSiteNavSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSitePage) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSitePageInsertSql, s.Uniacid, s.Multiid, s.Title, s.Description, s.Params, s.Html, s.Multipage, s.Type, s.Status, s.Createtime, s.Goodnum)
}

func (s *ImsSitePage) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSitePageDeleteSql, s.Id)
}

func (s *ImsSitePage) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_page` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSitePage) Get() (err error) {
	var tmp *ImsSitePage
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSitePageScan(rows)
		return
	}, ImsSitePageSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSitePage) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSitePageInsertSql, s.Uniacid, s.Multiid, s.Title, s.Description, s.Params, s.Html, s.Multipage, s.Type, s.Status, s.Createtime, s.Goodnum)
}

func (s *ImsSitePage) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSitePageDeleteSql, s.Id)
}

func (s *ImsSitePage) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_page` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSitePage) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSitePage
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSitePageScan(rows)
		return
	}, ImsSitePageSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteSlide) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteSlideInsertSql, s.Uniacid, s.Multiid, s.Title, s.Url, s.Thumb, s.Displayorder)
}

func (s *ImsSiteSlide) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteSlideDeleteSql, s.Id)
}

func (s *ImsSiteSlide) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_slide` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteSlide) Get() (err error) {
	var tmp *ImsSiteSlide
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteSlideScan(rows)
		return
	}, ImsSiteSlideSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteSlide) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteSlideInsertSql, s.Uniacid, s.Multiid, s.Title, s.Url, s.Thumb, s.Displayorder)
}

func (s *ImsSiteSlide) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteSlideDeleteSql, s.Id)
}

func (s *ImsSiteSlide) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_slide` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteSlide) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteSlide
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteSlideScan(rows)
		return
	}, ImsSiteSlideSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreCashLog) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteStoreCashLogInsertSql, s.FounderUid, s.Number, s.Amount, s.Status, s.CreateTime)
}

func (s *ImsSiteStoreCashLog) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteStoreCashLogDeleteSql, s.Id)
}

func (s *ImsSiteStoreCashLog) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_store_cash_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreCashLog) Get() (err error) {
	var tmp *ImsSiteStoreCashLog
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreCashLogScan(rows)
		return
	}, ImsSiteStoreCashLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreCashLog) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteStoreCashLogInsertSql, s.FounderUid, s.Number, s.Amount, s.Status, s.CreateTime)
}

func (s *ImsSiteStoreCashLog) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteStoreCashLogDeleteSql, s.Id)
}

func (s *ImsSiteStoreCashLog) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_store_cash_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreCashLog) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteStoreCashLog
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreCashLogScan(rows)
		return
	}, ImsSiteStoreCashLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreCashOrder) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteStoreCashOrderInsertSql, s.Number, s.FounderUid, s.OrderId, s.GoodsId, s.OrderAmount, s.CashLogId, s.Status, s.CreateTime)
}

func (s *ImsSiteStoreCashOrder) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteStoreCashOrderDeleteSql, s.Id)
}

func (s *ImsSiteStoreCashOrder) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_store_cash_order` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreCashOrder) Get() (err error) {
	var tmp *ImsSiteStoreCashOrder
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreCashOrderScan(rows)
		return
	}, ImsSiteStoreCashOrderSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreCashOrder) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteStoreCashOrderInsertSql, s.Number, s.FounderUid, s.OrderId, s.GoodsId, s.OrderAmount, s.CashLogId, s.Status, s.CreateTime)
}

func (s *ImsSiteStoreCashOrder) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteStoreCashOrderDeleteSql, s.Id)
}

func (s *ImsSiteStoreCashOrder) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_store_cash_order` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreCashOrder) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteStoreCashOrder
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreCashOrderScan(rows)
		return
	}, ImsSiteStoreCashOrderSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreCreateAccount) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteStoreCreateAccountInsertSql, s.Uid, s.Uniacid, s.Type, s.Endtime)
}

func (s *ImsSiteStoreCreateAccount) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteStoreCreateAccountDeleteSql, s.Id)
}

func (s *ImsSiteStoreCreateAccount) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_store_create_account` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreCreateAccount) Get() (err error) {
	var tmp *ImsSiteStoreCreateAccount
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreCreateAccountScan(rows)
		return
	}, ImsSiteStoreCreateAccountSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreCreateAccount) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteStoreCreateAccountInsertSql, s.Uid, s.Uniacid, s.Type, s.Endtime)
}

func (s *ImsSiteStoreCreateAccount) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteStoreCreateAccountDeleteSql, s.Id)
}

func (s *ImsSiteStoreCreateAccount) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_store_create_account` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreCreateAccount) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteStoreCreateAccount
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreCreateAccountScan(rows)
		return
	}, ImsSiteStoreCreateAccountSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreGoods) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteStoreGoodsInsertSql, s.Type, s.Title, s.Module, s.AccountNum, s.WxappNum, s.Price, s.Unit, s.Slide, s.CategoryId, s.TitleInitial, s.Status, s.Createtime, s.Synopsis, s.Description, s.ModuleGroup, s.ApiNum, s.UserGroupPrice, s.UserGroup, s.AccountGroup, s.IsWish, s.Logo, s.PlatformNum, s.AliappNum, s.BaiduappNum, s.PhoneappNum, s.ToutiaoappNum, s.WebappNum, s.XzappNum)
}

func (s *ImsSiteStoreGoods) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteStoreGoodsDeleteSql, s.Id)
}

func (s *ImsSiteStoreGoods) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_store_goods` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreGoods) Get() (err error) {
	var tmp *ImsSiteStoreGoods
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreGoodsScan(rows)
		return
	}, ImsSiteStoreGoodsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreGoods) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteStoreGoodsInsertSql, s.Type, s.Title, s.Module, s.AccountNum, s.WxappNum, s.Price, s.Unit, s.Slide, s.CategoryId, s.TitleInitial, s.Status, s.Createtime, s.Synopsis, s.Description, s.ModuleGroup, s.ApiNum, s.UserGroupPrice, s.UserGroup, s.AccountGroup, s.IsWish, s.Logo, s.PlatformNum, s.AliappNum, s.BaiduappNum, s.PhoneappNum, s.ToutiaoappNum, s.WebappNum, s.XzappNum)
}

func (s *ImsSiteStoreGoods) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteStoreGoodsDeleteSql, s.Id)
}

func (s *ImsSiteStoreGoods) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_store_goods` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreGoods) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteStoreGoods
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreGoodsScan(rows)
		return
	}, ImsSiteStoreGoodsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreGoodsCloud) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteStoreGoodsCloudInsertSql, s.CloudId, s.Name, s.Title, s.Logo, s.WishBranch, s.IsEdited, s.Isdeleted, s.Branchs)
}

func (s *ImsSiteStoreGoodsCloud) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteStoreGoodsCloudDeleteSql, s.Id)
}

func (s *ImsSiteStoreGoodsCloud) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_store_goods_cloud` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreGoodsCloud) Get() (err error) {
	var tmp *ImsSiteStoreGoodsCloud
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreGoodsCloudScan(rows)
		return
	}, ImsSiteStoreGoodsCloudSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreGoodsCloud) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteStoreGoodsCloudInsertSql, s.CloudId, s.Name, s.Title, s.Logo, s.WishBranch, s.IsEdited, s.Isdeleted, s.Branchs)
}

func (s *ImsSiteStoreGoodsCloud) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteStoreGoodsCloudDeleteSql, s.Id)
}

func (s *ImsSiteStoreGoodsCloud) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_store_goods_cloud` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreGoodsCloud) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteStoreGoodsCloud
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreGoodsCloudScan(rows)
		return
	}, ImsSiteStoreGoodsCloudSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreOrder) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteStoreOrderInsertSql, s.Orderid, s.Goodsid, s.Duration, s.Buyer, s.Buyerid, s.Amount, s.Type, s.Changeprice, s.Createtime, s.Uniacid, s.Endtime, s.Wxapp, s.IsWish)
}

func (s *ImsSiteStoreOrder) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteStoreOrderDeleteSql, s.Id)
}

func (s *ImsSiteStoreOrder) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_store_order` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreOrder) Get() (err error) {
	var tmp *ImsSiteStoreOrder
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreOrderScan(rows)
		return
	}, ImsSiteStoreOrderSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStoreOrder) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteStoreOrderInsertSql, s.Orderid, s.Goodsid, s.Duration, s.Buyer, s.Buyerid, s.Amount, s.Type, s.Changeprice, s.Createtime, s.Uniacid, s.Endtime, s.Wxapp, s.IsWish)
}

func (s *ImsSiteStoreOrder) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteStoreOrderDeleteSql, s.Id)
}

func (s *ImsSiteStoreOrder) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_store_order` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStoreOrder) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteStoreOrder
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStoreOrderScan(rows)
		return
	}, ImsSiteStoreOrderSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStyles) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteStylesInsertSql, s.Uniacid, s.Templateid, s.Name)
}

func (s *ImsSiteStyles) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteStylesDeleteSql, s.Id)
}

func (s *ImsSiteStyles) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_styles` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStyles) Get() (err error) {
	var tmp *ImsSiteStyles
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStylesScan(rows)
		return
	}, ImsSiteStylesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStyles) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteStylesInsertSql, s.Uniacid, s.Templateid, s.Name)
}

func (s *ImsSiteStyles) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteStylesDeleteSql, s.Id)
}

func (s *ImsSiteStyles) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_styles` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStyles) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteStyles
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStylesScan(rows)
		return
	}, ImsSiteStylesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStylesVars) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteStylesVarsInsertSql, s.Uniacid, s.Templateid, s.Styleid, s.Variable, s.Content, s.Description)
}

func (s *ImsSiteStylesVars) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteStylesVarsDeleteSql, s.Id)
}

func (s *ImsSiteStylesVars) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_styles_vars` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStylesVars) Get() (err error) {
	var tmp *ImsSiteStylesVars
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStylesVarsScan(rows)
		return
	}, ImsSiteStylesVarsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteStylesVars) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteStylesVarsInsertSql, s.Uniacid, s.Templateid, s.Styleid, s.Variable, s.Content, s.Description)
}

func (s *ImsSiteStylesVars) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteStylesVarsDeleteSql, s.Id)
}

func (s *ImsSiteStylesVars) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_styles_vars` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteStylesVars) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteStylesVars
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteStylesVarsScan(rows)
		return
	}, ImsSiteStylesVarsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteTemplates) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSiteTemplatesInsertSql, s.Name, s.Title, s.Version, s.Description, s.Author, s.Url, s.Type, s.Sections)
}

func (s *ImsSiteTemplates) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSiteTemplatesDeleteSql, s.Id)
}

func (s *ImsSiteTemplates) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_site_templates` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteTemplates) Get() (err error) {
	var tmp *ImsSiteTemplates
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteTemplatesScan(rows)
		return
	}, ImsSiteTemplatesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSiteTemplates) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSiteTemplatesInsertSql, s.Name, s.Title, s.Version, s.Description, s.Author, s.Url, s.Type, s.Sections)
}

func (s *ImsSiteTemplates) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSiteTemplatesDeleteSql, s.Id)
}

func (s *ImsSiteTemplates) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_site_templates` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSiteTemplates) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSiteTemplates
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSiteTemplatesScan(rows)
		return
	}, ImsSiteTemplatesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatFans) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsStatFansInsertSql, s.Uniacid, s.New, s.Cancel, s.Cumulate, s.Date)
}

func (s *ImsStatFans) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsStatFansDeleteSql, s.Id)
}

func (s *ImsStatFans) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_stat_fans` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatFans) Get() (err error) {
	var tmp *ImsStatFans
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatFansScan(rows)
		return
	}, ImsStatFansSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatFans) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsStatFansInsertSql, s.Uniacid, s.New, s.Cancel, s.Cumulate, s.Date)
}

func (s *ImsStatFans) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsStatFansDeleteSql, s.Id)
}

func (s *ImsStatFans) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_stat_fans` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatFans) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsStatFans
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatFansScan(rows)
		return
	}, ImsStatFansSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatKeyword) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsStatKeywordInsertSql, s.Uniacid, s.Rid, s.Kid, s.Hit, s.Lastupdate, s.Createtime)
}

func (s *ImsStatKeyword) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsStatKeywordDeleteSql, s.Id)
}

func (s *ImsStatKeyword) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_stat_keyword` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatKeyword) Get() (err error) {
	var tmp *ImsStatKeyword
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatKeywordScan(rows)
		return
	}, ImsStatKeywordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatKeyword) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsStatKeywordInsertSql, s.Uniacid, s.Rid, s.Kid, s.Hit, s.Lastupdate, s.Createtime)
}

func (s *ImsStatKeyword) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsStatKeywordDeleteSql, s.Id)
}

func (s *ImsStatKeyword) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_stat_keyword` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatKeyword) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsStatKeyword
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatKeywordScan(rows)
		return
	}, ImsStatKeywordSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatMsgHistory) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsStatMsgHistoryInsertSql, s.Uniacid, s.Rid, s.Kid, s.FromUser, s.Module, s.Message, s.Type, s.Createtime)
}

func (s *ImsStatMsgHistory) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsStatMsgHistoryDeleteSql, s.Id)
}

func (s *ImsStatMsgHistory) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_stat_msg_history` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatMsgHistory) Get() (err error) {
	var tmp *ImsStatMsgHistory
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatMsgHistoryScan(rows)
		return
	}, ImsStatMsgHistorySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatMsgHistory) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsStatMsgHistoryInsertSql, s.Uniacid, s.Rid, s.Kid, s.FromUser, s.Module, s.Message, s.Type, s.Createtime)
}

func (s *ImsStatMsgHistory) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsStatMsgHistoryDeleteSql, s.Id)
}

func (s *ImsStatMsgHistory) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_stat_msg_history` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatMsgHistory) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsStatMsgHistory
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatMsgHistoryScan(rows)
		return
	}, ImsStatMsgHistorySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatRule) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsStatRuleInsertSql, s.Uniacid, s.Rid, s.Hit, s.Lastupdate, s.Createtime)
}

func (s *ImsStatRule) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsStatRuleDeleteSql, s.Id)
}

func (s *ImsStatRule) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_stat_rule` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatRule) Get() (err error) {
	var tmp *ImsStatRule
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatRuleScan(rows)
		return
	}, ImsStatRuleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatRule) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsStatRuleInsertSql, s.Uniacid, s.Rid, s.Hit, s.Lastupdate, s.Createtime)
}

func (s *ImsStatRule) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsStatRuleDeleteSql, s.Id)
}

func (s *ImsStatRule) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_stat_rule` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatRule) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsStatRule
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatRuleScan(rows)
		return
	}, ImsStatRuleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatVisit) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsStatVisitInsertSql, s.Uniacid, s.Module, s.Count, s.Date, s.Type, s.IpCount)
}

func (s *ImsStatVisit) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsStatVisitDeleteSql, s.Id)
}

func (s *ImsStatVisit) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_stat_visit` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatVisit) Get() (err error) {
	var tmp *ImsStatVisit
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatVisitScan(rows)
		return
	}, ImsStatVisitSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatVisit) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsStatVisitInsertSql, s.Uniacid, s.Module, s.Count, s.Date, s.Type, s.IpCount)
}

func (s *ImsStatVisit) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsStatVisitDeleteSql, s.Id)
}

func (s *ImsStatVisit) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_stat_visit` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatVisit) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsStatVisit
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatVisitScan(rows)
		return
	}, ImsStatVisitSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatVisitIp) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsStatVisitIpInsertSql, s.Ip, s.Uniacid, s.Type, s.Module, s.Date)
}

func (s *ImsStatVisitIp) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsStatVisitIpDeleteSql, s.Id)
}

func (s *ImsStatVisitIp) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_stat_visit_ip` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatVisitIp) Get() (err error) {
	var tmp *ImsStatVisitIp
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatVisitIpScan(rows)
		return
	}, ImsStatVisitIpSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsStatVisitIp) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsStatVisitIpInsertSql, s.Ip, s.Uniacid, s.Type, s.Module, s.Date)
}

func (s *ImsStatVisitIp) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsStatVisitIpDeleteSql, s.Id)
}

func (s *ImsStatVisitIp) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_stat_visit_ip` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsStatVisitIp) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsStatVisitIp
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsStatVisitIpScan(rows)
		return
	}, ImsStatVisitIpSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSystemStatVisit) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSystemStatVisitInsertSql, s.Uniacid, s.Modulename, s.Uid, s.Displayorder, s.Createtime, s.Updatetime)
}

func (s *ImsSystemStatVisit) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSystemStatVisitDeleteSql, s.Id)
}

func (s *ImsSystemStatVisit) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_system_stat_visit` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSystemStatVisit) Get() (err error) {
	var tmp *ImsSystemStatVisit
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSystemStatVisitScan(rows)
		return
	}, ImsSystemStatVisitSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSystemStatVisit) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSystemStatVisitInsertSql, s.Uniacid, s.Modulename, s.Uid, s.Displayorder, s.Createtime, s.Updatetime)
}

func (s *ImsSystemStatVisit) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSystemStatVisitDeleteSql, s.Id)
}

func (s *ImsSystemStatVisit) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_system_stat_visit` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSystemStatVisit) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSystemStatVisit
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSystemStatVisitScan(rows)
		return
	}, ImsSystemStatVisitSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSystemWelcomeBinddomain) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsSystemWelcomeBinddomainInsertSql, s.Uid, s.ModuleName, s.Domain, s.Createtime)
}

func (s *ImsSystemWelcomeBinddomain) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsSystemWelcomeBinddomainDeleteSql, s.Id)
}

func (s *ImsSystemWelcomeBinddomain) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_system_welcome_binddomain` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSystemWelcomeBinddomain) Get() (err error) {
	var tmp *ImsSystemWelcomeBinddomain
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSystemWelcomeBinddomainScan(rows)
		return
	}, ImsSystemWelcomeBinddomainSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsSystemWelcomeBinddomain) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsSystemWelcomeBinddomainInsertSql, s.Uid, s.ModuleName, s.Domain, s.Createtime)
}

func (s *ImsSystemWelcomeBinddomain) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsSystemWelcomeBinddomainDeleteSql, s.Id)
}

func (s *ImsSystemWelcomeBinddomain) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_system_welcome_binddomain` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsSystemWelcomeBinddomain) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsSystemWelcomeBinddomain
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsSystemWelcomeBinddomainScan(rows)
		return
	}, ImsSystemWelcomeBinddomainSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccount) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniAccountInsertSql, s.Groupid, s.Name, s.Description, s.DefaultAcid, s.Rank, s.TitleInitial, s.Createtime, s.Logo, s.Qrcode, s.CreateUid)
}

func (s *ImsUniAccount) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniAccountDeleteSql, s.Uniacid)
}

func (s *ImsUniAccount) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Uniacid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_account` SET %s WHERE ( `uniacid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccount) Get() (err error) {
	var tmp *ImsUniAccount
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountScan(rows)
		return
	}, ImsUniAccountSelectSql, s.Uniacid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccount) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniAccountInsertSql, s.Groupid, s.Name, s.Description, s.DefaultAcid, s.Rank, s.TitleInitial, s.Createtime, s.Logo, s.Qrcode, s.CreateUid)
}

func (s *ImsUniAccount) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniAccountDeleteSql, s.Uniacid)
}

func (s *ImsUniAccount) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Uniacid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_account` SET %s WHERE ( `uniacid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccount) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniAccount
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountScan(rows)
		return
	}, ImsUniAccountSelectSql, s.Uniacid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountExtraModules) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniAccountExtraModulesInsertSql, s.Uniacid, s.Modules)
}

func (s *ImsUniAccountExtraModules) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniAccountExtraModulesDeleteSql, s.Id)
}

func (s *ImsUniAccountExtraModules) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_extra_modules` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountExtraModules) Get() (err error) {
	var tmp *ImsUniAccountExtraModules
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountExtraModulesScan(rows)
		return
	}, ImsUniAccountExtraModulesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountExtraModules) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniAccountExtraModulesInsertSql, s.Uniacid, s.Modules)
}

func (s *ImsUniAccountExtraModules) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniAccountExtraModulesDeleteSql, s.Id)
}

func (s *ImsUniAccountExtraModules) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_extra_modules` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountExtraModules) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniAccountExtraModules
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountExtraModulesScan(rows)
		return
	}, ImsUniAccountExtraModulesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountGroup) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniAccountGroupInsertSql, s.Uniacid, s.Groupid)
}

func (s *ImsUniAccountGroup) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniAccountGroupDeleteSql, s.Id)
}

func (s *ImsUniAccountGroup) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountGroup) Get() (err error) {
	var tmp *ImsUniAccountGroup
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountGroupScan(rows)
		return
	}, ImsUniAccountGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountGroup) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniAccountGroupInsertSql, s.Uniacid, s.Groupid)
}

func (s *ImsUniAccountGroup) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniAccountGroupDeleteSql, s.Id)
}

func (s *ImsUniAccountGroup) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountGroup) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniAccountGroup
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountGroupScan(rows)
		return
	}, ImsUniAccountGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountMenus) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniAccountMenusInsertSql, s.Uniacid, s.Menuid, s.Type, s.Title, s.Sex, s.GroupId, s.ClientPlatformType, s.Area, s.Data, s.Status, s.Createtime, s.Isdeleted)
}

func (s *ImsUniAccountMenus) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniAccountMenusDeleteSql, s.Id)
}

func (s *ImsUniAccountMenus) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_menus` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountMenus) Get() (err error) {
	var tmp *ImsUniAccountMenus
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountMenusScan(rows)
		return
	}, ImsUniAccountMenusSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountMenus) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniAccountMenusInsertSql, s.Uniacid, s.Menuid, s.Type, s.Title, s.Sex, s.GroupId, s.ClientPlatformType, s.Area, s.Data, s.Status, s.Createtime, s.Isdeleted)
}

func (s *ImsUniAccountMenus) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniAccountMenusDeleteSql, s.Id)
}

func (s *ImsUniAccountMenus) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_menus` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountMenus) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniAccountMenus
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountMenusScan(rows)
		return
	}, ImsUniAccountMenusSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountModules) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniAccountModulesInsertSql, s.Uniacid, s.Module, s.Enabled, s.Settings, s.Shortcut, s.Displayorder)
}

func (s *ImsUniAccountModules) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniAccountModulesDeleteSql, s.Id)
}

func (s *ImsUniAccountModules) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_modules` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountModules) Get() (err error) {
	var tmp *ImsUniAccountModules
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountModulesScan(rows)
		return
	}, ImsUniAccountModulesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountModules) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniAccountModulesInsertSql, s.Uniacid, s.Module, s.Enabled, s.Settings, s.Shortcut, s.Displayorder)
}

func (s *ImsUniAccountModules) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniAccountModulesDeleteSql, s.Id)
}

func (s *ImsUniAccountModules) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_modules` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountModules) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniAccountModules
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountModulesScan(rows)
		return
	}, ImsUniAccountModulesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountModulesShortcut) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniAccountModulesShortcutInsertSql, s.Title, s.Url, s.Icon, s.Uniacid, s.VersionId, s.ModuleName)
}

func (s *ImsUniAccountModulesShortcut) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniAccountModulesShortcutDeleteSql, s.Id)
}

func (s *ImsUniAccountModulesShortcut) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_modules_shortcut` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountModulesShortcut) Get() (err error) {
	var tmp *ImsUniAccountModulesShortcut
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountModulesShortcutScan(rows)
		return
	}, ImsUniAccountModulesShortcutSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountModulesShortcut) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniAccountModulesShortcutInsertSql, s.Title, s.Url, s.Icon, s.Uniacid, s.VersionId, s.ModuleName)
}

func (s *ImsUniAccountModulesShortcut) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniAccountModulesShortcutDeleteSql, s.Id)
}

func (s *ImsUniAccountModulesShortcut) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_modules_shortcut` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountModulesShortcut) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniAccountModulesShortcut
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountModulesShortcutScan(rows)
		return
	}, ImsUniAccountModulesShortcutSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountUsers) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniAccountUsersInsertSql, s.Uniacid, s.Uid, s.Role, s.Rank)
}

func (s *ImsUniAccountUsers) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniAccountUsersDeleteSql, s.Id)
}

func (s *ImsUniAccountUsers) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_users` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountUsers) Get() (err error) {
	var tmp *ImsUniAccountUsers
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountUsersScan(rows)
		return
	}, ImsUniAccountUsersSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniAccountUsers) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniAccountUsersInsertSql, s.Uniacid, s.Uid, s.Role, s.Rank)
}

func (s *ImsUniAccountUsers) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniAccountUsersDeleteSql, s.Id)
}

func (s *ImsUniAccountUsers) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_account_users` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniAccountUsers) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniAccountUsers
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniAccountUsersScan(rows)
		return
	}, ImsUniAccountUsersSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniGroup) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniGroupInsertSql, s.OwnerUid, s.Name, s.Modules, s.Templates, s.Uniacid, s.Uid)
}

func (s *ImsUniGroup) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniGroupDeleteSql, s.Id)
}

func (s *ImsUniGroup) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniGroup) Get() (err error) {
	var tmp *ImsUniGroup
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniGroupScan(rows)
		return
	}, ImsUniGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniGroup) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniGroupInsertSql, s.OwnerUid, s.Name, s.Modules, s.Templates, s.Uniacid, s.Uid)
}

func (s *ImsUniGroup) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniGroupDeleteSql, s.Id)
}

func (s *ImsUniGroup) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniGroup) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniGroup
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniGroupScan(rows)
		return
	}, ImsUniGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniLinkUniacid) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniLinkUniacidInsertSql, s.Uniacid, s.LinkUniacid, s.VersionId, s.ModuleName)
}

func (s *ImsUniLinkUniacid) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniLinkUniacidDeleteSql, s.Id)
}

func (s *ImsUniLinkUniacid) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_link_uniacid` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniLinkUniacid) Get() (err error) {
	var tmp *ImsUniLinkUniacid
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniLinkUniacidScan(rows)
		return
	}, ImsUniLinkUniacidSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniLinkUniacid) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniLinkUniacidInsertSql, s.Uniacid, s.LinkUniacid, s.VersionId, s.ModuleName)
}

func (s *ImsUniLinkUniacid) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniLinkUniacidDeleteSql, s.Id)
}

func (s *ImsUniLinkUniacid) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_link_uniacid` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniLinkUniacid) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniLinkUniacid
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniLinkUniacidScan(rows)
		return
	}, ImsUniLinkUniacidSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniModules) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniModulesInsertSql, s.Uniacid, s.ModuleName)
}

func (s *ImsUniModules) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniModulesDeleteSql, s.Id)
}

func (s *ImsUniModules) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_modules` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniModules) Get() (err error) {
	var tmp *ImsUniModules
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniModulesScan(rows)
		return
	}, ImsUniModulesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniModules) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniModulesInsertSql, s.Uniacid, s.ModuleName)
}

func (s *ImsUniModules) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniModulesDeleteSql, s.Id)
}

func (s *ImsUniModules) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_modules` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniModules) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniModules
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniModulesScan(rows)
		return
	}, ImsUniModulesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniSettings) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniSettingsInsertSql, s.Passport, s.Oauth, s.JsauthAcid, s.Notify, s.Creditnames, s.Creditbehaviors, s.Welcome, s.Default, s.DefaultMessage, s.Payment, s.Stat, s.DefaultSite, s.Sync, s.Recharge, s.Tplnotice, s.Grouplevel, s.Mcplugin, s.ExchangeEnable, s.CouponType, s.Menuset, s.Statistics, s.BindDomain, s.CommentStatus, s.ReplySetting, s.DefaultModule, s.AttachmentLimit, s.AttachmentSize, s.SyncMember, s.Remote)
}

func (s *ImsUniSettings) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniSettingsDeleteSql, s.Uniacid)
}

func (s *ImsUniSettings) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Uniacid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_settings` SET %s WHERE ( `uniacid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniSettings) Get() (err error) {
	var tmp *ImsUniSettings
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniSettingsScan(rows)
		return
	}, ImsUniSettingsSelectSql, s.Uniacid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniSettings) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniSettingsInsertSql, s.Passport, s.Oauth, s.JsauthAcid, s.Notify, s.Creditnames, s.Creditbehaviors, s.Welcome, s.Default, s.DefaultMessage, s.Payment, s.Stat, s.DefaultSite, s.Sync, s.Recharge, s.Tplnotice, s.Grouplevel, s.Mcplugin, s.ExchangeEnable, s.CouponType, s.Menuset, s.Statistics, s.BindDomain, s.CommentStatus, s.ReplySetting, s.DefaultModule, s.AttachmentLimit, s.AttachmentSize, s.SyncMember, s.Remote)
}

func (s *ImsUniSettings) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniSettingsDeleteSql, s.Uniacid)
}

func (s *ImsUniSettings) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Uniacid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_settings` SET %s WHERE ( `uniacid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniSettings) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniSettings
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniSettingsScan(rows)
		return
	}, ImsUniSettingsSelectSql, s.Uniacid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniVerifycode) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUniVerifycodeInsertSql, s.Uniacid, s.Receiver, s.Verifycode, s.Total, s.Createtime, s.FailedCount)
}

func (s *ImsUniVerifycode) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUniVerifycodeDeleteSql, s.Id)
}

func (s *ImsUniVerifycode) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_uni_verifycode` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniVerifycode) Get() (err error) {
	var tmp *ImsUniVerifycode
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniVerifycodeScan(rows)
		return
	}, ImsUniVerifycodeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUniVerifycode) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUniVerifycodeInsertSql, s.Uniacid, s.Receiver, s.Verifycode, s.Total, s.Createtime, s.FailedCount)
}

func (s *ImsUniVerifycode) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUniVerifycodeDeleteSql, s.Id)
}

func (s *ImsUniVerifycode) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_uni_verifycode` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUniVerifycode) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUniVerifycode
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUniVerifycodeScan(rows)
		return
	}, ImsUniVerifycodeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUserapiCache) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUserapiCacheInsertSql, s.Key, s.Content, s.Lastupdate)
}

func (s *ImsUserapiCache) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUserapiCacheDeleteSql, s.Id)
}

func (s *ImsUserapiCache) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_userapi_cache` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUserapiCache) Get() (err error) {
	var tmp *ImsUserapiCache
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUserapiCacheScan(rows)
		return
	}, ImsUserapiCacheSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUserapiCache) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUserapiCacheInsertSql, s.Key, s.Content, s.Lastupdate)
}

func (s *ImsUserapiCache) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUserapiCacheDeleteSql, s.Id)
}

func (s *ImsUserapiCache) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_userapi_cache` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUserapiCache) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUserapiCache
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUserapiCacheScan(rows)
		return
	}, ImsUserapiCacheSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUserapiReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUserapiReplyInsertSql, s.Rid, s.Description, s.Apiurl, s.Token, s.DefaultText, s.Cachetime)
}

func (s *ImsUserapiReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUserapiReplyDeleteSql, s.Id)
}

func (s *ImsUserapiReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_userapi_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUserapiReply) Get() (err error) {
	var tmp *ImsUserapiReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUserapiReplyScan(rows)
		return
	}, ImsUserapiReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUserapiReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUserapiReplyInsertSql, s.Rid, s.Description, s.Apiurl, s.Token, s.DefaultText, s.Cachetime)
}

func (s *ImsUserapiReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUserapiReplyDeleteSql, s.Id)
}

func (s *ImsUserapiReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_userapi_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUserapiReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUserapiReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUserapiReplyScan(rows)
		return
	}, ImsUserapiReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsers) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersInsertSql, s.OwnerUid, s.Groupid, s.FounderGroupid, s.Username, s.Password, s.Salt, s.Type, s.Status, s.Joindate, s.Joinip, s.Lastvisit, s.Lastip, s.Remark, s.Starttime, s.Endtime, s.RegisterType, s.Openid, s.WelcomeLink, s.NoticeSetting, s.IsBind)
}

func (s *ImsUsers) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersDeleteSql, s.Uid)
}

func (s *ImsUsers) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Uid
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users` SET %s WHERE ( `uid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsers) Get() (err error) {
	var tmp *ImsUsers
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersScan(rows)
		return
	}, ImsUsersSelectSql, s.Uid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsers) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersInsertSql, s.OwnerUid, s.Groupid, s.FounderGroupid, s.Username, s.Password, s.Salt, s.Type, s.Status, s.Joindate, s.Joinip, s.Lastvisit, s.Lastip, s.Remark, s.Starttime, s.Endtime, s.RegisterType, s.Openid, s.WelcomeLink, s.NoticeSetting, s.IsBind)
}

func (s *ImsUsers) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersDeleteSql, s.Uid)
}

func (s *ImsUsers) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Uid
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users` SET %s WHERE ( `uid` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsers) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsers
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersScan(rows)
		return
	}, ImsUsersSelectSql, s.Uid)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersBind) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersBindInsertSql, s.Uid, s.BindSign, s.ThirdType, s.ThirdNickname)
}

func (s *ImsUsersBind) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersBindDeleteSql, s.Id)
}

func (s *ImsUsersBind) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_bind` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersBind) Get() (err error) {
	var tmp *ImsUsersBind
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersBindScan(rows)
		return
	}, ImsUsersBindSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersBind) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersBindInsertSql, s.Uid, s.BindSign, s.ThirdType, s.ThirdNickname)
}

func (s *ImsUsersBind) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersBindDeleteSql, s.Id)
}

func (s *ImsUsersBind) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_bind` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersBind) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersBind
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersBindScan(rows)
		return
	}, ImsUsersBindSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersCreateGroup) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersCreateGroupInsertSql, s.GroupName, s.Maxaccount, s.Maxwxapp, s.Maxwebapp, s.Maxphoneapp, s.Maxxzapp, s.Maxaliapp, s.Createtime, s.Maxbaiduapp, s.Maxtoutiaoapp)
}

func (s *ImsUsersCreateGroup) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersCreateGroupDeleteSql, s.Id)
}

func (s *ImsUsersCreateGroup) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_create_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersCreateGroup) Get() (err error) {
	var tmp *ImsUsersCreateGroup
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersCreateGroupScan(rows)
		return
	}, ImsUsersCreateGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersCreateGroup) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersCreateGroupInsertSql, s.GroupName, s.Maxaccount, s.Maxwxapp, s.Maxwebapp, s.Maxphoneapp, s.Maxxzapp, s.Maxaliapp, s.Createtime, s.Maxbaiduapp, s.Maxtoutiaoapp)
}

func (s *ImsUsersCreateGroup) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersCreateGroupDeleteSql, s.Id)
}

func (s *ImsUsersCreateGroup) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_create_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersCreateGroup) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersCreateGroup
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersCreateGroupScan(rows)
		return
	}, ImsUsersCreateGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersExtraGroup) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersExtraGroupInsertSql, s.Uid, s.UniGroupId, s.CreateGroupId)
}

func (s *ImsUsersExtraGroup) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersExtraGroupDeleteSql, s.Id)
}

func (s *ImsUsersExtraGroup) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_extra_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersExtraGroup) Get() (err error) {
	var tmp *ImsUsersExtraGroup
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersExtraGroupScan(rows)
		return
	}, ImsUsersExtraGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersExtraGroup) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersExtraGroupInsertSql, s.Uid, s.UniGroupId, s.CreateGroupId)
}

func (s *ImsUsersExtraGroup) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersExtraGroupDeleteSql, s.Id)
}

func (s *ImsUsersExtraGroup) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_extra_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersExtraGroup) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersExtraGroup
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersExtraGroupScan(rows)
		return
	}, ImsUsersExtraGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersExtraLimit) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersExtraLimitInsertSql, s.Uid, s.Maxaccount, s.Maxwxapp, s.Maxwebapp, s.Maxphoneapp, s.Maxxzapp, s.Maxaliapp, s.Timelimit, s.Maxbaiduapp, s.Maxtoutiaoapp)
}

func (s *ImsUsersExtraLimit) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersExtraLimitDeleteSql, s.Id)
}

func (s *ImsUsersExtraLimit) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_extra_limit` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersExtraLimit) Get() (err error) {
	var tmp *ImsUsersExtraLimit
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersExtraLimitScan(rows)
		return
	}, ImsUsersExtraLimitSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersExtraLimit) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersExtraLimitInsertSql, s.Uid, s.Maxaccount, s.Maxwxapp, s.Maxwebapp, s.Maxphoneapp, s.Maxxzapp, s.Maxaliapp, s.Timelimit, s.Maxbaiduapp, s.Maxtoutiaoapp)
}

func (s *ImsUsersExtraLimit) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersExtraLimitDeleteSql, s.Id)
}

func (s *ImsUsersExtraLimit) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_extra_limit` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersExtraLimit) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersExtraLimit
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersExtraLimitScan(rows)
		return
	}, ImsUsersExtraLimitSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersExtraModules) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersExtraModulesInsertSql, s.Uid, s.ModuleName, s.Support)
}

func (s *ImsUsersExtraModules) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersExtraModulesDeleteSql, s.Id)
}

func (s *ImsUsersExtraModules) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_extra_modules` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersExtraModules) Get() (err error) {
	var tmp *ImsUsersExtraModules
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersExtraModulesScan(rows)
		return
	}, ImsUsersExtraModulesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersExtraModules) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersExtraModulesInsertSql, s.Uid, s.ModuleName, s.Support)
}

func (s *ImsUsersExtraModules) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersExtraModulesDeleteSql, s.Id)
}

func (s *ImsUsersExtraModules) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_extra_modules` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersExtraModules) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersExtraModules
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersExtraModulesScan(rows)
		return
	}, ImsUsersExtraModulesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersExtraTemplates) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersExtraTemplatesInsertSql, s.Uid, s.TemplateId)
}

func (s *ImsUsersExtraTemplates) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersExtraTemplatesDeleteSql, s.Id)
}

func (s *ImsUsersExtraTemplates) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_extra_templates` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersExtraTemplates) Get() (err error) {
	var tmp *ImsUsersExtraTemplates
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersExtraTemplatesScan(rows)
		return
	}, ImsUsersExtraTemplatesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersExtraTemplates) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersExtraTemplatesInsertSql, s.Uid, s.TemplateId)
}

func (s *ImsUsersExtraTemplates) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersExtraTemplatesDeleteSql, s.Id)
}

func (s *ImsUsersExtraTemplates) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_extra_templates` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersExtraTemplates) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersExtraTemplates
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersExtraTemplatesScan(rows)
		return
	}, ImsUsersExtraTemplatesSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFailedLogin) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersFailedLoginInsertSql, s.Ip, s.Username, s.Count, s.Lastupdate)
}

func (s *ImsUsersFailedLogin) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersFailedLoginDeleteSql, s.Id)
}

func (s *ImsUsersFailedLogin) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_failed_login` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFailedLogin) Get() (err error) {
	var tmp *ImsUsersFailedLogin
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFailedLoginScan(rows)
		return
	}, ImsUsersFailedLoginSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFailedLogin) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersFailedLoginInsertSql, s.Ip, s.Username, s.Count, s.Lastupdate)
}

func (s *ImsUsersFailedLogin) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersFailedLoginDeleteSql, s.Id)
}

func (s *ImsUsersFailedLogin) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_failed_login` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFailedLogin) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersFailedLogin
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFailedLoginScan(rows)
		return
	}, ImsUsersFailedLoginSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderGroup) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersFounderGroupInsertSql, s.Name, s.Package, s.Maxaccount, s.Timelimit, s.Maxwxapp, s.Maxwebapp, s.Maxphoneapp, s.Maxxzapp, s.Maxaliapp, s.Maxbaiduapp, s.Maxtoutiaoapp)
}

func (s *ImsUsersFounderGroup) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersFounderGroupDeleteSql, s.Id)
}

func (s *ImsUsersFounderGroup) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderGroup) Get() (err error) {
	var tmp *ImsUsersFounderGroup
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderGroupScan(rows)
		return
	}, ImsUsersFounderGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderGroup) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersFounderGroupInsertSql, s.Name, s.Package, s.Maxaccount, s.Timelimit, s.Maxwxapp, s.Maxwebapp, s.Maxphoneapp, s.Maxxzapp, s.Maxaliapp, s.Maxbaiduapp, s.Maxtoutiaoapp)
}

func (s *ImsUsersFounderGroup) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersFounderGroupDeleteSql, s.Id)
}

func (s *ImsUsersFounderGroup) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderGroup) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersFounderGroup
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderGroupScan(rows)
		return
	}, ImsUsersFounderGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderOwnCreateGroups) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersFounderOwnCreateGroupsInsertSql, s.FounderUid, s.CreateGroupId)
}

func (s *ImsUsersFounderOwnCreateGroups) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersFounderOwnCreateGroupsDeleteSql, s.Id)
}

func (s *ImsUsersFounderOwnCreateGroups) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_own_create_groups` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderOwnCreateGroups) Get() (err error) {
	var tmp *ImsUsersFounderOwnCreateGroups
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderOwnCreateGroupsScan(rows)
		return
	}, ImsUsersFounderOwnCreateGroupsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderOwnCreateGroups) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersFounderOwnCreateGroupsInsertSql, s.FounderUid, s.CreateGroupId)
}

func (s *ImsUsersFounderOwnCreateGroups) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersFounderOwnCreateGroupsDeleteSql, s.Id)
}

func (s *ImsUsersFounderOwnCreateGroups) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_own_create_groups` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderOwnCreateGroups) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersFounderOwnCreateGroups
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderOwnCreateGroupsScan(rows)
		return
	}, ImsUsersFounderOwnCreateGroupsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderOwnUniGroups) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersFounderOwnUniGroupsInsertSql, s.FounderUid, s.UniGroupId)
}

func (s *ImsUsersFounderOwnUniGroups) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersFounderOwnUniGroupsDeleteSql, s.Id)
}

func (s *ImsUsersFounderOwnUniGroups) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_own_uni_groups` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderOwnUniGroups) Get() (err error) {
	var tmp *ImsUsersFounderOwnUniGroups
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderOwnUniGroupsScan(rows)
		return
	}, ImsUsersFounderOwnUniGroupsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderOwnUniGroups) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersFounderOwnUniGroupsInsertSql, s.FounderUid, s.UniGroupId)
}

func (s *ImsUsersFounderOwnUniGroups) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersFounderOwnUniGroupsDeleteSql, s.Id)
}

func (s *ImsUsersFounderOwnUniGroups) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_own_uni_groups` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderOwnUniGroups) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersFounderOwnUniGroups
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderOwnUniGroupsScan(rows)
		return
	}, ImsUsersFounderOwnUniGroupsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderOwnUsers) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersFounderOwnUsersInsertSql, s.Uid, s.FounderUid)
}

func (s *ImsUsersFounderOwnUsers) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersFounderOwnUsersDeleteSql, s.Id)
}

func (s *ImsUsersFounderOwnUsers) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_own_users` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderOwnUsers) Get() (err error) {
	var tmp *ImsUsersFounderOwnUsers
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderOwnUsersScan(rows)
		return
	}, ImsUsersFounderOwnUsersSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderOwnUsers) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersFounderOwnUsersInsertSql, s.Uid, s.FounderUid)
}

func (s *ImsUsersFounderOwnUsers) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersFounderOwnUsersDeleteSql, s.Id)
}

func (s *ImsUsersFounderOwnUsers) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_own_users` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderOwnUsers) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersFounderOwnUsers
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderOwnUsersScan(rows)
		return
	}, ImsUsersFounderOwnUsersSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderOwnUsersGroups) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersFounderOwnUsersGroupsInsertSql, s.FounderUid, s.UsersGroupId)
}

func (s *ImsUsersFounderOwnUsersGroups) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersFounderOwnUsersGroupsDeleteSql, s.Id)
}

func (s *ImsUsersFounderOwnUsersGroups) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_own_users_groups` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderOwnUsersGroups) Get() (err error) {
	var tmp *ImsUsersFounderOwnUsersGroups
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderOwnUsersGroupsScan(rows)
		return
	}, ImsUsersFounderOwnUsersGroupsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersFounderOwnUsersGroups) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersFounderOwnUsersGroupsInsertSql, s.FounderUid, s.UsersGroupId)
}

func (s *ImsUsersFounderOwnUsersGroups) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersFounderOwnUsersGroupsDeleteSql, s.Id)
}

func (s *ImsUsersFounderOwnUsersGroups) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_founder_own_users_groups` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersFounderOwnUsersGroups) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersFounderOwnUsersGroups
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersFounderOwnUsersGroupsScan(rows)
		return
	}, ImsUsersFounderOwnUsersGroupsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersGroup) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersGroupInsertSql, s.OwnerUid, s.Name, s.Package, s.Maxaccount, s.Timelimit, s.Maxwxapp, s.Maxwebapp, s.Maxphoneapp, s.Maxxzapp, s.Maxaliapp, s.Maxbaiduapp, s.Maxtoutiaoapp)
}

func (s *ImsUsersGroup) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersGroupDeleteSql, s.Id)
}

func (s *ImsUsersGroup) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersGroup) Get() (err error) {
	var tmp *ImsUsersGroup
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersGroupScan(rows)
		return
	}, ImsUsersGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersGroup) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersGroupInsertSql, s.OwnerUid, s.Name, s.Package, s.Maxaccount, s.Timelimit, s.Maxwxapp, s.Maxwebapp, s.Maxphoneapp, s.Maxxzapp, s.Maxaliapp, s.Maxbaiduapp, s.Maxtoutiaoapp)
}

func (s *ImsUsersGroup) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersGroupDeleteSql, s.Id)
}

func (s *ImsUsersGroup) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_group` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersGroup) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersGroup
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersGroupScan(rows)
		return
	}, ImsUsersGroupSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersInvitation) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersInvitationInsertSql, s.Code, s.Fromuid, s.Inviteuid, s.Createtime)
}

func (s *ImsUsersInvitation) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersInvitationDeleteSql, s.Id)
}

func (s *ImsUsersInvitation) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_invitation` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersInvitation) Get() (err error) {
	var tmp *ImsUsersInvitation
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersInvitationScan(rows)
		return
	}, ImsUsersInvitationSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersInvitation) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersInvitationInsertSql, s.Code, s.Fromuid, s.Inviteuid, s.Createtime)
}

func (s *ImsUsersInvitation) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersInvitationDeleteSql, s.Id)
}

func (s *ImsUsersInvitation) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_invitation` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersInvitation) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersInvitation
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersInvitationScan(rows)
		return
	}, ImsUsersInvitationSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersLastuse) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersLastuseInsertSql, s.Uid, s.Uniacid, s.Modulename, s.Type)
}

func (s *ImsUsersLastuse) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersLastuseDeleteSql, s.Id)
}

func (s *ImsUsersLastuse) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_lastuse` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersLastuse) Get() (err error) {
	var tmp *ImsUsersLastuse
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersLastuseScan(rows)
		return
	}, ImsUsersLastuseSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersLastuse) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersLastuseInsertSql, s.Uid, s.Uniacid, s.Modulename, s.Type)
}

func (s *ImsUsersLastuse) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersLastuseDeleteSql, s.Id)
}

func (s *ImsUsersLastuse) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_lastuse` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersLastuse) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersLastuse
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersLastuseScan(rows)
		return
	}, ImsUsersLastuseSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersLoginLogs) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersLoginLogsInsertSql, s.Uid, s.Ip, s.City, s.LoginAt, s.Createtime)
}

func (s *ImsUsersLoginLogs) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersLoginLogsDeleteSql, s.Id)
}

func (s *ImsUsersLoginLogs) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_login_logs` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersLoginLogs) Get() (err error) {
	var tmp *ImsUsersLoginLogs
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersLoginLogsScan(rows)
		return
	}, ImsUsersLoginLogsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersLoginLogs) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersLoginLogsInsertSql, s.Uid, s.Ip, s.City, s.LoginAt, s.Createtime)
}

func (s *ImsUsersLoginLogs) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersLoginLogsDeleteSql, s.Id)
}

func (s *ImsUsersLoginLogs) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_login_logs` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersLoginLogs) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersLoginLogs
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersLoginLogsScan(rows)
		return
	}, ImsUsersLoginLogsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersOperateHistory) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersOperateHistoryInsertSql, s.Type, s.Uid, s.Uniacid, s.ModuleName, s.Createtime)
}

func (s *ImsUsersOperateHistory) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersOperateHistoryDeleteSql, s.Id)
}

func (s *ImsUsersOperateHistory) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_operate_history` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersOperateHistory) Get() (err error) {
	var tmp *ImsUsersOperateHistory
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersOperateHistoryScan(rows)
		return
	}, ImsUsersOperateHistorySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersOperateHistory) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersOperateHistoryInsertSql, s.Type, s.Uid, s.Uniacid, s.ModuleName, s.Createtime)
}

func (s *ImsUsersOperateHistory) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersOperateHistoryDeleteSql, s.Id)
}

func (s *ImsUsersOperateHistory) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_operate_history` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersOperateHistory) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersOperateHistory
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersOperateHistoryScan(rows)
		return
	}, ImsUsersOperateHistorySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersOperateStar) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersOperateStarInsertSql, s.Type, s.Uid, s.Uniacid, s.ModuleName, s.Rank, s.Createtime)
}

func (s *ImsUsersOperateStar) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersOperateStarDeleteSql, s.Id)
}

func (s *ImsUsersOperateStar) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_operate_star` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersOperateStar) Get() (err error) {
	var tmp *ImsUsersOperateStar
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersOperateStarScan(rows)
		return
	}, ImsUsersOperateStarSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersOperateStar) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersOperateStarInsertSql, s.Type, s.Uid, s.Uniacid, s.ModuleName, s.Rank, s.Createtime)
}

func (s *ImsUsersOperateStar) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersOperateStarDeleteSql, s.Id)
}

func (s *ImsUsersOperateStar) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_operate_star` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersOperateStar) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersOperateStar
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersOperateStarScan(rows)
		return
	}, ImsUsersOperateStarSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersPermission) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersPermissionInsertSql, s.Uniacid, s.Uid, s.Type, s.Permission, s.Url, s.Modules, s.Templates)
}

func (s *ImsUsersPermission) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersPermissionDeleteSql, s.Id)
}

func (s *ImsUsersPermission) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_permission` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersPermission) Get() (err error) {
	var tmp *ImsUsersPermission
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersPermissionScan(rows)
		return
	}, ImsUsersPermissionSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersPermission) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersPermissionInsertSql, s.Uniacid, s.Uid, s.Type, s.Permission, s.Url, s.Modules, s.Templates)
}

func (s *ImsUsersPermission) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersPermissionDeleteSql, s.Id)
}

func (s *ImsUsersPermission) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_permission` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersPermission) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersPermission
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersPermissionScan(rows)
		return
	}, ImsUsersPermissionSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersProfile) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsUsersProfileInsertSql, s.Uid, s.Createtime, s.Edittime, s.Realname, s.Nickname, s.Avatar, s.Qq, s.Mobile, s.Fakeid, s.Vip, s.Gender, s.Birthyear, s.Birthmonth, s.Birthday, s.Constellation, s.Zodiac, s.Telephone, s.Idcard, s.Studentid, s.Grade, s.Address, s.Zipcode, s.Nationality, s.Resideprovince, s.Residecity, s.Residedist, s.Graduateschool, s.Company, s.Education, s.Occupation, s.Position, s.Revenue, s.Affectivestatus, s.Lookingfor, s.Bloodtype, s.Height, s.Weight, s.Alipay, s.Msn, s.Email, s.Taobao, s.Site, s.Bio, s.Interest, s.Workerid, s.IsSendMobileStatus, s.SendExpireStatus)
}

func (s *ImsUsersProfile) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsUsersProfileDeleteSql, s.Id)
}

func (s *ImsUsersProfile) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_users_profile` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersProfile) Get() (err error) {
	var tmp *ImsUsersProfile
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersProfileScan(rows)
		return
	}, ImsUsersProfileSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsUsersProfile) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsUsersProfileInsertSql, s.Uid, s.Createtime, s.Edittime, s.Realname, s.Nickname, s.Avatar, s.Qq, s.Mobile, s.Fakeid, s.Vip, s.Gender, s.Birthyear, s.Birthmonth, s.Birthday, s.Constellation, s.Zodiac, s.Telephone, s.Idcard, s.Studentid, s.Grade, s.Address, s.Zipcode, s.Nationality, s.Resideprovince, s.Residecity, s.Residedist, s.Graduateschool, s.Company, s.Education, s.Occupation, s.Position, s.Revenue, s.Affectivestatus, s.Lookingfor, s.Bloodtype, s.Height, s.Weight, s.Alipay, s.Msn, s.Email, s.Taobao, s.Site, s.Bio, s.Interest, s.Workerid, s.IsSendMobileStatus, s.SendExpireStatus)
}

func (s *ImsUsersProfile) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsUsersProfileDeleteSql, s.Id)
}

func (s *ImsUsersProfile) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_users_profile` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsUsersProfile) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsUsersProfile
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsUsersProfileScan(rows)
		return
	}, ImsUsersProfileSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVideoReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVideoReplyInsertSql, s.Rid, s.Title, s.Description, s.Mediaid, s.Createtime)
}

func (s *ImsVideoReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVideoReplyDeleteSql, s.Id)
}

func (s *ImsVideoReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_video_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVideoReply) Get() (err error) {
	var tmp *ImsVideoReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVideoReplyScan(rows)
		return
	}, ImsVideoReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVideoReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVideoReplyInsertSql, s.Rid, s.Title, s.Description, s.Mediaid, s.Createtime)
}

func (s *ImsVideoReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVideoReplyDeleteSql, s.Id)
}

func (s *ImsVideoReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_video_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVideoReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVideoReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVideoReplyScan(rows)
		return
	}, ImsVideoReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVoiceReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVoiceReplyInsertSql, s.Rid, s.Title, s.Mediaid, s.Createtime)
}

func (s *ImsVoiceReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVoiceReplyDeleteSql, s.Id)
}

func (s *ImsVoiceReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_voice_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVoiceReply) Get() (err error) {
	var tmp *ImsVoiceReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVoiceReplyScan(rows)
		return
	}, ImsVoiceReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVoiceReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVoiceReplyInsertSql, s.Rid, s.Title, s.Mediaid, s.Createtime)
}

func (s *ImsVoiceReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVoiceReplyDeleteSql, s.Id)
}

func (s *ImsVoiceReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_voice_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVoiceReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVoiceReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVoiceReplyScan(rows)
		return
	}, ImsVoiceReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFan) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanInsertSql, s.Pid, s.AccountId, s.AccountCategory, s.Openid, s.GroupId, s.UnionId, s.Mobile, s.Category, s.CategoryMsg, s.Subscribe, s.Nickname, s.Sex, s.Language, s.City, s.Province, s.Country, s.Avatar, s.SubscribeAt, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note, s.RecommendQrCode)
}

func (s *ImsVxFan) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanDeleteSql, s.Id)
}

func (s *ImsVxFan) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFan) Get() (err error) {
	var tmp *ImsVxFan
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanScan(rows)
		return
	}, ImsVxFanSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFan) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanInsertSql, s.Pid, s.AccountId, s.AccountCategory, s.Openid, s.GroupId, s.UnionId, s.Mobile, s.Category, s.CategoryMsg, s.Subscribe, s.Nickname, s.Sex, s.Language, s.City, s.Province, s.Country, s.Avatar, s.SubscribeAt, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note, s.RecommendQrCode)
}

func (s *ImsVxFan) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanDeleteSql, s.Id)
}

func (s *ImsVxFan) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFan) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFan
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanScan(rows)
		return
	}, ImsVxFanSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanAccount) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanAccountInsertSql, s.Uid, s.Balance, s.Prepare, s.Withdraw, s.LastWithdrawAt, s.LastWithdrawDay, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanAccount) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanAccountDeleteSql, s.Id)
}

func (s *ImsVxFanAccount) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_account` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanAccount) Get() (err error) {
	var tmp *ImsVxFanAccount
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanAccountScan(rows)
		return
	}, ImsVxFanAccountSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanAccount) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanAccountInsertSql, s.Uid, s.Balance, s.Prepare, s.Withdraw, s.LastWithdrawAt, s.LastWithdrawDay, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanAccount) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanAccountDeleteSql, s.Id)
}

func (s *ImsVxFanAccount) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_account` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanAccount) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanAccount
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanAccountScan(rows)
		return
	}, ImsVxFanAccountSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanAddress) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanAddressInsertSql, s.Uid, s.Consignee, s.Mobile, s.Address1, s.Address2, s.Address3, s.Address4, s.PostalCode, s.IsDefault, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.DeletedAt, s.Note)
}

func (s *ImsVxFanAddress) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanAddressDeleteSql, s.Id)
}

func (s *ImsVxFanAddress) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_address` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanAddress) Get() (err error) {
	var tmp *ImsVxFanAddress
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanAddressScan(rows)
		return
	}, ImsVxFanAddressSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanAddress) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanAddressInsertSql, s.Uid, s.Consignee, s.Mobile, s.Address1, s.Address2, s.Address3, s.Address4, s.PostalCode, s.IsDefault, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.DeletedAt, s.Note)
}

func (s *ImsVxFanAddress) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanAddressDeleteSql, s.Id)
}

func (s *ImsVxFanAddress) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_address` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanAddress) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanAddress
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanAddressScan(rows)
		return
	}, ImsVxFanAddressSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanBalanceLog) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanBalanceLogInsertSql, s.Uid, s.Val1, s.Val2, s.Val3, s.Category, s.CategoryMsg, s.OrderId, s.OrderNo, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanBalanceLog) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanBalanceLogDeleteSql, s.Id)
}

func (s *ImsVxFanBalanceLog) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_balance_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanBalanceLog) Get() (err error) {
	var tmp *ImsVxFanBalanceLog
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanBalanceLogScan(rows)
		return
	}, ImsVxFanBalanceLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanBalanceLog) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanBalanceLogInsertSql, s.Uid, s.Val1, s.Val2, s.Val3, s.Category, s.CategoryMsg, s.OrderId, s.OrderNo, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanBalanceLog) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanBalanceLogDeleteSql, s.Id)
}

func (s *ImsVxFanBalanceLog) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_balance_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanBalanceLog) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanBalanceLog
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanBalanceLogScan(rows)
		return
	}, ImsVxFanBalanceLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanGoods) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanGoodsInsertSql, s.AccountId, s.Avatar, s.Name, s.Title, s.TitleDesc, s.Label, s.Details, s.Video, s.PriceNominal, s.Price, s.Postage, s.Stock, s.Sold, s.Serial, s.Category, s.CategoryMsg, s.Uid, s.DeliveryAddress, s.Recommend, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.DeletedAt, s.Note)
}

func (s *ImsVxFanGoods) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanGoodsDeleteSql, s.Id)
}

func (s *ImsVxFanGoods) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_goods` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanGoods) Get() (err error) {
	var tmp *ImsVxFanGoods
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanGoodsScan(rows)
		return
	}, ImsVxFanGoodsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanGoods) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanGoodsInsertSql, s.AccountId, s.Avatar, s.Name, s.Title, s.TitleDesc, s.Label, s.Details, s.Video, s.PriceNominal, s.Price, s.Postage, s.Stock, s.Sold, s.Serial, s.Category, s.CategoryMsg, s.Uid, s.DeliveryAddress, s.Recommend, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.DeletedAt, s.Note)
}

func (s *ImsVxFanGoods) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanGoodsDeleteSql, s.Id)
}

func (s *ImsVxFanGoods) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_goods` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanGoods) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanGoods
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanGoodsScan(rows)
		return
	}, ImsVxFanGoodsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanGoodsRule) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanGoodsRuleInsertSql, s.GoodsId, s.Attributes, s.Serial, s.Avatar, s.Details, s.PriceNominal, s.Price, s.Stock, s.Sold, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.DeletedAt, s.Note)
}

func (s *ImsVxFanGoodsRule) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanGoodsRuleDeleteSql, s.Id)
}

func (s *ImsVxFanGoodsRule) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_goods_rule` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanGoodsRule) Get() (err error) {
	var tmp *ImsVxFanGoodsRule
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanGoodsRuleScan(rows)
		return
	}, ImsVxFanGoodsRuleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanGoodsRule) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanGoodsRuleInsertSql, s.GoodsId, s.Attributes, s.Serial, s.Avatar, s.Details, s.PriceNominal, s.Price, s.Stock, s.Sold, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.DeletedAt, s.Note)
}

func (s *ImsVxFanGoodsRule) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanGoodsRuleDeleteSql, s.Id)
}

func (s *ImsVxFanGoodsRule) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_goods_rule` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanGoodsRule) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanGoodsRule
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanGoodsRuleScan(rows)
		return
	}, ImsVxFanGoodsRuleSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanOrderBagQuotaCard) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanOrderBagQuotaCardInsertSql, s.Uid, s.AccountId, s.GoodsId, s.GoodsNum, s.GoodsPrice, s.GoodsTitle, s.OrderNo, s.Amount, s.PayAmount, s.PayType, s.PayState, s.PayNo, s.PayDetails, s.GoodsDetails, s.UserDetails, s.State, s.State1, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanOrderBagQuotaCard) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanOrderBagQuotaCardDeleteSql, s.Id)
}

func (s *ImsVxFanOrderBagQuotaCard) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_order_bag_quota_card` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanOrderBagQuotaCard) Get() (err error) {
	var tmp *ImsVxFanOrderBagQuotaCard
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanOrderBagQuotaCardScan(rows)
		return
	}, ImsVxFanOrderBagQuotaCardSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanOrderBagQuotaCard) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanOrderBagQuotaCardInsertSql, s.Uid, s.AccountId, s.GoodsId, s.GoodsNum, s.GoodsPrice, s.GoodsTitle, s.OrderNo, s.Amount, s.PayAmount, s.PayType, s.PayState, s.PayNo, s.PayDetails, s.GoodsDetails, s.UserDetails, s.State, s.State1, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanOrderBagQuotaCard) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanOrderBagQuotaCardDeleteSql, s.Id)
}

func (s *ImsVxFanOrderBagQuotaCard) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_order_bag_quota_card` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanOrderBagQuotaCard) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanOrderBagQuotaCard
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanOrderBagQuotaCardScan(rows)
		return
	}, ImsVxFanOrderBagQuotaCardSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanOrderGoods) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanOrderGoodsInsertSql, s.Uid, s.AccountId, s.GoodsId, s.GoodsNum, s.GoodsPrice, s.OrderNo, s.Amount, s.DiscountAmount, s.Discount, s.PayAmount, s.PayType, s.PayState, s.PayNo, s.PayDetails, s.UserDetails, s.GoodsDetails, s.AddressDetails, s.AddressId, s.Address, s.State, s.State1, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanOrderGoods) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanOrderGoodsDeleteSql, s.Id)
}

func (s *ImsVxFanOrderGoods) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_order_goods` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanOrderGoods) Get() (err error) {
	var tmp *ImsVxFanOrderGoods
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanOrderGoodsScan(rows)
		return
	}, ImsVxFanOrderGoodsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanOrderGoods) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanOrderGoodsInsertSql, s.Uid, s.AccountId, s.GoodsId, s.GoodsNum, s.GoodsPrice, s.OrderNo, s.Amount, s.DiscountAmount, s.Discount, s.PayAmount, s.PayType, s.PayState, s.PayNo, s.PayDetails, s.UserDetails, s.GoodsDetails, s.AddressDetails, s.AddressId, s.Address, s.State, s.State1, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanOrderGoods) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanOrderGoodsDeleteSql, s.Id)
}

func (s *ImsVxFanOrderGoods) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_order_goods` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanOrderGoods) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanOrderGoods
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanOrderGoodsScan(rows)
		return
	}, ImsVxFanOrderGoodsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanOrderRightCard) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanOrderRightCardInsertSql, s.Uid, s.AccountId, s.GoodsId, s.GoodsNum, s.GoodsPrice, s.GoodsTitle, s.OrderNo, s.Amount, s.PayAmount, s.PayType, s.PayState, s.PayNo, s.PayDetails, s.GoodsDetails, s.UserDetails, s.State, s.State1, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanOrderRightCard) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanOrderRightCardDeleteSql, s.Id)
}

func (s *ImsVxFanOrderRightCard) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_order_right_card` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanOrderRightCard) Get() (err error) {
	var tmp *ImsVxFanOrderRightCard
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanOrderRightCardScan(rows)
		return
	}, ImsVxFanOrderRightCardSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanOrderRightCard) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanOrderRightCardInsertSql, s.Uid, s.AccountId, s.GoodsId, s.GoodsNum, s.GoodsPrice, s.GoodsTitle, s.OrderNo, s.Amount, s.PayAmount, s.PayType, s.PayState, s.PayNo, s.PayDetails, s.GoodsDetails, s.UserDetails, s.State, s.State1, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanOrderRightCard) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanOrderRightCardDeleteSql, s.Id)
}

func (s *ImsVxFanOrderRightCard) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_order_right_card` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanOrderRightCard) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanOrderRightCard
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanOrderRightCardScan(rows)
		return
	}, ImsVxFanOrderRightCardSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanPrepareLog) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanPrepareLogInsertSql, s.Uid, s.Val1, s.Val2, s.Val3, s.Category, s.CategoryMsg, s.OrderId, s.OrderNo, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanPrepareLog) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanPrepareLogDeleteSql, s.Id)
}

func (s *ImsVxFanPrepareLog) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_prepare_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanPrepareLog) Get() (err error) {
	var tmp *ImsVxFanPrepareLog
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanPrepareLogScan(rows)
		return
	}, ImsVxFanPrepareLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanPrepareLog) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanPrepareLogInsertSql, s.Uid, s.Val1, s.Val2, s.Val3, s.Category, s.CategoryMsg, s.OrderId, s.OrderNo, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanPrepareLog) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanPrepareLogDeleteSql, s.Id)
}

func (s *ImsVxFanPrepareLog) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_prepare_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanPrepareLog) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanPrepareLog
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanPrepareLogScan(rows)
		return
	}, ImsVxFanPrepareLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanRightCard) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanRightCardInsertSql, s.AccountId, s.Name, s.Money, s.Avatar, s.Title, s.Details, s.Multiple, s.Vip, s.Discount, s.Serial, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanRightCard) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanRightCardDeleteSql, s.Id)
}

func (s *ImsVxFanRightCard) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_right_card` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanRightCard) Get() (err error) {
	var tmp *ImsVxFanRightCard
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanRightCardScan(rows)
		return
	}, ImsVxFanRightCardSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanRightCard) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanRightCardInsertSql, s.AccountId, s.Name, s.Money, s.Avatar, s.Title, s.Details, s.Multiple, s.Vip, s.Discount, s.Serial, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanRightCard) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanRightCardDeleteSql, s.Id)
}

func (s *ImsVxFanRightCard) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_right_card` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanRightCard) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanRightCard
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanRightCardScan(rows)
		return
	}, ImsVxFanRightCardSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanRightCardUpgrade) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanRightCardUpgradeInsertSql, s.Uid, s.UpgradeRatio, s.FromUid, s.RightCardAmount, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note, s.Cards)
}

func (s *ImsVxFanRightCardUpgrade) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanRightCardUpgradeDeleteSql, s.Id)
}

func (s *ImsVxFanRightCardUpgrade) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_right_card_upgrade` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanRightCardUpgrade) Get() (err error) {
	var tmp *ImsVxFanRightCardUpgrade
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanRightCardUpgradeScan(rows)
		return
	}, ImsVxFanRightCardUpgradeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanRightCardUpgrade) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanRightCardUpgradeInsertSql, s.Uid, s.UpgradeRatio, s.FromUid, s.RightCardAmount, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note, s.Cards)
}

func (s *ImsVxFanRightCardUpgrade) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanRightCardUpgradeDeleteSql, s.Id)
}

func (s *ImsVxFanRightCardUpgrade) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_right_card_upgrade` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanRightCardUpgrade) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanRightCardUpgrade
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanRightCardUpgradeScan(rows)
		return
	}, ImsVxFanRightCardUpgradeSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanRightCardUsed) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanRightCardUsedInsertSql, s.Cid, s.Uid, s.Serial, s.Name, s.Money, s.Multiple, s.Amount, s.Days, s.AmountRemain, s.DaysRemain, s.DayRatio, s.DayMoney, s.AmountAchieved, s.AmountAchievedDays, s.Discount, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note, s.Tips, s.Card)
}

func (s *ImsVxFanRightCardUsed) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanRightCardUsedDeleteSql, s.Id)
}

func (s *ImsVxFanRightCardUsed) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_right_card_used` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanRightCardUsed) Get() (err error) {
	var tmp *ImsVxFanRightCardUsed
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanRightCardUsedScan(rows)
		return
	}, ImsVxFanRightCardUsedSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanRightCardUsed) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanRightCardUsedInsertSql, s.Cid, s.Uid, s.Serial, s.Name, s.Money, s.Multiple, s.Amount, s.Days, s.AmountRemain, s.DaysRemain, s.DayRatio, s.DayMoney, s.AmountAchieved, s.AmountAchievedDays, s.Discount, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note, s.Tips, s.Card)
}

func (s *ImsVxFanRightCardUsed) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanRightCardUsedDeleteSql, s.Id)
}

func (s *ImsVxFanRightCardUsed) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_right_card_used` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanRightCardUsed) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanRightCardUsed
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanRightCardUsedScan(rows)
		return
	}, ImsVxFanRightCardUsedSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanRightCardUsing) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanRightCardUsingInsertSql, s.Cid, s.Uid, s.Serial, s.Name, s.Money, s.Multiple, s.Amount, s.Days, s.AmountRemain, s.DaysRemain, s.DayRatio, s.DayMoney, s.AmountAchieved, s.AmountAchievedDays, s.Discount, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note, s.Tips, s.Card)
}

func (s *ImsVxFanRightCardUsing) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanRightCardUsingDeleteSql, s.Id)
}

func (s *ImsVxFanRightCardUsing) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_right_card_using` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanRightCardUsing) Get() (err error) {
	var tmp *ImsVxFanRightCardUsing
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanRightCardUsingScan(rows)
		return
	}, ImsVxFanRightCardUsingSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanRightCardUsing) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanRightCardUsingInsertSql, s.Cid, s.Uid, s.Serial, s.Name, s.Money, s.Multiple, s.Amount, s.Days, s.AmountRemain, s.DaysRemain, s.DayRatio, s.DayMoney, s.AmountAchieved, s.AmountAchievedDays, s.Discount, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note, s.Tips, s.Card)
}

func (s *ImsVxFanRightCardUsing) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanRightCardUsingDeleteSql, s.Id)
}

func (s *ImsVxFanRightCardUsing) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_right_card_using` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanRightCardUsing) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanRightCardUsing
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanRightCardUsingScan(rows)
		return
	}, ImsVxFanRightCardUsingSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanSign) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanSignInsertSql, s.Uid, s.Val1, s.Val2, s.Val3, s.Day, s.ContinuedDays, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanSign) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanSignDeleteSql, s.Id)
}

func (s *ImsVxFanSign) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_sign` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanSign) Get() (err error) {
	var tmp *ImsVxFanSign
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanSignScan(rows)
		return
	}, ImsVxFanSignSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanSign) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanSignInsertSql, s.Uid, s.Val1, s.Val2, s.Val3, s.Day, s.ContinuedDays, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanSign) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanSignDeleteSql, s.Id)
}

func (s *ImsVxFanSign) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_sign` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanSign) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanSign
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanSignScan(rows)
		return
	}, ImsVxFanSignSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanSignAdvertising) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanSignAdvertisingInsertSql, s.AccountId, s.Name, s.Avatar, s.Video, s.Details, s.Duration, s.Category, s.CategoryMsg, s.Uid, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanSignAdvertising) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanSignAdvertisingDeleteSql, s.Id)
}

func (s *ImsVxFanSignAdvertising) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_sign_advertising` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanSignAdvertising) Get() (err error) {
	var tmp *ImsVxFanSignAdvertising
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanSignAdvertisingScan(rows)
		return
	}, ImsVxFanSignAdvertisingSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanSignAdvertising) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanSignAdvertisingInsertSql, s.AccountId, s.Name, s.Avatar, s.Video, s.Details, s.Duration, s.Category, s.CategoryMsg, s.Uid, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanSignAdvertising) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanSignAdvertisingDeleteSql, s.Id)
}

func (s *ImsVxFanSignAdvertising) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_sign_advertising` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanSignAdvertising) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanSignAdvertising
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanSignAdvertisingScan(rows)
		return
	}, ImsVxFanSignAdvertisingSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanVip) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanVipInsertSql, s.Uid, s.Vip, s.VipMsg, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.ExpiredAt, s.Note)
}

func (s *ImsVxFanVip) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanVipDeleteSql, s.Id)
}

func (s *ImsVxFanVip) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_vip` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanVip) Get() (err error) {
	var tmp *ImsVxFanVip
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanVipScan(rows)
		return
	}, ImsVxFanVipSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanVip) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanVipInsertSql, s.Uid, s.Vip, s.VipMsg, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.ExpiredAt, s.Note)
}

func (s *ImsVxFanVip) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanVipDeleteSql, s.Id)
}

func (s *ImsVxFanVip) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_vip` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanVip) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanVip
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanVipScan(rows)
		return
	}, ImsVxFanVipSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanWithdrawLog) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxFanWithdrawLogInsertSql, s.Uid, s.Val1, s.Val2, s.Val3, s.Category, s.CategoryMsg, s.OrderId, s.OrderNo, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanWithdrawLog) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxFanWithdrawLogDeleteSql, s.Id)
}

func (s *ImsVxFanWithdrawLog) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_withdraw_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanWithdrawLog) Get() (err error) {
	var tmp *ImsVxFanWithdrawLog
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanWithdrawLogScan(rows)
		return
	}, ImsVxFanWithdrawLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxFanWithdrawLog) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxFanWithdrawLogInsertSql, s.Uid, s.Val1, s.Val2, s.Val3, s.Category, s.CategoryMsg, s.OrderId, s.OrderNo, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxFanWithdrawLog) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxFanWithdrawLogDeleteSql, s.Id)
}

func (s *ImsVxFanWithdrawLog) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_fan_withdraw_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxFanWithdrawLog) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxFanWithdrawLog
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxFanWithdrawLogScan(rows)
		return
	}, ImsVxFanWithdrawLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxRedBag) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxRedBagInsertSql, s.Uid, s.AccountId, s.AccountAppId, s.MchId, s.Openid, s.Amount, s.SendListId, s.MchNo, s.ReturnCode, s.ReturnMsg, s.ResultCode, s.ErrCode, s.ErrCodeDes, s.Day, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxRedBag) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxRedBagDeleteSql, s.Id)
}

func (s *ImsVxRedBag) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_red_bag` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxRedBag) Get() (err error) {
	var tmp *ImsVxRedBag
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxRedBagScan(rows)
		return
	}, ImsVxRedBagSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxRedBag) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxRedBagInsertSql, s.Uid, s.AccountId, s.AccountAppId, s.MchId, s.Openid, s.Amount, s.SendListId, s.MchNo, s.ReturnCode, s.ReturnMsg, s.ResultCode, s.ErrCode, s.ErrCodeDes, s.Day, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxRedBag) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxRedBagDeleteSql, s.Id)
}

func (s *ImsVxRedBag) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_red_bag` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxRedBag) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxRedBag
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxRedBagScan(rows)
		return
	}, ImsVxRedBagSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxRedBagQuotaCard) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsVxRedBagQuotaCardInsertSql, s.AccountId, s.Name, s.Avatar, s.Title, s.Details, s.Money, s.Prepare, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxRedBagQuotaCard) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsVxRedBagQuotaCardDeleteSql, s.Id)
}

func (s *ImsVxRedBagQuotaCard) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_vx_red_bag_quota_card` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxRedBagQuotaCard) Get() (err error) {
	var tmp *ImsVxRedBagQuotaCard
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxRedBagQuotaCardScan(rows)
		return
	}, ImsVxRedBagQuotaCardSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsVxRedBagQuotaCard) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsVxRedBagQuotaCardInsertSql, s.AccountId, s.Name, s.Avatar, s.Title, s.Details, s.Money, s.Prepare, s.State, s.StateMsg, s.CreatedAt, s.UpdatedAt, s.Note)
}

func (s *ImsVxRedBagQuotaCard) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsVxRedBagQuotaCardDeleteSql, s.Id)
}

func (s *ImsVxRedBagQuotaCard) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_vx_red_bag_quota_card` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsVxRedBagQuotaCard) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsVxRedBagQuotaCard
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsVxRedBagQuotaCardScan(rows)
		return
	}, ImsVxRedBagQuotaCardSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWechatAttachment) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsWechatAttachmentInsertSql, s.Uniacid, s.Acid, s.Uid, s.Filename, s.Attachment, s.MediaId, s.Width, s.Height, s.Type, s.Model, s.Tag, s.Createtime, s.ModuleUploadDir, s.GroupId)
}

func (s *ImsWechatAttachment) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsWechatAttachmentDeleteSql, s.Id)
}

func (s *ImsWechatAttachment) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_wechat_attachment` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWechatAttachment) Get() (err error) {
	var tmp *ImsWechatAttachment
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWechatAttachmentScan(rows)
		return
	}, ImsWechatAttachmentSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWechatAttachment) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsWechatAttachmentInsertSql, s.Uniacid, s.Acid, s.Uid, s.Filename, s.Attachment, s.MediaId, s.Width, s.Height, s.Type, s.Model, s.Tag, s.Createtime, s.ModuleUploadDir, s.GroupId)
}

func (s *ImsWechatAttachment) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsWechatAttachmentDeleteSql, s.Id)
}

func (s *ImsWechatAttachment) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_wechat_attachment` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWechatAttachment) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsWechatAttachment
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWechatAttachmentScan(rows)
		return
	}, ImsWechatAttachmentSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWechatNews) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsWechatNewsInsertSql, s.Uniacid, s.AttachId, s.ThumbMediaId, s.ThumbUrl, s.Title, s.Author, s.Digest, s.Content, s.ContentSourceUrl, s.ShowCoverPic, s.Url, s.Displayorder, s.NeedOpenComment, s.OnlyFansCanComment)
}

func (s *ImsWechatNews) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsWechatNewsDeleteSql, s.Id)
}

func (s *ImsWechatNews) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_wechat_news` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWechatNews) Get() (err error) {
	var tmp *ImsWechatNews
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWechatNewsScan(rows)
		return
	}, ImsWechatNewsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWechatNews) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsWechatNewsInsertSql, s.Uniacid, s.AttachId, s.ThumbMediaId, s.ThumbUrl, s.Title, s.Author, s.Digest, s.Content, s.ContentSourceUrl, s.ShowCoverPic, s.Url, s.Displayorder, s.NeedOpenComment, s.OnlyFansCanComment)
}

func (s *ImsWechatNews) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsWechatNewsDeleteSql, s.Id)
}

func (s *ImsWechatNews) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_wechat_news` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWechatNews) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsWechatNews
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWechatNewsScan(rows)
		return
	}, ImsWechatNewsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxappGeneralAnalysis) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsWxappGeneralAnalysisInsertSql, s.Uniacid, s.SessionCnt, s.VisitPv, s.VisitUv, s.VisitUvNew, s.Type, s.StayTimeUv, s.StayTimeSession, s.VisitDepth, s.RefDate)
}

func (s *ImsWxappGeneralAnalysis) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsWxappGeneralAnalysisDeleteSql, s.Id)
}

func (s *ImsWxappGeneralAnalysis) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_wxapp_general_analysis` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxappGeneralAnalysis) Get() (err error) {
	var tmp *ImsWxappGeneralAnalysis
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxappGeneralAnalysisScan(rows)
		return
	}, ImsWxappGeneralAnalysisSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxappGeneralAnalysis) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsWxappGeneralAnalysisInsertSql, s.Uniacid, s.SessionCnt, s.VisitPv, s.VisitUv, s.VisitUvNew, s.Type, s.StayTimeUv, s.StayTimeSession, s.VisitDepth, s.RefDate)
}

func (s *ImsWxappGeneralAnalysis) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsWxappGeneralAnalysisDeleteSql, s.Id)
}

func (s *ImsWxappGeneralAnalysis) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_wxapp_general_analysis` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxappGeneralAnalysis) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsWxappGeneralAnalysis
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxappGeneralAnalysisScan(rows)
		return
	}, ImsWxappGeneralAnalysisSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxappRegisterVersion) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsWxappRegisterVersionInsertSql, s.Uniacid, s.VersionId, s.Auditid, s.Version, s.Description, s.Status, s.Reason, s.UploadTime, s.AuditInfo, s.SubmitInfo, s.Developer)
}

func (s *ImsWxappRegisterVersion) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsWxappRegisterVersionDeleteSql, s.Id)
}

func (s *ImsWxappRegisterVersion) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_wxapp_register_version` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxappRegisterVersion) Get() (err error) {
	var tmp *ImsWxappRegisterVersion
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxappRegisterVersionScan(rows)
		return
	}, ImsWxappRegisterVersionSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxappRegisterVersion) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsWxappRegisterVersionInsertSql, s.Uniacid, s.VersionId, s.Auditid, s.Version, s.Description, s.Status, s.Reason, s.UploadTime, s.AuditInfo, s.SubmitInfo, s.Developer)
}

func (s *ImsWxappRegisterVersion) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsWxappRegisterVersionDeleteSql, s.Id)
}

func (s *ImsWxappRegisterVersion) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_wxapp_register_version` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxappRegisterVersion) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsWxappRegisterVersion
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxappRegisterVersionScan(rows)
		return
	}, ImsWxappRegisterVersionSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxappUndocodeauditLog) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsWxappUndocodeauditLogInsertSql, s.Uniacid, s.VersionId, s.Auditid, s.RevokeTime)
}

func (s *ImsWxappUndocodeauditLog) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsWxappUndocodeauditLogDeleteSql, s.Id)
}

func (s *ImsWxappUndocodeauditLog) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_wxapp_undocodeaudit_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxappUndocodeauditLog) Get() (err error) {
	var tmp *ImsWxappUndocodeauditLog
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxappUndocodeauditLogScan(rows)
		return
	}, ImsWxappUndocodeauditLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxappUndocodeauditLog) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsWxappUndocodeauditLogInsertSql, s.Uniacid, s.VersionId, s.Auditid, s.RevokeTime)
}

func (s *ImsWxappUndocodeauditLog) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsWxappUndocodeauditLogDeleteSql, s.Id)
}

func (s *ImsWxappUndocodeauditLog) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_wxapp_undocodeaudit_log` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxappUndocodeauditLog) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsWxappUndocodeauditLog
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxappUndocodeauditLogScan(rows)
		return
	}, ImsWxappUndocodeauditLogSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxappVersions) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsWxappVersionsInsertSql, s.Uniacid, s.Multiid, s.Version, s.Description, s.Modules, s.DesignMethod, s.Template, s.Quickmenu, s.Createtime, s.Type, s.EntryId, s.Appjson, s.DefaultAppjson, s.UseDefault, s.LastModules, s.Tominiprogram, s.UploadTime)
}

func (s *ImsWxappVersions) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsWxappVersionsDeleteSql, s.Id)
}

func (s *ImsWxappVersions) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_wxapp_versions` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxappVersions) Get() (err error) {
	var tmp *ImsWxappVersions
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxappVersionsScan(rows)
		return
	}, ImsWxappVersionsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxappVersions) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsWxappVersionsInsertSql, s.Uniacid, s.Multiid, s.Version, s.Description, s.Modules, s.DesignMethod, s.Template, s.Quickmenu, s.Createtime, s.Type, s.EntryId, s.Appjson, s.DefaultAppjson, s.UseDefault, s.LastModules, s.Tominiprogram, s.UploadTime)
}

func (s *ImsWxappVersions) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsWxappVersionsDeleteSql, s.Id)
}

func (s *ImsWxappVersions) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_wxapp_versions` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxappVersions) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsWxappVersions
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxappVersionsScan(rows)
		return
	}, ImsWxappVersionsSelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxcardReply) Add() (int64, error) {
	return gomysql.Db2().RightCreate(ImsWxcardReplyInsertSql, s.Rid, s.Title, s.CardId, s.Cid, s.BrandName, s.LogoUrl, s.Success, s.Error)
}

func (s *ImsWxcardReply) Del() (int64, error) {
	return gomysql.Db2().RightExecute(ImsWxcardReplyDeleteSql, s.Id)
}

func (s *ImsWxcardReply) Mod(m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return gomysql.Db2().RightExecute(fmt.Sprintf("UPDATE `ims_wxcard_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxcardReply) Get() (err error) {
	var tmp *ImsWxcardReply
	err = gomysql.Db2().RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxcardReplyScan(rows)
		return
	}, ImsWxcardReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}

func (s *ImsWxcardReply) Add1(e *gomysql.Execs) (int64, error) {
	return e.RightCreate(ImsWxcardReplyInsertSql, s.Rid, s.Title, s.CardId, s.Cid, s.BrandName, s.LogoUrl, s.Success, s.Error)
}

func (s *ImsWxcardReply) Del1(e *gomysql.Execs) (int64, error) {
	return e.RightExecute(ImsWxcardReplyDeleteSql, s.Id)
}

func (s *ImsWxcardReply) Mod1(e *gomysql.Execs, m map[string]interface{}) (int64, error) {
	length := len(m)
	if length == 0 {
		return 0, nil
	}
	cols := make([]string, length, length)
	args := make([]interface{}, length+1, length+1)
	i := 0
	for k, v := range m {
		cols[i] = fmt.Sprintf("`%s` = ?", k)
		args[i] = v
		i++
	}
	args[i] = s.Id
	return e.RightExecute(fmt.Sprintf("UPDATE `ims_wxcard_reply` SET %s WHERE ( `id` = ? );", strings.Join(cols, ", ")), args...)
}

func (s *ImsWxcardReply) Get1(e *gomysql.Execs) (err error) {
	var tmp *ImsWxcardReply
	err = e.RightQuery(func(rows *sql.Rows) (err error) {
		tmp, err = ImsWxcardReplyScan(rows)
		return
	}, ImsWxcardReplySelectSql, s.Id)
	if err != nil {
		return
	}
	if tmp != nil {
		*s = *tmp
	}
	return
}
