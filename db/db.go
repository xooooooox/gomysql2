package db

import (
	"database/sql"
)

const (
	ImsAccountInsertSql                     = "INSERT INTO `ims_account` ( `uniacid`, `hash`, `type`, `isconnect`, `isdeleted`, `endtime`, `send_account_expire_status`, `send_api_expire_status` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsAccountDeleteSql                     = "DELETE FROM `ims_account` WHERE ( `acid` = ? );"
	ImsAccountUpdateSql                     = "UPDATE `ims_account` SET `uniacid` = ?, `hash` = ?, `type` = ?, `isconnect` = ?, `isdeleted` = ?, `endtime` = ?, `send_account_expire_status` = ?, `send_api_expire_status` = ? WHERE ( `acid` = ? );"
	ImsAccountSelectSql                     = "SELECT `acid`, `uniacid`, `hash`, `type`, `isconnect`, `isdeleted`, `endtime`, `send_account_expire_status`, `send_api_expire_status` FROM `ims_account` WHERE ( `acid` = ? ) ORDER BY `acid` DESC LIMIT 0, 1;"
	ImsAccountAliappInsertSql               = "INSERT INTO `ims_account_aliapp` ( `uniacid`, `level`, `name`, `description`, `key` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsAccountAliappDeleteSql               = "DELETE FROM `ims_account_aliapp` WHERE ( `acid` = ? );"
	ImsAccountAliappUpdateSql               = "UPDATE `ims_account_aliapp` SET `uniacid` = ?, `level` = ?, `name` = ?, `description` = ?, `key` = ? WHERE ( `acid` = ? );"
	ImsAccountAliappSelectSql               = "SELECT `acid`, `uniacid`, `level`, `name`, `description`, `key` FROM `ims_account_aliapp` WHERE ( `acid` = ? ) ORDER BY `acid` DESC LIMIT 0, 1;"
	ImsAccountBaiduappInsertSql             = "INSERT INTO `ims_account_baiduapp` ( `uniacid`, `name`, `appid`, `key`, `secret`, `description` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsAccountBaiduappDeleteSql             = "DELETE FROM `ims_account_baiduapp` WHERE ( `acid` = ? );"
	ImsAccountBaiduappUpdateSql             = "UPDATE `ims_account_baiduapp` SET `uniacid` = ?, `name` = ?, `appid` = ?, `key` = ?, `secret` = ?, `description` = ? WHERE ( `acid` = ? );"
	ImsAccountBaiduappSelectSql             = "SELECT `acid`, `uniacid`, `name`, `appid`, `key`, `secret`, `description` FROM `ims_account_baiduapp` WHERE ( `acid` = ? ) ORDER BY `acid` DESC LIMIT 0, 1;"
	ImsAccountPhoneappInsertSql             = "INSERT INTO `ims_account_phoneapp` ( `uniacid`, `name` ) VALUES ( ?, ? );"
	ImsAccountPhoneappDeleteSql             = "DELETE FROM `ims_account_phoneapp` WHERE ( `acid` = ? );"
	ImsAccountPhoneappUpdateSql             = "UPDATE `ims_account_phoneapp` SET `uniacid` = ?, `name` = ? WHERE ( `acid` = ? );"
	ImsAccountPhoneappSelectSql             = "SELECT `acid`, `uniacid`, `name` FROM `ims_account_phoneapp` WHERE ( `acid` = ? ) ORDER BY `acid` DESC LIMIT 0, 1;"
	ImsAccountToutiaoappInsertSql           = "INSERT INTO `ims_account_toutiaoapp` ( `uniacid`, `name`, `appid`, `key`, `secret`, `description` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsAccountToutiaoappDeleteSql           = "DELETE FROM `ims_account_toutiaoapp` WHERE ( `acid` = ? );"
	ImsAccountToutiaoappUpdateSql           = "UPDATE `ims_account_toutiaoapp` SET `uniacid` = ?, `name` = ?, `appid` = ?, `key` = ?, `secret` = ?, `description` = ? WHERE ( `acid` = ? );"
	ImsAccountToutiaoappSelectSql           = "SELECT `acid`, `uniacid`, `name`, `appid`, `key`, `secret`, `description` FROM `ims_account_toutiaoapp` WHERE ( `acid` = ? ) ORDER BY `acid` DESC LIMIT 0, 1;"
	ImsAccountWebappInsertSql               = "INSERT INTO `ims_account_webapp` ( `uniacid`, `name` ) VALUES ( ?, ? );"
	ImsAccountWebappDeleteSql               = "DELETE FROM `ims_account_webapp` WHERE ( `acid` = ? );"
	ImsAccountWebappUpdateSql               = "UPDATE `ims_account_webapp` SET `uniacid` = ?, `name` = ? WHERE ( `acid` = ? );"
	ImsAccountWebappSelectSql               = "SELECT `acid`, `uniacid`, `name` FROM `ims_account_webapp` WHERE ( `acid` = ? ) ORDER BY `acid` DESC LIMIT 0, 1;"
	ImsAccountWechatsInsertSql              = "INSERT INTO `ims_account_wechats` ( `uniacid`, `token`, `encodingaeskey`, `level`, `name`, `account`, `original`, `signature`, `country`, `province`, `city`, `username`, `password`, `lastupdate`, `key`, `secret`, `styleid`, `subscribeurl`, `auth_refresh_token` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsAccountWechatsDeleteSql              = "DELETE FROM `ims_account_wechats` WHERE ( `acid` = ? );"
	ImsAccountWechatsUpdateSql              = "UPDATE `ims_account_wechats` SET `uniacid` = ?, `token` = ?, `encodingaeskey` = ?, `level` = ?, `name` = ?, `account` = ?, `original` = ?, `signature` = ?, `country` = ?, `province` = ?, `city` = ?, `username` = ?, `password` = ?, `lastupdate` = ?, `key` = ?, `secret` = ?, `styleid` = ?, `subscribeurl` = ?, `auth_refresh_token` = ? WHERE ( `acid` = ? );"
	ImsAccountWechatsSelectSql              = "SELECT `acid`, `uniacid`, `token`, `encodingaeskey`, `level`, `name`, `account`, `original`, `signature`, `country`, `province`, `city`, `username`, `password`, `lastupdate`, `key`, `secret`, `styleid`, `subscribeurl`, `auth_refresh_token` FROM `ims_account_wechats` WHERE ( `acid` = ? ) ORDER BY `acid` DESC LIMIT 0, 1;"
	ImsAccountWxappInsertSql                = "INSERT INTO `ims_account_wxapp` ( `uniacid`, `token`, `encodingaeskey`, `level`, `account`, `original`, `key`, `secret`, `name`, `appdomain`, `auth_refresh_token` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsAccountWxappDeleteSql                = "DELETE FROM `ims_account_wxapp` WHERE ( `acid` = ? );"
	ImsAccountWxappUpdateSql                = "UPDATE `ims_account_wxapp` SET `uniacid` = ?, `token` = ?, `encodingaeskey` = ?, `level` = ?, `account` = ?, `original` = ?, `key` = ?, `secret` = ?, `name` = ?, `appdomain` = ?, `auth_refresh_token` = ? WHERE ( `acid` = ? );"
	ImsAccountWxappSelectSql                = "SELECT `acid`, `uniacid`, `token`, `encodingaeskey`, `level`, `account`, `original`, `key`, `secret`, `name`, `appdomain`, `auth_refresh_token` FROM `ims_account_wxapp` WHERE ( `acid` = ? ) ORDER BY `acid` DESC LIMIT 0, 1;"
	ImsAccountXzappInsertSql                = "INSERT INTO `ims_account_xzapp` ( `uniacid`, `name`, `original`, `lastupdate`, `styleid`, `createtime`, `token`, `encodingaeskey`, `xzapp_id`, `level`, `key`, `secret` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsAccountXzappDeleteSql                = "DELETE FROM `ims_account_xzapp` WHERE ( `acid` = ? );"
	ImsAccountXzappUpdateSql                = "UPDATE `ims_account_xzapp` SET `uniacid` = ?, `name` = ?, `original` = ?, `lastupdate` = ?, `styleid` = ?, `createtime` = ?, `token` = ?, `encodingaeskey` = ?, `xzapp_id` = ?, `level` = ?, `key` = ?, `secret` = ? WHERE ( `acid` = ? );"
	ImsAccountXzappSelectSql                = "SELECT `acid`, `uniacid`, `name`, `original`, `lastupdate`, `styleid`, `createtime`, `token`, `encodingaeskey`, `xzapp_id`, `level`, `key`, `secret` FROM `ims_account_xzapp` WHERE ( `acid` = ? ) ORDER BY `acid` DESC LIMIT 0, 1;"
	ImsActivityClerksInsertSql              = "INSERT INTO `ims_activity_clerks` ( `uniacid`, `uid`, `storeid`, `name`, `password`, `mobile`, `openid`, `nickname` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsActivityClerksDeleteSql              = "DELETE FROM `ims_activity_clerks` WHERE ( `id` = ? );"
	ImsActivityClerksUpdateSql              = "UPDATE `ims_activity_clerks` SET `uniacid` = ?, `uid` = ?, `storeid` = ?, `name` = ?, `password` = ?, `mobile` = ?, `openid` = ?, `nickname` = ? WHERE ( `id` = ? );"
	ImsActivityClerksSelectSql              = "SELECT `id`, `uniacid`, `uid`, `storeid`, `name`, `password`, `mobile`, `openid`, `nickname` FROM `ims_activity_clerks` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsActivityClerkMenuInsertSql           = "INSERT INTO `ims_activity_clerk_menu` ( `uniacid`, `displayorder`, `pid`, `group_name`, `title`, `icon`, `url`, `type`, `permission`, `system` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsActivityClerkMenuDeleteSql           = "DELETE FROM `ims_activity_clerk_menu` WHERE ( `id` = ? );"
	ImsActivityClerkMenuUpdateSql           = "UPDATE `ims_activity_clerk_menu` SET `uniacid` = ?, `displayorder` = ?, `pid` = ?, `group_name` = ?, `title` = ?, `icon` = ?, `url` = ?, `type` = ?, `permission` = ?, `system` = ? WHERE ( `id` = ? );"
	ImsActivityClerkMenuSelectSql           = "SELECT `id`, `uniacid`, `displayorder`, `pid`, `group_name`, `title`, `icon`, `url`, `type`, `permission`, `system` FROM `ims_activity_clerk_menu` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsAddressInsertSql                     = "INSERT INTO `ims_address` ( `pid`, `name`, `level` ) VALUES ( ?, ?, ? );"
	ImsAddressDeleteSql                     = "DELETE FROM `ims_address` WHERE ( `id` = ? );"
	ImsAddressUpdateSql                     = "UPDATE `ims_address` SET `pid` = ?, `name` = ?, `level` = ? WHERE ( `id` = ? );"
	ImsAddressSelectSql                     = "SELECT `id`, `pid`, `name`, `level` FROM `ims_address` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsArticleCategoryInsertSql             = "INSERT INTO `ims_article_category` ( `title`, `displayorder`, `type` ) VALUES ( ?, ?, ? );"
	ImsArticleCategoryDeleteSql             = "DELETE FROM `ims_article_category` WHERE ( `id` = ? );"
	ImsArticleCategoryUpdateSql             = "UPDATE `ims_article_category` SET `title` = ?, `displayorder` = ?, `type` = ? WHERE ( `id` = ? );"
	ImsArticleCategorySelectSql             = "SELECT `id`, `title`, `displayorder`, `type` FROM `ims_article_category` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsArticleCommentInsertSql              = "INSERT INTO `ims_article_comment` ( `articleid`, `parentid`, `uid`, `content`, `is_like`, `is_reply`, `like_num`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsArticleCommentDeleteSql              = "DELETE FROM `ims_article_comment` WHERE ( `id` = ? );"
	ImsArticleCommentUpdateSql              = "UPDATE `ims_article_comment` SET `articleid` = ?, `parentid` = ?, `uid` = ?, `content` = ?, `is_like` = ?, `is_reply` = ?, `like_num` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsArticleCommentSelectSql              = "SELECT `id`, `articleid`, `parentid`, `uid`, `content`, `is_like`, `is_reply`, `like_num`, `createtime` FROM `ims_article_comment` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsArticleNewsInsertSql                 = "INSERT INTO `ims_article_news` ( `cateid`, `title`, `content`, `thumb`, `source`, `author`, `displayorder`, `is_display`, `is_show_home`, `createtime`, `click` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsArticleNewsDeleteSql                 = "DELETE FROM `ims_article_news` WHERE ( `id` = ? );"
	ImsArticleNewsUpdateSql                 = "UPDATE `ims_article_news` SET `cateid` = ?, `title` = ?, `content` = ?, `thumb` = ?, `source` = ?, `author` = ?, `displayorder` = ?, `is_display` = ?, `is_show_home` = ?, `createtime` = ?, `click` = ? WHERE ( `id` = ? );"
	ImsArticleNewsSelectSql                 = "SELECT `id`, `cateid`, `title`, `content`, `thumb`, `source`, `author`, `displayorder`, `is_display`, `is_show_home`, `createtime`, `click` FROM `ims_article_news` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsArticleNoticeInsertSql               = "INSERT INTO `ims_article_notice` ( `cateid`, `title`, `content`, `displayorder`, `is_display`, `is_show_home`, `createtime`, `click`, `style`, `group` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsArticleNoticeDeleteSql               = "DELETE FROM `ims_article_notice` WHERE ( `id` = ? );"
	ImsArticleNoticeUpdateSql               = "UPDATE `ims_article_notice` SET `cateid` = ?, `title` = ?, `content` = ?, `displayorder` = ?, `is_display` = ?, `is_show_home` = ?, `createtime` = ?, `click` = ?, `style` = ?, `group` = ? WHERE ( `id` = ? );"
	ImsArticleNoticeSelectSql               = "SELECT `id`, `cateid`, `title`, `content`, `displayorder`, `is_display`, `is_show_home`, `createtime`, `click`, `style`, `group` FROM `ims_article_notice` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsArticleUnreadNoticeInsertSql         = "INSERT INTO `ims_article_unread_notice` ( `notice_id`, `uid`, `is_new` ) VALUES ( ?, ?, ? );"
	ImsArticleUnreadNoticeDeleteSql         = "DELETE FROM `ims_article_unread_notice` WHERE ( `id` = ? );"
	ImsArticleUnreadNoticeUpdateSql         = "UPDATE `ims_article_unread_notice` SET `notice_id` = ?, `uid` = ?, `is_new` = ? WHERE ( `id` = ? );"
	ImsArticleUnreadNoticeSelectSql         = "SELECT `id`, `notice_id`, `uid`, `is_new` FROM `ims_article_unread_notice` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsAttachmentGroupInsertSql             = "INSERT INTO `ims_attachment_group` ( `pid`, `name`, `uniacid`, `uid`, `type` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsAttachmentGroupDeleteSql             = "DELETE FROM `ims_attachment_group` WHERE ( `id` = ? );"
	ImsAttachmentGroupUpdateSql             = "UPDATE `ims_attachment_group` SET `pid` = ?, `name` = ?, `uniacid` = ?, `uid` = ?, `type` = ? WHERE ( `id` = ? );"
	ImsAttachmentGroupSelectSql             = "SELECT `id`, `pid`, `name`, `uniacid`, `uid`, `type` FROM `ims_attachment_group` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsBasicReplyInsertSql                  = "INSERT INTO `ims_basic_reply` ( `rid`, `content` ) VALUES ( ?, ? );"
	ImsBasicReplyDeleteSql                  = "DELETE FROM `ims_basic_reply` WHERE ( `id` = ? );"
	ImsBasicReplyUpdateSql                  = "UPDATE `ims_basic_reply` SET `rid` = ?, `content` = ? WHERE ( `id` = ? );"
	ImsBasicReplySelectSql                  = "SELECT `id`, `rid`, `content` FROM `ims_basic_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsBusinessInsertSql                    = "INSERT INTO `ims_business` ( `weid`, `title`, `thumb`, `content`, `phone`, `qq`, `province`, `city`, `dist`, `address`, `lng`, `lat`, `industry1`, `industry2`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsBusinessDeleteSql                    = "DELETE FROM `ims_business` WHERE ( `id` = ? );"
	ImsBusinessUpdateSql                    = "UPDATE `ims_business` SET `weid` = ?, `title` = ?, `thumb` = ?, `content` = ?, `phone` = ?, `qq` = ?, `province` = ?, `city` = ?, `dist` = ?, `address` = ?, `lng` = ?, `lat` = ?, `industry1` = ?, `industry2` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsBusinessSelectSql                    = "SELECT `id`, `weid`, `title`, `thumb`, `content`, `phone`, `qq`, `province`, `city`, `dist`, `address`, `lng`, `lat`, `industry1`, `industry2`, `createtime` FROM `ims_business` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoreAttachmentInsertSql              = "INSERT INTO `ims_core_attachment` ( `uniacid`, `uid`, `filename`, `attachment`, `type`, `createtime`, `module_upload_dir`, `group_id`, `displayorder` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsCoreAttachmentDeleteSql              = "DELETE FROM `ims_core_attachment` WHERE ( `id` = ? );"
	ImsCoreAttachmentUpdateSql              = "UPDATE `ims_core_attachment` SET `uniacid` = ?, `uid` = ?, `filename` = ?, `attachment` = ?, `type` = ?, `createtime` = ?, `module_upload_dir` = ?, `group_id` = ?, `displayorder` = ? WHERE ( `id` = ? );"
	ImsCoreAttachmentSelectSql              = "SELECT `id`, `uniacid`, `uid`, `filename`, `attachment`, `type`, `createtime`, `module_upload_dir`, `group_id`, `displayorder` FROM `ims_core_attachment` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoreCacheInsertSql                   = "INSERT INTO `ims_core_cache` ( `value` ) VALUES ( ? );"
	ImsCoreCacheDeleteSql                   = "DELETE FROM `ims_core_cache` WHERE ( `key` = ? );"
	ImsCoreCacheUpdateSql                   = "UPDATE `ims_core_cache` SET `value` = ? WHERE ( `key` = ? );"
	ImsCoreCacheSelectSql                   = "SELECT `key`, `value` FROM `ims_core_cache` WHERE ( `key` = ? ) ORDER BY `key` DESC LIMIT 0, 1;"
	ImsCoreCronInsertSql                    = "INSERT INTO `ims_core_cron` ( `cloudid`, `module`, `uniacid`, `type`, `name`, `filename`, `lastruntime`, `nextruntime`, `weekday`, `day`, `hour`, `minute`, `extra`, `status`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsCoreCronDeleteSql                    = "DELETE FROM `ims_core_cron` WHERE ( `id` = ? );"
	ImsCoreCronUpdateSql                    = "UPDATE `ims_core_cron` SET `cloudid` = ?, `module` = ?, `uniacid` = ?, `type` = ?, `name` = ?, `filename` = ?, `lastruntime` = ?, `nextruntime` = ?, `weekday` = ?, `day` = ?, `hour` = ?, `minute` = ?, `extra` = ?, `status` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsCoreCronSelectSql                    = "SELECT `id`, `cloudid`, `module`, `uniacid`, `type`, `name`, `filename`, `lastruntime`, `nextruntime`, `weekday`, `day`, `hour`, `minute`, `extra`, `status`, `createtime` FROM `ims_core_cron` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoreCronRecordInsertSql              = "INSERT INTO `ims_core_cron_record` ( `uniacid`, `module`, `type`, `tid`, `note`, `tag`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ?, ? );"
	ImsCoreCronRecordDeleteSql              = "DELETE FROM `ims_core_cron_record` WHERE ( `id` = ? );"
	ImsCoreCronRecordUpdateSql              = "UPDATE `ims_core_cron_record` SET `uniacid` = ?, `module` = ?, `type` = ?, `tid` = ?, `note` = ?, `tag` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsCoreCronRecordSelectSql              = "SELECT `id`, `uniacid`, `module`, `type`, `tid`, `note`, `tag`, `createtime` FROM `ims_core_cron_record` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoreJobInsertSql                     = "INSERT INTO `ims_core_job` ( `type`, `uniacid`, `payload`, `status`, `title`, `handled`, `total`, `createtime`, `updatetime`, `endtime`, `uid`, `isdeleted` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsCoreJobDeleteSql                     = "DELETE FROM `ims_core_job` WHERE ( `id` = ? );"
	ImsCoreJobUpdateSql                     = "UPDATE `ims_core_job` SET `type` = ?, `uniacid` = ?, `payload` = ?, `status` = ?, `title` = ?, `handled` = ?, `total` = ?, `createtime` = ?, `updatetime` = ?, `endtime` = ?, `uid` = ?, `isdeleted` = ? WHERE ( `id` = ? );"
	ImsCoreJobSelectSql                     = "SELECT `id`, `type`, `uniacid`, `payload`, `status`, `title`, `handled`, `total`, `createtime`, `updatetime`, `endtime`, `uid`, `isdeleted` FROM `ims_core_job` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoreMenuInsertSql                    = "INSERT INTO `ims_core_menu` ( `pid`, `title`, `name`, `url`, `append_title`, `append_url`, `displayorder`, `type`, `is_display`, `is_system`, `permission_name`, `group_name`, `icon` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsCoreMenuDeleteSql                    = "DELETE FROM `ims_core_menu` WHERE ( `id` = ? );"
	ImsCoreMenuUpdateSql                    = "UPDATE `ims_core_menu` SET `pid` = ?, `title` = ?, `name` = ?, `url` = ?, `append_title` = ?, `append_url` = ?, `displayorder` = ?, `type` = ?, `is_display` = ?, `is_system` = ?, `permission_name` = ?, `group_name` = ?, `icon` = ? WHERE ( `id` = ? );"
	ImsCoreMenuSelectSql                    = "SELECT `id`, `pid`, `title`, `name`, `url`, `append_title`, `append_url`, `displayorder`, `type`, `is_display`, `is_system`, `permission_name`, `group_name`, `icon` FROM `ims_core_menu` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoreMenuShortcutInsertSql            = "INSERT INTO `ims_core_menu_shortcut` ( `uid`, `uniacid`, `modulename`, `displayorder`, `position`, `updatetime` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsCoreMenuShortcutDeleteSql            = "DELETE FROM `ims_core_menu_shortcut` WHERE ( `id` = ? );"
	ImsCoreMenuShortcutUpdateSql            = "UPDATE `ims_core_menu_shortcut` SET `uid` = ?, `uniacid` = ?, `modulename` = ?, `displayorder` = ?, `position` = ?, `updatetime` = ? WHERE ( `id` = ? );"
	ImsCoreMenuShortcutSelectSql            = "SELECT `id`, `uid`, `uniacid`, `modulename`, `displayorder`, `position`, `updatetime` FROM `ims_core_menu_shortcut` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCorePaylogInsertSql                  = "INSERT INTO `ims_core_paylog` ( `type`, `uniacid`, `acid`, `openid`, `uniontid`, `tid`, `fee`, `status`, `module`, `tag`, `is_usecard`, `card_type`, `card_id`, `card_fee`, `encrypt_code`, `is_wish` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsCorePaylogDeleteSql                  = "DELETE FROM `ims_core_paylog` WHERE ( `plid` = ? );"
	ImsCorePaylogUpdateSql                  = "UPDATE `ims_core_paylog` SET `type` = ?, `uniacid` = ?, `acid` = ?, `openid` = ?, `uniontid` = ?, `tid` = ?, `fee` = ?, `status` = ?, `module` = ?, `tag` = ?, `is_usecard` = ?, `card_type` = ?, `card_id` = ?, `card_fee` = ?, `encrypt_code` = ?, `is_wish` = ? WHERE ( `plid` = ? );"
	ImsCorePaylogSelectSql                  = "SELECT `plid`, `type`, `uniacid`, `acid`, `openid`, `uniontid`, `tid`, `fee`, `status`, `module`, `tag`, `is_usecard`, `card_type`, `card_id`, `card_fee`, `encrypt_code`, `is_wish` FROM `ims_core_paylog` WHERE ( `plid` = ? ) ORDER BY `plid` DESC LIMIT 0, 1;"
	ImsCorePerformanceInsertSql             = "INSERT INTO `ims_core_performance` ( `type`, `runtime`, `runurl`, `runsql`, `createtime` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsCorePerformanceDeleteSql             = "DELETE FROM `ims_core_performance` WHERE ( `id` = ? );"
	ImsCorePerformanceUpdateSql             = "UPDATE `ims_core_performance` SET `type` = ?, `runtime` = ?, `runurl` = ?, `runsql` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsCorePerformanceSelectSql             = "SELECT `id`, `type`, `runtime`, `runurl`, `runsql`, `createtime` FROM `ims_core_performance` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoreQueueInsertSql                   = "INSERT INTO `ims_core_queue` ( `uniacid`, `acid`, `message`, `params`, `keyword`, `response`, `module`, `type`, `dateline` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsCoreQueueDeleteSql                   = "DELETE FROM `ims_core_queue` WHERE ( `qid` = ? );"
	ImsCoreQueueUpdateSql                   = "UPDATE `ims_core_queue` SET `uniacid` = ?, `acid` = ?, `message` = ?, `params` = ?, `keyword` = ?, `response` = ?, `module` = ?, `type` = ?, `dateline` = ? WHERE ( `qid` = ? );"
	ImsCoreQueueSelectSql                   = "SELECT `qid`, `uniacid`, `acid`, `message`, `params`, `keyword`, `response`, `module`, `type`, `dateline` FROM `ims_core_queue` WHERE ( `qid` = ? ) ORDER BY `qid` DESC LIMIT 0, 1;"
	ImsCoreRefundlogInsertSql               = "INSERT INTO `ims_core_refundlog` ( `uniacid`, `refund_uniontid`, `reason`, `uniontid`, `fee`, `status`, `is_wish` ) VALUES ( ?, ?, ?, ?, ?, ?, ? );"
	ImsCoreRefundlogDeleteSql               = "DELETE FROM `ims_core_refundlog` WHERE ( `id` = ? );"
	ImsCoreRefundlogUpdateSql               = "UPDATE `ims_core_refundlog` SET `uniacid` = ?, `refund_uniontid` = ?, `reason` = ?, `uniontid` = ?, `fee` = ?, `status` = ?, `is_wish` = ? WHERE ( `id` = ? );"
	ImsCoreRefundlogSelectSql               = "SELECT `id`, `uniacid`, `refund_uniontid`, `reason`, `uniontid`, `fee`, `status`, `is_wish` FROM `ims_core_refundlog` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoreResourceInsertSql                = "INSERT INTO `ims_core_resource` ( `uniacid`, `media_id`, `trunk`, `type`, `dateline` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsCoreResourceDeleteSql                = "DELETE FROM `ims_core_resource` WHERE ( `mid` = ? );"
	ImsCoreResourceUpdateSql                = "UPDATE `ims_core_resource` SET `uniacid` = ?, `media_id` = ?, `trunk` = ?, `type` = ?, `dateline` = ? WHERE ( `mid` = ? );"
	ImsCoreResourceSelectSql                = "SELECT `mid`, `uniacid`, `media_id`, `trunk`, `type`, `dateline` FROM `ims_core_resource` WHERE ( `mid` = ? ) ORDER BY `mid` DESC LIMIT 0, 1;"
	ImsCoreSendsmsLogInsertSql              = "INSERT INTO `ims_core_sendsms_log` ( `uniacid`, `mobile`, `content`, `result`, `createtime` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsCoreSendsmsLogDeleteSql              = "DELETE FROM `ims_core_sendsms_log` WHERE ( `id` = ? );"
	ImsCoreSendsmsLogUpdateSql              = "UPDATE `ims_core_sendsms_log` SET `uniacid` = ?, `mobile` = ?, `content` = ?, `result` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsCoreSendsmsLogSelectSql              = "SELECT `id`, `uniacid`, `mobile`, `content`, `result`, `createtime` FROM `ims_core_sendsms_log` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoreSessionsInsertSql                = "INSERT INTO `ims_core_sessions` ( `uniacid`, `openid`, `data`, `expiretime` ) VALUES ( ?, ?, ?, ? );"
	ImsCoreSessionsDeleteSql                = "DELETE FROM `ims_core_sessions` WHERE ( `sid` = ? );"
	ImsCoreSessionsUpdateSql                = "UPDATE `ims_core_sessions` SET `uniacid` = ?, `openid` = ?, `data` = ?, `expiretime` = ? WHERE ( `sid` = ? );"
	ImsCoreSessionsSelectSql                = "SELECT `sid`, `uniacid`, `openid`, `data`, `expiretime` FROM `ims_core_sessions` WHERE ( `sid` = ? ) ORDER BY `sid` DESC LIMIT 0, 1;"
	ImsCoreSettingsInsertSql                = "INSERT INTO `ims_core_settings` ( `value` ) VALUES ( ? );"
	ImsCoreSettingsDeleteSql                = "DELETE FROM `ims_core_settings` WHERE ( `key` = ? );"
	ImsCoreSettingsUpdateSql                = "UPDATE `ims_core_settings` SET `value` = ? WHERE ( `key` = ? );"
	ImsCoreSettingsSelectSql                = "SELECT `key`, `value` FROM `ims_core_settings` WHERE ( `key` = ? ) ORDER BY `key` DESC LIMIT 0, 1;"
	ImsCouponLocationInsertSql              = "INSERT INTO `ims_coupon_location` ( `uniacid`, `acid`, `sid`, `location_id`, `business_name`, `branch_name`, `category`, `province`, `city`, `district`, `address`, `longitude`, `latitude`, `telephone`, `photo_list`, `avg_price`, `open_time`, `recommend`, `special`, `introduction`, `offset_type`, `status`, `message` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsCouponLocationDeleteSql              = "DELETE FROM `ims_coupon_location` WHERE ( `id` = ? );"
	ImsCouponLocationUpdateSql              = "UPDATE `ims_coupon_location` SET `uniacid` = ?, `acid` = ?, `sid` = ?, `location_id` = ?, `business_name` = ?, `branch_name` = ?, `category` = ?, `province` = ?, `city` = ?, `district` = ?, `address` = ?, `longitude` = ?, `latitude` = ?, `telephone` = ?, `photo_list` = ?, `avg_price` = ?, `open_time` = ?, `recommend` = ?, `special` = ?, `introduction` = ?, `offset_type` = ?, `status` = ?, `message` = ? WHERE ( `id` = ? );"
	ImsCouponLocationSelectSql              = "SELECT `id`, `uniacid`, `acid`, `sid`, `location_id`, `business_name`, `branch_name`, `category`, `province`, `city`, `district`, `address`, `longitude`, `latitude`, `telephone`, `photo_list`, `avg_price`, `open_time`, `recommend`, `special`, `introduction`, `offset_type`, `status`, `message` FROM `ims_coupon_location` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCoverReplyInsertSql                  = "INSERT INTO `ims_cover_reply` ( `uniacid`, `multiid`, `rid`, `module`, `do`, `title`, `description`, `thumb`, `url` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsCoverReplyDeleteSql                  = "DELETE FROM `ims_cover_reply` WHERE ( `id` = ? );"
	ImsCoverReplyUpdateSql                  = "UPDATE `ims_cover_reply` SET `uniacid` = ?, `multiid` = ?, `rid` = ?, `module` = ?, `do` = ?, `title` = ?, `description` = ?, `thumb` = ?, `url` = ? WHERE ( `id` = ? );"
	ImsCoverReplySelectSql                  = "SELECT `id`, `uniacid`, `multiid`, `rid`, `module`, `do`, `title`, `description`, `thumb`, `url` FROM `ims_cover_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsCustomReplyInsertSql                 = "INSERT INTO `ims_custom_reply` ( `rid`, `start1`, `end1`, `start2`, `end2` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsCustomReplyDeleteSql                 = "DELETE FROM `ims_custom_reply` WHERE ( `id` = ? );"
	ImsCustomReplyUpdateSql                 = "UPDATE `ims_custom_reply` SET `rid` = ?, `start1` = ?, `end1` = ?, `start2` = ?, `end2` = ? WHERE ( `id` = ? );"
	ImsCustomReplySelectSql                 = "SELECT `id`, `rid`, `start1`, `end1`, `start2`, `end2` FROM `ims_custom_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsImagesReplyInsertSql                 = "INSERT INTO `ims_images_reply` ( `rid`, `title`, `description`, `mediaid`, `createtime` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsImagesReplyDeleteSql                 = "DELETE FROM `ims_images_reply` WHERE ( `id` = ? );"
	ImsImagesReplyUpdateSql                 = "UPDATE `ims_images_reply` SET `rid` = ?, `title` = ?, `description` = ?, `mediaid` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsImagesReplySelectSql                 = "SELECT `id`, `rid`, `title`, `description`, `mediaid`, `createtime` FROM `ims_images_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcCashRecordInsertSql                = "INSERT INTO `ims_mc_cash_record` ( `uniacid`, `uid`, `clerk_id`, `store_id`, `clerk_type`, `fee`, `final_fee`, `credit1`, `credit1_fee`, `credit2`, `cash`, `return_cash`, `final_cash`, `remark`, `createtime`, `trade_type` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMcCashRecordDeleteSql                = "DELETE FROM `ims_mc_cash_record` WHERE ( `id` = ? );"
	ImsMcCashRecordUpdateSql                = "UPDATE `ims_mc_cash_record` SET `uniacid` = ?, `uid` = ?, `clerk_id` = ?, `store_id` = ?, `clerk_type` = ?, `fee` = ?, `final_fee` = ?, `credit1` = ?, `credit1_fee` = ?, `credit2` = ?, `cash` = ?, `return_cash` = ?, `final_cash` = ?, `remark` = ?, `createtime` = ?, `trade_type` = ? WHERE ( `id` = ? );"
	ImsMcCashRecordSelectSql                = "SELECT `id`, `uniacid`, `uid`, `clerk_id`, `store_id`, `clerk_type`, `fee`, `final_fee`, `credit1`, `credit1_fee`, `credit2`, `cash`, `return_cash`, `final_cash`, `remark`, `createtime`, `trade_type` FROM `ims_mc_cash_record` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcChatsRecordInsertSql               = "INSERT INTO `ims_mc_chats_record` ( `uniacid`, `acid`, `flag`, `openid`, `msgtype`, `content`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ?, ? );"
	ImsMcChatsRecordDeleteSql               = "DELETE FROM `ims_mc_chats_record` WHERE ( `id` = ? );"
	ImsMcChatsRecordUpdateSql               = "UPDATE `ims_mc_chats_record` SET `uniacid` = ?, `acid` = ?, `flag` = ?, `openid` = ?, `msgtype` = ?, `content` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsMcChatsRecordSelectSql               = "SELECT `id`, `uniacid`, `acid`, `flag`, `openid`, `msgtype`, `content`, `createtime` FROM `ims_mc_chats_record` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcCreditsRechargeInsertSql           = "INSERT INTO `ims_mc_credits_recharge` ( `uniacid`, `uid`, `openid`, `tid`, `transid`, `fee`, `type`, `tag`, `status`, `createtime`, `backtype` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMcCreditsRechargeDeleteSql           = "DELETE FROM `ims_mc_credits_recharge` WHERE ( `id` = ? );"
	ImsMcCreditsRechargeUpdateSql           = "UPDATE `ims_mc_credits_recharge` SET `uniacid` = ?, `uid` = ?, `openid` = ?, `tid` = ?, `transid` = ?, `fee` = ?, `type` = ?, `tag` = ?, `status` = ?, `createtime` = ?, `backtype` = ? WHERE ( `id` = ? );"
	ImsMcCreditsRechargeSelectSql           = "SELECT `id`, `uniacid`, `uid`, `openid`, `tid`, `transid`, `fee`, `type`, `tag`, `status`, `createtime`, `backtype` FROM `ims_mc_credits_recharge` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcCreditsRecordInsertSql             = "INSERT INTO `ims_mc_credits_record` ( `uid`, `uniacid`, `credittype`, `num`, `operator`, `module`, `clerk_id`, `store_id`, `clerk_type`, `createtime`, `remark`, `real_uniacid` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMcCreditsRecordDeleteSql             = "DELETE FROM `ims_mc_credits_record` WHERE ( `id` = ? );"
	ImsMcCreditsRecordUpdateSql             = "UPDATE `ims_mc_credits_record` SET `uid` = ?, `uniacid` = ?, `credittype` = ?, `num` = ?, `operator` = ?, `module` = ?, `clerk_id` = ?, `store_id` = ?, `clerk_type` = ?, `createtime` = ?, `remark` = ?, `real_uniacid` = ? WHERE ( `id` = ? );"
	ImsMcCreditsRecordSelectSql             = "SELECT `id`, `uid`, `uniacid`, `credittype`, `num`, `operator`, `module`, `clerk_id`, `store_id`, `clerk_type`, `createtime`, `remark`, `real_uniacid` FROM `ims_mc_credits_record` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcFansGroupsInsertSql                = "INSERT INTO `ims_mc_fans_groups` ( `uniacid`, `acid`, `groups` ) VALUES ( ?, ?, ? );"
	ImsMcFansGroupsDeleteSql                = "DELETE FROM `ims_mc_fans_groups` WHERE ( `id` = ? );"
	ImsMcFansGroupsUpdateSql                = "UPDATE `ims_mc_fans_groups` SET `uniacid` = ?, `acid` = ?, `groups` = ? WHERE ( `id` = ? );"
	ImsMcFansGroupsSelectSql                = "SELECT `id`, `uniacid`, `acid`, `groups` FROM `ims_mc_fans_groups` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcFansTagInsertSql                   = "INSERT INTO `ims_mc_fans_tag` ( `uniacid`, `fanid`, `openid`, `subscribe`, `nickname`, `sex`, `language`, `city`, `province`, `country`, `headimgurl`, `subscribe_time`, `unionid`, `remark`, `groupid`, `tagid_list`, `subscribe_scene`, `qr_scene_str`, `qr_scene` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMcFansTagDeleteSql                   = "DELETE FROM `ims_mc_fans_tag` WHERE ( `id` = ? );"
	ImsMcFansTagUpdateSql                   = "UPDATE `ims_mc_fans_tag` SET `uniacid` = ?, `fanid` = ?, `openid` = ?, `subscribe` = ?, `nickname` = ?, `sex` = ?, `language` = ?, `city` = ?, `province` = ?, `country` = ?, `headimgurl` = ?, `subscribe_time` = ?, `unionid` = ?, `remark` = ?, `groupid` = ?, `tagid_list` = ?, `subscribe_scene` = ?, `qr_scene_str` = ?, `qr_scene` = ? WHERE ( `id` = ? );"
	ImsMcFansTagSelectSql                   = "SELECT `id`, `uniacid`, `fanid`, `openid`, `subscribe`, `nickname`, `sex`, `language`, `city`, `province`, `country`, `headimgurl`, `subscribe_time`, `unionid`, `remark`, `groupid`, `tagid_list`, `subscribe_scene`, `qr_scene_str`, `qr_scene` FROM `ims_mc_fans_tag` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcFansTagMappingInsertSql            = "INSERT INTO `ims_mc_fans_tag_mapping` ( `fanid`, `tagid` ) VALUES ( ?, ? );"
	ImsMcFansTagMappingDeleteSql            = "DELETE FROM `ims_mc_fans_tag_mapping` WHERE ( `id` = ? );"
	ImsMcFansTagMappingUpdateSql            = "UPDATE `ims_mc_fans_tag_mapping` SET `fanid` = ?, `tagid` = ? WHERE ( `id` = ? );"
	ImsMcFansTagMappingSelectSql            = "SELECT `id`, `fanid`, `tagid` FROM `ims_mc_fans_tag_mapping` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcGroupsInsertSql                    = "INSERT INTO `ims_mc_groups` ( `uniacid`, `title`, `credit`, `isdefault` ) VALUES ( ?, ?, ?, ? );"
	ImsMcGroupsDeleteSql                    = "DELETE FROM `ims_mc_groups` WHERE ( `groupid` = ? );"
	ImsMcGroupsUpdateSql                    = "UPDATE `ims_mc_groups` SET `uniacid` = ?, `title` = ?, `credit` = ?, `isdefault` = ? WHERE ( `groupid` = ? );"
	ImsMcGroupsSelectSql                    = "SELECT `groupid`, `uniacid`, `title`, `credit`, `isdefault` FROM `ims_mc_groups` WHERE ( `groupid` = ? ) ORDER BY `groupid` DESC LIMIT 0, 1;"
	ImsMcHandselInsertSql                   = "INSERT INTO `ims_mc_handsel` ( `uniacid`, `touid`, `fromuid`, `module`, `sign`, `action`, `credit_value`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMcHandselDeleteSql                   = "DELETE FROM `ims_mc_handsel` WHERE ( `id` = ? );"
	ImsMcHandselUpdateSql                   = "UPDATE `ims_mc_handsel` SET `uniacid` = ?, `touid` = ?, `fromuid` = ?, `module` = ?, `sign` = ?, `action` = ?, `credit_value` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsMcHandselSelectSql                   = "SELECT `id`, `uniacid`, `touid`, `fromuid`, `module`, `sign`, `action`, `credit_value`, `createtime` FROM `ims_mc_handsel` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcMappingFansInsertSql               = "INSERT INTO `ims_mc_mapping_fans` ( `acid`, `uniacid`, `uid`, `openid`, `nickname`, `groupid`, `salt`, `follow`, `followtime`, `unfollowtime`, `tag`, `updatetime`, `unionid`, `user_from` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMcMappingFansDeleteSql               = "DELETE FROM `ims_mc_mapping_fans` WHERE ( `fanid` = ? );"
	ImsMcMappingFansUpdateSql               = "UPDATE `ims_mc_mapping_fans` SET `acid` = ?, `uniacid` = ?, `uid` = ?, `openid` = ?, `nickname` = ?, `groupid` = ?, `salt` = ?, `follow` = ?, `followtime` = ?, `unfollowtime` = ?, `tag` = ?, `updatetime` = ?, `unionid` = ?, `user_from` = ? WHERE ( `fanid` = ? );"
	ImsMcMappingFansSelectSql               = "SELECT `fanid`, `acid`, `uniacid`, `uid`, `openid`, `nickname`, `groupid`, `salt`, `follow`, `followtime`, `unfollowtime`, `tag`, `updatetime`, `unionid`, `user_from` FROM `ims_mc_mapping_fans` WHERE ( `fanid` = ? ) ORDER BY `fanid` DESC LIMIT 0, 1;"
	ImsMcMassRecordInsertSql                = "INSERT INTO `ims_mc_mass_record` ( `uniacid`, `acid`, `groupname`, `fansnum`, `msgtype`, `content`, `group`, `attach_id`, `media_id`, `type`, `status`, `cron_id`, `sendtime`, `finalsendtime`, `createtime`, `msg_id`, `msg_data_id` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMcMassRecordDeleteSql                = "DELETE FROM `ims_mc_mass_record` WHERE ( `id` = ? );"
	ImsMcMassRecordUpdateSql                = "UPDATE `ims_mc_mass_record` SET `uniacid` = ?, `acid` = ?, `groupname` = ?, `fansnum` = ?, `msgtype` = ?, `content` = ?, `group` = ?, `attach_id` = ?, `media_id` = ?, `type` = ?, `status` = ?, `cron_id` = ?, `sendtime` = ?, `finalsendtime` = ?, `createtime` = ?, `msg_id` = ?, `msg_data_id` = ? WHERE ( `id` = ? );"
	ImsMcMassRecordSelectSql                = "SELECT `id`, `uniacid`, `acid`, `groupname`, `fansnum`, `msgtype`, `content`, `group`, `attach_id`, `media_id`, `type`, `status`, `cron_id`, `sendtime`, `finalsendtime`, `createtime`, `msg_id`, `msg_data_id` FROM `ims_mc_mass_record` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcMembersInsertSql                   = "INSERT INTO `ims_mc_members` ( `uniacid`, `mobile`, `email`, `password`, `salt`, `groupid`, `credit1`, `credit2`, `credit3`, `credit4`, `credit5`, `credit6`, `createtime`, `realname`, `nickname`, `avatar`, `qq`, `vip`, `gender`, `birthyear`, `birthmonth`, `birthday`, `constellation`, `zodiac`, `telephone`, `idcard`, `studentid`, `grade`, `address`, `zipcode`, `nationality`, `resideprovince`, `residecity`, `residedist`, `graduateschool`, `company`, `education`, `occupation`, `position`, `revenue`, `affectivestatus`, `lookingfor`, `bloodtype`, `height`, `weight`, `alipay`, `msn`, `taobao`, `site`, `bio`, `interest`, `pay_password`, `user_from` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMcMembersDeleteSql                   = "DELETE FROM `ims_mc_members` WHERE ( `uid` = ? );"
	ImsMcMembersUpdateSql                   = "UPDATE `ims_mc_members` SET `uniacid` = ?, `mobile` = ?, `email` = ?, `password` = ?, `salt` = ?, `groupid` = ?, `credit1` = ?, `credit2` = ?, `credit3` = ?, `credit4` = ?, `credit5` = ?, `credit6` = ?, `createtime` = ?, `realname` = ?, `nickname` = ?, `avatar` = ?, `qq` = ?, `vip` = ?, `gender` = ?, `birthyear` = ?, `birthmonth` = ?, `birthday` = ?, `constellation` = ?, `zodiac` = ?, `telephone` = ?, `idcard` = ?, `studentid` = ?, `grade` = ?, `address` = ?, `zipcode` = ?, `nationality` = ?, `resideprovince` = ?, `residecity` = ?, `residedist` = ?, `graduateschool` = ?, `company` = ?, `education` = ?, `occupation` = ?, `position` = ?, `revenue` = ?, `affectivestatus` = ?, `lookingfor` = ?, `bloodtype` = ?, `height` = ?, `weight` = ?, `alipay` = ?, `msn` = ?, `taobao` = ?, `site` = ?, `bio` = ?, `interest` = ?, `pay_password` = ?, `user_from` = ? WHERE ( `uid` = ? );"
	ImsMcMembersSelectSql                   = "SELECT `uid`, `uniacid`, `mobile`, `email`, `password`, `salt`, `groupid`, `credit1`, `credit2`, `credit3`, `credit4`, `credit5`, `credit6`, `createtime`, `realname`, `nickname`, `avatar`, `qq`, `vip`, `gender`, `birthyear`, `birthmonth`, `birthday`, `constellation`, `zodiac`, `telephone`, `idcard`, `studentid`, `grade`, `address`, `zipcode`, `nationality`, `resideprovince`, `residecity`, `residedist`, `graduateschool`, `company`, `education`, `occupation`, `position`, `revenue`, `affectivestatus`, `lookingfor`, `bloodtype`, `height`, `weight`, `alipay`, `msn`, `taobao`, `site`, `bio`, `interest`, `pay_password`, `user_from` FROM `ims_mc_members` WHERE ( `uid` = ? ) ORDER BY `uid` DESC LIMIT 0, 1;"
	ImsMcMemberAddressInsertSql             = "INSERT INTO `ims_mc_member_address` ( `uniacid`, `uid`, `username`, `mobile`, `zipcode`, `province`, `city`, `district`, `address`, `isdefault` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMcMemberAddressDeleteSql             = "DELETE FROM `ims_mc_member_address` WHERE ( `id` = ? );"
	ImsMcMemberAddressUpdateSql             = "UPDATE `ims_mc_member_address` SET `uniacid` = ?, `uid` = ?, `username` = ?, `mobile` = ?, `zipcode` = ?, `province` = ?, `city` = ?, `district` = ?, `address` = ?, `isdefault` = ? WHERE ( `id` = ? );"
	ImsMcMemberAddressSelectSql             = "SELECT `id`, `uniacid`, `uid`, `username`, `mobile`, `zipcode`, `province`, `city`, `district`, `address`, `isdefault` FROM `ims_mc_member_address` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcMemberFieldsInsertSql              = "INSERT INTO `ims_mc_member_fields` ( `uniacid`, `fieldid`, `title`, `available`, `displayorder` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsMcMemberFieldsDeleteSql              = "DELETE FROM `ims_mc_member_fields` WHERE ( `id` = ? );"
	ImsMcMemberFieldsUpdateSql              = "UPDATE `ims_mc_member_fields` SET `uniacid` = ?, `fieldid` = ?, `title` = ?, `available` = ?, `displayorder` = ? WHERE ( `id` = ? );"
	ImsMcMemberFieldsSelectSql              = "SELECT `id`, `uniacid`, `fieldid`, `title`, `available`, `displayorder` FROM `ims_mc_member_fields` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcMemberPropertyInsertSql            = "INSERT INTO `ims_mc_member_property` ( `uniacid`, `property` ) VALUES ( ?, ? );"
	ImsMcMemberPropertyDeleteSql            = "DELETE FROM `ims_mc_member_property` WHERE ( `id` = ? );"
	ImsMcMemberPropertyUpdateSql            = "UPDATE `ims_mc_member_property` SET `uniacid` = ?, `property` = ? WHERE ( `id` = ? );"
	ImsMcMemberPropertySelectSql            = "SELECT `id`, `uniacid`, `property` FROM `ims_mc_member_property` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMcOauthFansInsertSql                 = "INSERT INTO `ims_mc_oauth_fans` ( `oauth_openid`, `acid`, `uid`, `openid` ) VALUES ( ?, ?, ?, ? );"
	ImsMcOauthFansDeleteSql                 = "DELETE FROM `ims_mc_oauth_fans` WHERE ( `id` = ? );"
	ImsMcOauthFansUpdateSql                 = "UPDATE `ims_mc_oauth_fans` SET `oauth_openid` = ?, `acid` = ?, `uid` = ?, `openid` = ? WHERE ( `id` = ? );"
	ImsMcOauthFansSelectSql                 = "SELECT `id`, `oauth_openid`, `acid`, `uid`, `openid` FROM `ims_mc_oauth_fans` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMenuEventInsertSql                   = "INSERT INTO `ims_menu_event` ( `uniacid`, `keyword`, `type`, `picmd5`, `openid`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsMenuEventDeleteSql                   = "DELETE FROM `ims_menu_event` WHERE ( `id` = ? );"
	ImsMenuEventUpdateSql                   = "UPDATE `ims_menu_event` SET `uniacid` = ?, `keyword` = ?, `type` = ?, `picmd5` = ?, `openid` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsMenuEventSelectSql                   = "SELECT `id`, `uniacid`, `keyword`, `type`, `picmd5`, `openid`, `createtime` FROM `ims_menu_event` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMessageNoticeLogInsertSql            = "INSERT INTO `ims_message_notice_log` ( `message`, `is_read`, `uid`, `sign`, `type`, `status`, `create_time`, `end_time`, `url` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsMessageNoticeLogDeleteSql            = "DELETE FROM `ims_message_notice_log` WHERE ( `id` = ? );"
	ImsMessageNoticeLogUpdateSql            = "UPDATE `ims_message_notice_log` SET `message` = ?, `is_read` = ?, `uid` = ?, `sign` = ?, `type` = ?, `status` = ?, `create_time` = ?, `end_time` = ?, `url` = ? WHERE ( `id` = ? );"
	ImsMessageNoticeLogSelectSql            = "SELECT `id`, `message`, `is_read`, `uid`, `sign`, `type`, `status`, `create_time`, `end_time`, `url` FROM `ims_message_notice_log` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMobilenumberInsertSql                = "INSERT INTO `ims_mobilenumber` ( `rid`, `enabled`, `dateline` ) VALUES ( ?, ?, ? );"
	ImsMobilenumberDeleteSql                = "DELETE FROM `ims_mobilenumber` WHERE ( `id` = ? );"
	ImsMobilenumberUpdateSql                = "UPDATE `ims_mobilenumber` SET `rid` = ?, `enabled` = ?, `dateline` = ? WHERE ( `id` = ? );"
	ImsMobilenumberSelectSql                = "SELECT `id`, `rid`, `enabled`, `dateline` FROM `ims_mobilenumber` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsModulesInsertSql                     = "INSERT INTO `ims_modules` ( `name`, `application_type`, `type`, `title`, `version`, `ability`, `description`, `author`, `url`, `settings`, `subscribes`, `handles`, `isrulefields`, `issystem`, `target`, `iscard`, `permissions`, `title_initial`, `wxapp_support`, `welcome_support`, `oauth_type`, `webapp_support`, `phoneapp_support`, `account_support`, `xzapp_support`, `aliapp_support`, `logo`, `baiduapp_support`, `toutiaoapp_support`, `from`, `cloud_record`, `sections`, `label` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsModulesDeleteSql                     = "DELETE FROM `ims_modules` WHERE ( `mid` = ? );"
	ImsModulesUpdateSql                     = "UPDATE `ims_modules` SET `name` = ?, `application_type` = ?, `type` = ?, `title` = ?, `version` = ?, `ability` = ?, `description` = ?, `author` = ?, `url` = ?, `settings` = ?, `subscribes` = ?, `handles` = ?, `isrulefields` = ?, `issystem` = ?, `target` = ?, `iscard` = ?, `permissions` = ?, `title_initial` = ?, `wxapp_support` = ?, `welcome_support` = ?, `oauth_type` = ?, `webapp_support` = ?, `phoneapp_support` = ?, `account_support` = ?, `xzapp_support` = ?, `aliapp_support` = ?, `logo` = ?, `baiduapp_support` = ?, `toutiaoapp_support` = ?, `from` = ?, `cloud_record` = ?, `sections` = ?, `label` = ? WHERE ( `mid` = ? );"
	ImsModulesSelectSql                     = "SELECT `mid`, `name`, `application_type`, `type`, `title`, `version`, `ability`, `description`, `author`, `url`, `settings`, `subscribes`, `handles`, `isrulefields`, `issystem`, `target`, `iscard`, `permissions`, `title_initial`, `wxapp_support`, `welcome_support`, `oauth_type`, `webapp_support`, `phoneapp_support`, `account_support`, `xzapp_support`, `aliapp_support`, `logo`, `baiduapp_support`, `toutiaoapp_support`, `from`, `cloud_record`, `sections`, `label` FROM `ims_modules` WHERE ( `mid` = ? ) ORDER BY `mid` DESC LIMIT 0, 1;"
	ImsModulesBindingsInsertSql             = "INSERT INTO `ims_modules_bindings` ( `module`, `entry`, `call`, `title`, `do`, `state`, `direct`, `url`, `icon`, `displayorder`, `multilevel`, `parent` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsModulesBindingsDeleteSql             = "DELETE FROM `ims_modules_bindings` WHERE ( `eid` = ? );"
	ImsModulesBindingsUpdateSql             = "UPDATE `ims_modules_bindings` SET `module` = ?, `entry` = ?, `call` = ?, `title` = ?, `do` = ?, `state` = ?, `direct` = ?, `url` = ?, `icon` = ?, `displayorder` = ?, `multilevel` = ?, `parent` = ? WHERE ( `eid` = ? );"
	ImsModulesBindingsSelectSql             = "SELECT `eid`, `module`, `entry`, `call`, `title`, `do`, `state`, `direct`, `url`, `icon`, `displayorder`, `multilevel`, `parent` FROM `ims_modules_bindings` WHERE ( `eid` = ? ) ORDER BY `eid` DESC LIMIT 0, 1;"
	ImsModulesCloudInsertSql                = "INSERT INTO `ims_modules_cloud` ( `name`, `application_type`, `title`, `title_initial`, `logo`, `version`, `install_status`, `account_support`, `wxapp_support`, `webapp_support`, `phoneapp_support`, `welcome_support`, `main_module_name`, `main_module_logo`, `has_new_version`, `has_new_branch`, `is_ban`, `lastupdatetime`, `xzapp_support`, `cloud_id`, `aliapp_support`, `baiduapp_support`, `toutiaoapp_support`, `buytime`, `module_status`, `label` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsModulesCloudDeleteSql                = "DELETE FROM `ims_modules_cloud` WHERE ( `id` = ? );"
	ImsModulesCloudUpdateSql                = "UPDATE `ims_modules_cloud` SET `name` = ?, `application_type` = ?, `title` = ?, `title_initial` = ?, `logo` = ?, `version` = ?, `install_status` = ?, `account_support` = ?, `wxapp_support` = ?, `webapp_support` = ?, `phoneapp_support` = ?, `welcome_support` = ?, `main_module_name` = ?, `main_module_logo` = ?, `has_new_version` = ?, `has_new_branch` = ?, `is_ban` = ?, `lastupdatetime` = ?, `xzapp_support` = ?, `cloud_id` = ?, `aliapp_support` = ?, `baiduapp_support` = ?, `toutiaoapp_support` = ?, `buytime` = ?, `module_status` = ?, `label` = ? WHERE ( `id` = ? );"
	ImsModulesCloudSelectSql                = "SELECT `id`, `name`, `application_type`, `title`, `title_initial`, `logo`, `version`, `install_status`, `account_support`, `wxapp_support`, `webapp_support`, `phoneapp_support`, `welcome_support`, `main_module_name`, `main_module_logo`, `has_new_version`, `has_new_branch`, `is_ban`, `lastupdatetime`, `xzapp_support`, `cloud_id`, `aliapp_support`, `baiduapp_support`, `toutiaoapp_support`, `buytime`, `module_status`, `label` FROM `ims_modules_cloud` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsModulesIgnoreInsertSql               = "INSERT INTO `ims_modules_ignore` ( `name`, `version` ) VALUES ( ?, ? );"
	ImsModulesIgnoreDeleteSql               = "DELETE FROM `ims_modules_ignore` WHERE ( `id` = ? );"
	ImsModulesIgnoreUpdateSql               = "UPDATE `ims_modules_ignore` SET `name` = ?, `version` = ? WHERE ( `id` = ? );"
	ImsModulesIgnoreSelectSql               = "SELECT `id`, `name`, `version` FROM `ims_modules_ignore` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsModulesPluginInsertSql               = "INSERT INTO `ims_modules_plugin` ( `name`, `main_module` ) VALUES ( ?, ? );"
	ImsModulesPluginDeleteSql               = "DELETE FROM `ims_modules_plugin` WHERE ( `id` = ? );"
	ImsModulesPluginUpdateSql               = "UPDATE `ims_modules_plugin` SET `name` = ?, `main_module` = ? WHERE ( `id` = ? );"
	ImsModulesPluginSelectSql               = "SELECT `id`, `name`, `main_module` FROM `ims_modules_plugin` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsModulesPluginRankInsertSql           = "INSERT INTO `ims_modules_plugin_rank` ( `uniacid`, `uid`, `rank`, `plugin_name`, `main_module_name` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsModulesPluginRankDeleteSql           = "DELETE FROM `ims_modules_plugin_rank` WHERE ( `id` = ? );"
	ImsModulesPluginRankUpdateSql           = "UPDATE `ims_modules_plugin_rank` SET `uniacid` = ?, `uid` = ?, `rank` = ?, `plugin_name` = ?, `main_module_name` = ? WHERE ( `id` = ? );"
	ImsModulesPluginRankSelectSql           = "SELECT `id`, `uniacid`, `uid`, `rank`, `plugin_name`, `main_module_name` FROM `ims_modules_plugin_rank` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsModulesRankInsertSql                 = "INSERT INTO `ims_modules_rank` ( `module_name`, `uid`, `rank`, `uniacid` ) VALUES ( ?, ?, ?, ? );"
	ImsModulesRankDeleteSql                 = "DELETE FROM `ims_modules_rank` WHERE ( `id` = ? );"
	ImsModulesRankUpdateSql                 = "UPDATE `ims_modules_rank` SET `module_name` = ?, `uid` = ?, `rank` = ?, `uniacid` = ? WHERE ( `id` = ? );"
	ImsModulesRankSelectSql                 = "SELECT `id`, `module_name`, `uid`, `rank`, `uniacid` FROM `ims_modules_rank` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsModulesRecycleInsertSql              = "INSERT INTO `ims_modules_recycle` ( `name`, `type`, `account_support`, `wxapp_support`, `welcome_support`, `webapp_support`, `phoneapp_support`, `xzapp_support`, `aliapp_support`, `baiduapp_support`, `toutiaoapp_support` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsModulesRecycleDeleteSql              = "DELETE FROM `ims_modules_recycle` WHERE ( `id` = ? );"
	ImsModulesRecycleUpdateSql              = "UPDATE `ims_modules_recycle` SET `name` = ?, `type` = ?, `account_support` = ?, `wxapp_support` = ?, `welcome_support` = ?, `webapp_support` = ?, `phoneapp_support` = ?, `xzapp_support` = ?, `aliapp_support` = ?, `baiduapp_support` = ?, `toutiaoapp_support` = ? WHERE ( `id` = ? );"
	ImsModulesRecycleSelectSql              = "SELECT `id`, `name`, `type`, `account_support`, `wxapp_support`, `welcome_support`, `webapp_support`, `phoneapp_support`, `xzapp_support`, `aliapp_support`, `baiduapp_support`, `toutiaoapp_support` FROM `ims_modules_recycle` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsMusicReplyInsertSql                  = "INSERT INTO `ims_music_reply` ( `rid`, `title`, `description`, `url`, `hqurl` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsMusicReplyDeleteSql                  = "DELETE FROM `ims_music_reply` WHERE ( `id` = ? );"
	ImsMusicReplyUpdateSql                  = "UPDATE `ims_music_reply` SET `rid` = ?, `title` = ?, `description` = ?, `url` = ?, `hqurl` = ? WHERE ( `id` = ? );"
	ImsMusicReplySelectSql                  = "SELECT `id`, `rid`, `title`, `description`, `url`, `hqurl` FROM `ims_music_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsNewsReplyInsertSql                   = "INSERT INTO `ims_news_reply` ( `rid`, `parent_id`, `title`, `author`, `description`, `thumb`, `content`, `url`, `displayorder`, `incontent`, `createtime`, `media_id` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsNewsReplyDeleteSql                   = "DELETE FROM `ims_news_reply` WHERE ( `id` = ? );"
	ImsNewsReplyUpdateSql                   = "UPDATE `ims_news_reply` SET `rid` = ?, `parent_id` = ?, `title` = ?, `author` = ?, `description` = ?, `thumb` = ?, `content` = ?, `url` = ?, `displayorder` = ?, `incontent` = ?, `createtime` = ?, `media_id` = ? WHERE ( `id` = ? );"
	ImsNewsReplySelectSql                   = "SELECT `id`, `rid`, `parent_id`, `title`, `author`, `description`, `thumb`, `content`, `url`, `displayorder`, `incontent`, `createtime`, `media_id` FROM `ims_news_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsPhoneappVersionsInsertSql            = "INSERT INTO `ims_phoneapp_versions` ( `uniacid`, `version`, `description`, `modules`, `createtime` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsPhoneappVersionsDeleteSql            = "DELETE FROM `ims_phoneapp_versions` WHERE ( `id` = ? );"
	ImsPhoneappVersionsUpdateSql            = "UPDATE `ims_phoneapp_versions` SET `uniacid` = ?, `version` = ?, `description` = ?, `modules` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsPhoneappVersionsSelectSql            = "SELECT `id`, `uniacid`, `version`, `description`, `modules`, `createtime` FROM `ims_phoneapp_versions` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsProfileFieldsInsertSql               = "INSERT INTO `ims_profile_fields` ( `field`, `available`, `title`, `description`, `displayorder`, `required`, `unchangeable`, `showinregister`, `field_length` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsProfileFieldsDeleteSql               = "DELETE FROM `ims_profile_fields` WHERE ( `id` = ? );"
	ImsProfileFieldsUpdateSql               = "UPDATE `ims_profile_fields` SET `field` = ?, `available` = ?, `title` = ?, `description` = ?, `displayorder` = ?, `required` = ?, `unchangeable` = ?, `showinregister` = ?, `field_length` = ? WHERE ( `id` = ? );"
	ImsProfileFieldsSelectSql               = "SELECT `id`, `field`, `available`, `title`, `description`, `displayorder`, `required`, `unchangeable`, `showinregister`, `field_length` FROM `ims_profile_fields` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsQrcodeInsertSql                      = "INSERT INTO `ims_qrcode` ( `uniacid`, `acid`, `type`, `extra`, `qrcid`, `scene_str`, `name`, `keyword`, `model`, `ticket`, `url`, `expire`, `subnum`, `createtime`, `status` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsQrcodeDeleteSql                      = "DELETE FROM `ims_qrcode` WHERE ( `id` = ? );"
	ImsQrcodeUpdateSql                      = "UPDATE `ims_qrcode` SET `uniacid` = ?, `acid` = ?, `type` = ?, `extra` = ?, `qrcid` = ?, `scene_str` = ?, `name` = ?, `keyword` = ?, `model` = ?, `ticket` = ?, `url` = ?, `expire` = ?, `subnum` = ?, `createtime` = ?, `status` = ? WHERE ( `id` = ? );"
	ImsQrcodeSelectSql                      = "SELECT `id`, `uniacid`, `acid`, `type`, `extra`, `qrcid`, `scene_str`, `name`, `keyword`, `model`, `ticket`, `url`, `expire`, `subnum`, `createtime`, `status` FROM `ims_qrcode` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsQrcodeStatInsertSql                  = "INSERT INTO `ims_qrcode_stat` ( `uniacid`, `acid`, `qid`, `openid`, `type`, `qrcid`, `scene_str`, `name`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsQrcodeStatDeleteSql                  = "DELETE FROM `ims_qrcode_stat` WHERE ( `id` = ? );"
	ImsQrcodeStatUpdateSql                  = "UPDATE `ims_qrcode_stat` SET `uniacid` = ?, `acid` = ?, `qid` = ?, `openid` = ?, `type` = ?, `qrcid` = ?, `scene_str` = ?, `name` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsQrcodeStatSelectSql                  = "SELECT `id`, `uniacid`, `acid`, `qid`, `openid`, `type`, `qrcid`, `scene_str`, `name`, `createtime` FROM `ims_qrcode_stat` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsRuleInsertSql                        = "INSERT INTO `ims_rule` ( `uniacid`, `name`, `module`, `displayorder`, `status`, `containtype` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsRuleDeleteSql                        = "DELETE FROM `ims_rule` WHERE ( `id` = ? );"
	ImsRuleUpdateSql                        = "UPDATE `ims_rule` SET `uniacid` = ?, `name` = ?, `module` = ?, `displayorder` = ?, `status` = ?, `containtype` = ? WHERE ( `id` = ? );"
	ImsRuleSelectSql                        = "SELECT `id`, `uniacid`, `name`, `module`, `displayorder`, `status`, `containtype` FROM `ims_rule` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsRuleKeywordInsertSql                 = "INSERT INTO `ims_rule_keyword` ( `rid`, `uniacid`, `module`, `content`, `type`, `displayorder`, `status` ) VALUES ( ?, ?, ?, ?, ?, ?, ? );"
	ImsRuleKeywordDeleteSql                 = "DELETE FROM `ims_rule_keyword` WHERE ( `id` = ? );"
	ImsRuleKeywordUpdateSql                 = "UPDATE `ims_rule_keyword` SET `rid` = ?, `uniacid` = ?, `module` = ?, `content` = ?, `type` = ?, `displayorder` = ?, `status` = ? WHERE ( `id` = ? );"
	ImsRuleKeywordSelectSql                 = "SELECT `id`, `rid`, `uniacid`, `module`, `content`, `type`, `displayorder`, `status` FROM `ims_rule_keyword` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteArticleInsertSql                 = "INSERT INTO `ims_site_article` ( `uniacid`, `rid`, `kid`, `iscommend`, `ishot`, `pcate`, `ccate`, `template`, `title`, `description`, `content`, `thumb`, `incontent`, `source`, `author`, `displayorder`, `linkurl`, `createtime`, `edittime`, `click`, `type`, `credit` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSiteArticleDeleteSql                 = "DELETE FROM `ims_site_article` WHERE ( `id` = ? );"
	ImsSiteArticleUpdateSql                 = "UPDATE `ims_site_article` SET `uniacid` = ?, `rid` = ?, `kid` = ?, `iscommend` = ?, `ishot` = ?, `pcate` = ?, `ccate` = ?, `template` = ?, `title` = ?, `description` = ?, `content` = ?, `thumb` = ?, `incontent` = ?, `source` = ?, `author` = ?, `displayorder` = ?, `linkurl` = ?, `createtime` = ?, `edittime` = ?, `click` = ?, `type` = ?, `credit` = ? WHERE ( `id` = ? );"
	ImsSiteArticleSelectSql                 = "SELECT `id`, `uniacid`, `rid`, `kid`, `iscommend`, `ishot`, `pcate`, `ccate`, `template`, `title`, `description`, `content`, `thumb`, `incontent`, `source`, `author`, `displayorder`, `linkurl`, `createtime`, `edittime`, `click`, `type`, `credit` FROM `ims_site_article` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteArticleCommentInsertSql          = "INSERT INTO `ims_site_article_comment` ( `uniacid`, `articleid`, `parentid`, `uid`, `openid`, `content`, `is_read`, `iscomment`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSiteArticleCommentDeleteSql          = "DELETE FROM `ims_site_article_comment` WHERE ( `id` = ? );"
	ImsSiteArticleCommentUpdateSql          = "UPDATE `ims_site_article_comment` SET `uniacid` = ?, `articleid` = ?, `parentid` = ?, `uid` = ?, `openid` = ?, `content` = ?, `is_read` = ?, `iscomment` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsSiteArticleCommentSelectSql          = "SELECT `id`, `uniacid`, `articleid`, `parentid`, `uid`, `openid`, `content`, `is_read`, `iscomment`, `createtime` FROM `ims_site_article_comment` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteCategoryInsertSql                = "INSERT INTO `ims_site_category` ( `uniacid`, `nid`, `name`, `parentid`, `displayorder`, `enabled`, `icon`, `description`, `styleid`, `linkurl`, `ishomepage`, `icontype`, `css`, `multiid` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSiteCategoryDeleteSql                = "DELETE FROM `ims_site_category` WHERE ( `id` = ? );"
	ImsSiteCategoryUpdateSql                = "UPDATE `ims_site_category` SET `uniacid` = ?, `nid` = ?, `name` = ?, `parentid` = ?, `displayorder` = ?, `enabled` = ?, `icon` = ?, `description` = ?, `styleid` = ?, `linkurl` = ?, `ishomepage` = ?, `icontype` = ?, `css` = ?, `multiid` = ? WHERE ( `id` = ? );"
	ImsSiteCategorySelectSql                = "SELECT `id`, `uniacid`, `nid`, `name`, `parentid`, `displayorder`, `enabled`, `icon`, `description`, `styleid`, `linkurl`, `ishomepage`, `icontype`, `css`, `multiid` FROM `ims_site_category` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteMultiInsertSql                   = "INSERT INTO `ims_site_multi` ( `uniacid`, `title`, `styleid`, `site_info`, `status`, `bindhost` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsSiteMultiDeleteSql                   = "DELETE FROM `ims_site_multi` WHERE ( `id` = ? );"
	ImsSiteMultiUpdateSql                   = "UPDATE `ims_site_multi` SET `uniacid` = ?, `title` = ?, `styleid` = ?, `site_info` = ?, `status` = ?, `bindhost` = ? WHERE ( `id` = ? );"
	ImsSiteMultiSelectSql                   = "SELECT `id`, `uniacid`, `title`, `styleid`, `site_info`, `status`, `bindhost` FROM `ims_site_multi` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteNavInsertSql                     = "INSERT INTO `ims_site_nav` ( `uniacid`, `multiid`, `section`, `module`, `displayorder`, `name`, `description`, `position`, `url`, `icon`, `css`, `status`, `categoryid` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSiteNavDeleteSql                     = "DELETE FROM `ims_site_nav` WHERE ( `id` = ? );"
	ImsSiteNavUpdateSql                     = "UPDATE `ims_site_nav` SET `uniacid` = ?, `multiid` = ?, `section` = ?, `module` = ?, `displayorder` = ?, `name` = ?, `description` = ?, `position` = ?, `url` = ?, `icon` = ?, `css` = ?, `status` = ?, `categoryid` = ? WHERE ( `id` = ? );"
	ImsSiteNavSelectSql                     = "SELECT `id`, `uniacid`, `multiid`, `section`, `module`, `displayorder`, `name`, `description`, `position`, `url`, `icon`, `css`, `status`, `categoryid` FROM `ims_site_nav` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSitePageInsertSql                    = "INSERT INTO `ims_site_page` ( `uniacid`, `multiid`, `title`, `description`, `params`, `html`, `multipage`, `type`, `status`, `createtime`, `goodnum` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSitePageDeleteSql                    = "DELETE FROM `ims_site_page` WHERE ( `id` = ? );"
	ImsSitePageUpdateSql                    = "UPDATE `ims_site_page` SET `uniacid` = ?, `multiid` = ?, `title` = ?, `description` = ?, `params` = ?, `html` = ?, `multipage` = ?, `type` = ?, `status` = ?, `createtime` = ?, `goodnum` = ? WHERE ( `id` = ? );"
	ImsSitePageSelectSql                    = "SELECT `id`, `uniacid`, `multiid`, `title`, `description`, `params`, `html`, `multipage`, `type`, `status`, `createtime`, `goodnum` FROM `ims_site_page` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteSlideInsertSql                   = "INSERT INTO `ims_site_slide` ( `uniacid`, `multiid`, `title`, `url`, `thumb`, `displayorder` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsSiteSlideDeleteSql                   = "DELETE FROM `ims_site_slide` WHERE ( `id` = ? );"
	ImsSiteSlideUpdateSql                   = "UPDATE `ims_site_slide` SET `uniacid` = ?, `multiid` = ?, `title` = ?, `url` = ?, `thumb` = ?, `displayorder` = ? WHERE ( `id` = ? );"
	ImsSiteSlideSelectSql                   = "SELECT `id`, `uniacid`, `multiid`, `title`, `url`, `thumb`, `displayorder` FROM `ims_site_slide` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteStoreCashLogInsertSql            = "INSERT INTO `ims_site_store_cash_log` ( `founder_uid`, `number`, `amount`, `status`, `create_time` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsSiteStoreCashLogDeleteSql            = "DELETE FROM `ims_site_store_cash_log` WHERE ( `id` = ? );"
	ImsSiteStoreCashLogUpdateSql            = "UPDATE `ims_site_store_cash_log` SET `founder_uid` = ?, `number` = ?, `amount` = ?, `status` = ?, `create_time` = ? WHERE ( `id` = ? );"
	ImsSiteStoreCashLogSelectSql            = "SELECT `id`, `founder_uid`, `number`, `amount`, `status`, `create_time` FROM `ims_site_store_cash_log` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteStoreCashOrderInsertSql          = "INSERT INTO `ims_site_store_cash_order` ( `number`, `founder_uid`, `order_id`, `goods_id`, `order_amount`, `cash_log_id`, `status`, `create_time` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSiteStoreCashOrderDeleteSql          = "DELETE FROM `ims_site_store_cash_order` WHERE ( `id` = ? );"
	ImsSiteStoreCashOrderUpdateSql          = "UPDATE `ims_site_store_cash_order` SET `number` = ?, `founder_uid` = ?, `order_id` = ?, `goods_id` = ?, `order_amount` = ?, `cash_log_id` = ?, `status` = ?, `create_time` = ? WHERE ( `id` = ? );"
	ImsSiteStoreCashOrderSelectSql          = "SELECT `id`, `number`, `founder_uid`, `order_id`, `goods_id`, `order_amount`, `cash_log_id`, `status`, `create_time` FROM `ims_site_store_cash_order` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteStoreCreateAccountInsertSql      = "INSERT INTO `ims_site_store_create_account` ( `uid`, `uniacid`, `type`, `endtime` ) VALUES ( ?, ?, ?, ? );"
	ImsSiteStoreCreateAccountDeleteSql      = "DELETE FROM `ims_site_store_create_account` WHERE ( `id` = ? );"
	ImsSiteStoreCreateAccountUpdateSql      = "UPDATE `ims_site_store_create_account` SET `uid` = ?, `uniacid` = ?, `type` = ?, `endtime` = ? WHERE ( `id` = ? );"
	ImsSiteStoreCreateAccountSelectSql      = "SELECT `id`, `uid`, `uniacid`, `type`, `endtime` FROM `ims_site_store_create_account` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteStoreGoodsInsertSql              = "INSERT INTO `ims_site_store_goods` ( `type`, `title`, `module`, `account_num`, `wxapp_num`, `price`, `unit`, `slide`, `category_id`, `title_initial`, `status`, `createtime`, `synopsis`, `description`, `module_group`, `api_num`, `user_group_price`, `user_group`, `account_group`, `is_wish`, `logo`, `platform_num`, `aliapp_num`, `baiduapp_num`, `phoneapp_num`, `toutiaoapp_num`, `webapp_num`, `xzapp_num` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSiteStoreGoodsDeleteSql              = "DELETE FROM `ims_site_store_goods` WHERE ( `id` = ? );"
	ImsSiteStoreGoodsUpdateSql              = "UPDATE `ims_site_store_goods` SET `type` = ?, `title` = ?, `module` = ?, `account_num` = ?, `wxapp_num` = ?, `price` = ?, `unit` = ?, `slide` = ?, `category_id` = ?, `title_initial` = ?, `status` = ?, `createtime` = ?, `synopsis` = ?, `description` = ?, `module_group` = ?, `api_num` = ?, `user_group_price` = ?, `user_group` = ?, `account_group` = ?, `is_wish` = ?, `logo` = ?, `platform_num` = ?, `aliapp_num` = ?, `baiduapp_num` = ?, `phoneapp_num` = ?, `toutiaoapp_num` = ?, `webapp_num` = ?, `xzapp_num` = ? WHERE ( `id` = ? );"
	ImsSiteStoreGoodsSelectSql              = "SELECT `id`, `type`, `title`, `module`, `account_num`, `wxapp_num`, `price`, `unit`, `slide`, `category_id`, `title_initial`, `status`, `createtime`, `synopsis`, `description`, `module_group`, `api_num`, `user_group_price`, `user_group`, `account_group`, `is_wish`, `logo`, `platform_num`, `aliapp_num`, `baiduapp_num`, `phoneapp_num`, `toutiaoapp_num`, `webapp_num`, `xzapp_num` FROM `ims_site_store_goods` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteStoreGoodsCloudInsertSql         = "INSERT INTO `ims_site_store_goods_cloud` ( `cloud_id`, `name`, `title`, `logo`, `wish_branch`, `is_edited`, `isdeleted`, `branchs` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSiteStoreGoodsCloudDeleteSql         = "DELETE FROM `ims_site_store_goods_cloud` WHERE ( `id` = ? );"
	ImsSiteStoreGoodsCloudUpdateSql         = "UPDATE `ims_site_store_goods_cloud` SET `cloud_id` = ?, `name` = ?, `title` = ?, `logo` = ?, `wish_branch` = ?, `is_edited` = ?, `isdeleted` = ?, `branchs` = ? WHERE ( `id` = ? );"
	ImsSiteStoreGoodsCloudSelectSql         = "SELECT `id`, `cloud_id`, `name`, `title`, `logo`, `wish_branch`, `is_edited`, `isdeleted`, `branchs` FROM `ims_site_store_goods_cloud` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteStoreOrderInsertSql              = "INSERT INTO `ims_site_store_order` ( `orderid`, `goodsid`, `duration`, `buyer`, `buyerid`, `amount`, `type`, `changeprice`, `createtime`, `uniacid`, `endtime`, `wxapp`, `is_wish` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSiteStoreOrderDeleteSql              = "DELETE FROM `ims_site_store_order` WHERE ( `id` = ? );"
	ImsSiteStoreOrderUpdateSql              = "UPDATE `ims_site_store_order` SET `orderid` = ?, `goodsid` = ?, `duration` = ?, `buyer` = ?, `buyerid` = ?, `amount` = ?, `type` = ?, `changeprice` = ?, `createtime` = ?, `uniacid` = ?, `endtime` = ?, `wxapp` = ?, `is_wish` = ? WHERE ( `id` = ? );"
	ImsSiteStoreOrderSelectSql              = "SELECT `id`, `orderid`, `goodsid`, `duration`, `buyer`, `buyerid`, `amount`, `type`, `changeprice`, `createtime`, `uniacid`, `endtime`, `wxapp`, `is_wish` FROM `ims_site_store_order` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteStylesInsertSql                  = "INSERT INTO `ims_site_styles` ( `uniacid`, `templateid`, `name` ) VALUES ( ?, ?, ? );"
	ImsSiteStylesDeleteSql                  = "DELETE FROM `ims_site_styles` WHERE ( `id` = ? );"
	ImsSiteStylesUpdateSql                  = "UPDATE `ims_site_styles` SET `uniacid` = ?, `templateid` = ?, `name` = ? WHERE ( `id` = ? );"
	ImsSiteStylesSelectSql                  = "SELECT `id`, `uniacid`, `templateid`, `name` FROM `ims_site_styles` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteStylesVarsInsertSql              = "INSERT INTO `ims_site_styles_vars` ( `uniacid`, `templateid`, `styleid`, `variable`, `content`, `description` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsSiteStylesVarsDeleteSql              = "DELETE FROM `ims_site_styles_vars` WHERE ( `id` = ? );"
	ImsSiteStylesVarsUpdateSql              = "UPDATE `ims_site_styles_vars` SET `uniacid` = ?, `templateid` = ?, `styleid` = ?, `variable` = ?, `content` = ?, `description` = ? WHERE ( `id` = ? );"
	ImsSiteStylesVarsSelectSql              = "SELECT `id`, `uniacid`, `templateid`, `styleid`, `variable`, `content`, `description` FROM `ims_site_styles_vars` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSiteTemplatesInsertSql               = "INSERT INTO `ims_site_templates` ( `name`, `title`, `version`, `description`, `author`, `url`, `type`, `sections` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsSiteTemplatesDeleteSql               = "DELETE FROM `ims_site_templates` WHERE ( `id` = ? );"
	ImsSiteTemplatesUpdateSql               = "UPDATE `ims_site_templates` SET `name` = ?, `title` = ?, `version` = ?, `description` = ?, `author` = ?, `url` = ?, `type` = ?, `sections` = ? WHERE ( `id` = ? );"
	ImsSiteTemplatesSelectSql               = "SELECT `id`, `name`, `title`, `version`, `description`, `author`, `url`, `type`, `sections` FROM `ims_site_templates` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsStatFansInsertSql                    = "INSERT INTO `ims_stat_fans` ( `uniacid`, `new`, `cancel`, `cumulate`, `date` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsStatFansDeleteSql                    = "DELETE FROM `ims_stat_fans` WHERE ( `id` = ? );"
	ImsStatFansUpdateSql                    = "UPDATE `ims_stat_fans` SET `uniacid` = ?, `new` = ?, `cancel` = ?, `cumulate` = ?, `date` = ? WHERE ( `id` = ? );"
	ImsStatFansSelectSql                    = "SELECT `id`, `uniacid`, `new`, `cancel`, `cumulate`, `date` FROM `ims_stat_fans` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsStatKeywordInsertSql                 = "INSERT INTO `ims_stat_keyword` ( `uniacid`, `rid`, `kid`, `hit`, `lastupdate`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsStatKeywordDeleteSql                 = "DELETE FROM `ims_stat_keyword` WHERE ( `id` = ? );"
	ImsStatKeywordUpdateSql                 = "UPDATE `ims_stat_keyword` SET `uniacid` = ?, `rid` = ?, `kid` = ?, `hit` = ?, `lastupdate` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsStatKeywordSelectSql                 = "SELECT `id`, `uniacid`, `rid`, `kid`, `hit`, `lastupdate`, `createtime` FROM `ims_stat_keyword` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsStatMsgHistoryInsertSql              = "INSERT INTO `ims_stat_msg_history` ( `uniacid`, `rid`, `kid`, `from_user`, `module`, `message`, `type`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsStatMsgHistoryDeleteSql              = "DELETE FROM `ims_stat_msg_history` WHERE ( `id` = ? );"
	ImsStatMsgHistoryUpdateSql              = "UPDATE `ims_stat_msg_history` SET `uniacid` = ?, `rid` = ?, `kid` = ?, `from_user` = ?, `module` = ?, `message` = ?, `type` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsStatMsgHistorySelectSql              = "SELECT `id`, `uniacid`, `rid`, `kid`, `from_user`, `module`, `message`, `type`, `createtime` FROM `ims_stat_msg_history` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsStatRuleInsertSql                    = "INSERT INTO `ims_stat_rule` ( `uniacid`, `rid`, `hit`, `lastupdate`, `createtime` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsStatRuleDeleteSql                    = "DELETE FROM `ims_stat_rule` WHERE ( `id` = ? );"
	ImsStatRuleUpdateSql                    = "UPDATE `ims_stat_rule` SET `uniacid` = ?, `rid` = ?, `hit` = ?, `lastupdate` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsStatRuleSelectSql                    = "SELECT `id`, `uniacid`, `rid`, `hit`, `lastupdate`, `createtime` FROM `ims_stat_rule` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsStatVisitInsertSql                   = "INSERT INTO `ims_stat_visit` ( `uniacid`, `module`, `count`, `date`, `type`, `ip_count` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsStatVisitDeleteSql                   = "DELETE FROM `ims_stat_visit` WHERE ( `id` = ? );"
	ImsStatVisitUpdateSql                   = "UPDATE `ims_stat_visit` SET `uniacid` = ?, `module` = ?, `count` = ?, `date` = ?, `type` = ?, `ip_count` = ? WHERE ( `id` = ? );"
	ImsStatVisitSelectSql                   = "SELECT `id`, `uniacid`, `module`, `count`, `date`, `type`, `ip_count` FROM `ims_stat_visit` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsStatVisitIpInsertSql                 = "INSERT INTO `ims_stat_visit_ip` ( `ip`, `uniacid`, `type`, `module`, `date` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsStatVisitIpDeleteSql                 = "DELETE FROM `ims_stat_visit_ip` WHERE ( `id` = ? );"
	ImsStatVisitIpUpdateSql                 = "UPDATE `ims_stat_visit_ip` SET `ip` = ?, `uniacid` = ?, `type` = ?, `module` = ?, `date` = ? WHERE ( `id` = ? );"
	ImsStatVisitIpSelectSql                 = "SELECT `id`, `ip`, `uniacid`, `type`, `module`, `date` FROM `ims_stat_visit_ip` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSystemStatVisitInsertSql             = "INSERT INTO `ims_system_stat_visit` ( `uniacid`, `modulename`, `uid`, `displayorder`, `createtime`, `updatetime` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsSystemStatVisitDeleteSql             = "DELETE FROM `ims_system_stat_visit` WHERE ( `id` = ? );"
	ImsSystemStatVisitUpdateSql             = "UPDATE `ims_system_stat_visit` SET `uniacid` = ?, `modulename` = ?, `uid` = ?, `displayorder` = ?, `createtime` = ?, `updatetime` = ? WHERE ( `id` = ? );"
	ImsSystemStatVisitSelectSql             = "SELECT `id`, `uniacid`, `modulename`, `uid`, `displayorder`, `createtime`, `updatetime` FROM `ims_system_stat_visit` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsSystemWelcomeBinddomainInsertSql     = "INSERT INTO `ims_system_welcome_binddomain` ( `uid`, `module_name`, `domain`, `createtime` ) VALUES ( ?, ?, ?, ? );"
	ImsSystemWelcomeBinddomainDeleteSql     = "DELETE FROM `ims_system_welcome_binddomain` WHERE ( `id` = ? );"
	ImsSystemWelcomeBinddomainUpdateSql     = "UPDATE `ims_system_welcome_binddomain` SET `uid` = ?, `module_name` = ?, `domain` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsSystemWelcomeBinddomainSelectSql     = "SELECT `id`, `uid`, `module_name`, `domain`, `createtime` FROM `ims_system_welcome_binddomain` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniAccountInsertSql                  = "INSERT INTO `ims_uni_account` ( `groupid`, `name`, `description`, `default_acid`, `rank`, `title_initial`, `createtime`, `logo`, `qrcode`, `create_uid` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsUniAccountDeleteSql                  = "DELETE FROM `ims_uni_account` WHERE ( `uniacid` = ? );"
	ImsUniAccountUpdateSql                  = "UPDATE `ims_uni_account` SET `groupid` = ?, `name` = ?, `description` = ?, `default_acid` = ?, `rank` = ?, `title_initial` = ?, `createtime` = ?, `logo` = ?, `qrcode` = ?, `create_uid` = ? WHERE ( `uniacid` = ? );"
	ImsUniAccountSelectSql                  = "SELECT `uniacid`, `groupid`, `name`, `description`, `default_acid`, `rank`, `title_initial`, `createtime`, `logo`, `qrcode`, `create_uid` FROM `ims_uni_account` WHERE ( `uniacid` = ? ) ORDER BY `uniacid` DESC LIMIT 0, 1;"
	ImsUniAccountExtraModulesInsertSql      = "INSERT INTO `ims_uni_account_extra_modules` ( `uniacid`, `modules` ) VALUES ( ?, ? );"
	ImsUniAccountExtraModulesDeleteSql      = "DELETE FROM `ims_uni_account_extra_modules` WHERE ( `id` = ? );"
	ImsUniAccountExtraModulesUpdateSql      = "UPDATE `ims_uni_account_extra_modules` SET `uniacid` = ?, `modules` = ? WHERE ( `id` = ? );"
	ImsUniAccountExtraModulesSelectSql      = "SELECT `id`, `uniacid`, `modules` FROM `ims_uni_account_extra_modules` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniAccountGroupInsertSql             = "INSERT INTO `ims_uni_account_group` ( `uniacid`, `groupid` ) VALUES ( ?, ? );"
	ImsUniAccountGroupDeleteSql             = "DELETE FROM `ims_uni_account_group` WHERE ( `id` = ? );"
	ImsUniAccountGroupUpdateSql             = "UPDATE `ims_uni_account_group` SET `uniacid` = ?, `groupid` = ? WHERE ( `id` = ? );"
	ImsUniAccountGroupSelectSql             = "SELECT `id`, `uniacid`, `groupid` FROM `ims_uni_account_group` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniAccountMenusInsertSql             = "INSERT INTO `ims_uni_account_menus` ( `uniacid`, `menuid`, `type`, `title`, `sex`, `group_id`, `client_platform_type`, `area`, `data`, `status`, `createtime`, `isdeleted` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsUniAccountMenusDeleteSql             = "DELETE FROM `ims_uni_account_menus` WHERE ( `id` = ? );"
	ImsUniAccountMenusUpdateSql             = "UPDATE `ims_uni_account_menus` SET `uniacid` = ?, `menuid` = ?, `type` = ?, `title` = ?, `sex` = ?, `group_id` = ?, `client_platform_type` = ?, `area` = ?, `data` = ?, `status` = ?, `createtime` = ?, `isdeleted` = ? WHERE ( `id` = ? );"
	ImsUniAccountMenusSelectSql             = "SELECT `id`, `uniacid`, `menuid`, `type`, `title`, `sex`, `group_id`, `client_platform_type`, `area`, `data`, `status`, `createtime`, `isdeleted` FROM `ims_uni_account_menus` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniAccountModulesInsertSql           = "INSERT INTO `ims_uni_account_modules` ( `uniacid`, `module`, `enabled`, `settings`, `shortcut`, `displayorder` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsUniAccountModulesDeleteSql           = "DELETE FROM `ims_uni_account_modules` WHERE ( `id` = ? );"
	ImsUniAccountModulesUpdateSql           = "UPDATE `ims_uni_account_modules` SET `uniacid` = ?, `module` = ?, `enabled` = ?, `settings` = ?, `shortcut` = ?, `displayorder` = ? WHERE ( `id` = ? );"
	ImsUniAccountModulesSelectSql           = "SELECT `id`, `uniacid`, `module`, `enabled`, `settings`, `shortcut`, `displayorder` FROM `ims_uni_account_modules` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniAccountModulesShortcutInsertSql   = "INSERT INTO `ims_uni_account_modules_shortcut` ( `title`, `url`, `icon`, `uniacid`, `version_id`, `module_name` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsUniAccountModulesShortcutDeleteSql   = "DELETE FROM `ims_uni_account_modules_shortcut` WHERE ( `id` = ? );"
	ImsUniAccountModulesShortcutUpdateSql   = "UPDATE `ims_uni_account_modules_shortcut` SET `title` = ?, `url` = ?, `icon` = ?, `uniacid` = ?, `version_id` = ?, `module_name` = ? WHERE ( `id` = ? );"
	ImsUniAccountModulesShortcutSelectSql   = "SELECT `id`, `title`, `url`, `icon`, `uniacid`, `version_id`, `module_name` FROM `ims_uni_account_modules_shortcut` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniAccountUsersInsertSql             = "INSERT INTO `ims_uni_account_users` ( `uniacid`, `uid`, `role`, `rank` ) VALUES ( ?, ?, ?, ? );"
	ImsUniAccountUsersDeleteSql             = "DELETE FROM `ims_uni_account_users` WHERE ( `id` = ? );"
	ImsUniAccountUsersUpdateSql             = "UPDATE `ims_uni_account_users` SET `uniacid` = ?, `uid` = ?, `role` = ?, `rank` = ? WHERE ( `id` = ? );"
	ImsUniAccountUsersSelectSql             = "SELECT `id`, `uniacid`, `uid`, `role`, `rank` FROM `ims_uni_account_users` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniGroupInsertSql                    = "INSERT INTO `ims_uni_group` ( `owner_uid`, `name`, `modules`, `templates`, `uniacid`, `uid` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsUniGroupDeleteSql                    = "DELETE FROM `ims_uni_group` WHERE ( `id` = ? );"
	ImsUniGroupUpdateSql                    = "UPDATE `ims_uni_group` SET `owner_uid` = ?, `name` = ?, `modules` = ?, `templates` = ?, `uniacid` = ?, `uid` = ? WHERE ( `id` = ? );"
	ImsUniGroupSelectSql                    = "SELECT `id`, `owner_uid`, `name`, `modules`, `templates`, `uniacid`, `uid` FROM `ims_uni_group` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniLinkUniacidInsertSql              = "INSERT INTO `ims_uni_link_uniacid` ( `uniacid`, `link_uniacid`, `version_id`, `module_name` ) VALUES ( ?, ?, ?, ? );"
	ImsUniLinkUniacidDeleteSql              = "DELETE FROM `ims_uni_link_uniacid` WHERE ( `id` = ? );"
	ImsUniLinkUniacidUpdateSql              = "UPDATE `ims_uni_link_uniacid` SET `uniacid` = ?, `link_uniacid` = ?, `version_id` = ?, `module_name` = ? WHERE ( `id` = ? );"
	ImsUniLinkUniacidSelectSql              = "SELECT `id`, `uniacid`, `link_uniacid`, `version_id`, `module_name` FROM `ims_uni_link_uniacid` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniModulesInsertSql                  = "INSERT INTO `ims_uni_modules` ( `uniacid`, `module_name` ) VALUES ( ?, ? );"
	ImsUniModulesDeleteSql                  = "DELETE FROM `ims_uni_modules` WHERE ( `id` = ? );"
	ImsUniModulesUpdateSql                  = "UPDATE `ims_uni_modules` SET `uniacid` = ?, `module_name` = ? WHERE ( `id` = ? );"
	ImsUniModulesSelectSql                  = "SELECT `id`, `uniacid`, `module_name` FROM `ims_uni_modules` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUniSettingsInsertSql                 = "INSERT INTO `ims_uni_settings` ( `passport`, `oauth`, `jsauth_acid`, `notify`, `creditnames`, `creditbehaviors`, `welcome`, `default`, `default_message`, `payment`, `stat`, `default_site`, `sync`, `recharge`, `tplnotice`, `grouplevel`, `mcplugin`, `exchange_enable`, `coupon_type`, `menuset`, `statistics`, `bind_domain`, `comment_status`, `reply_setting`, `default_module`, `attachment_limit`, `attachment_size`, `sync_member`, `remote` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsUniSettingsDeleteSql                 = "DELETE FROM `ims_uni_settings` WHERE ( `uniacid` = ? );"
	ImsUniSettingsUpdateSql                 = "UPDATE `ims_uni_settings` SET `passport` = ?, `oauth` = ?, `jsauth_acid` = ?, `notify` = ?, `creditnames` = ?, `creditbehaviors` = ?, `welcome` = ?, `default` = ?, `default_message` = ?, `payment` = ?, `stat` = ?, `default_site` = ?, `sync` = ?, `recharge` = ?, `tplnotice` = ?, `grouplevel` = ?, `mcplugin` = ?, `exchange_enable` = ?, `coupon_type` = ?, `menuset` = ?, `statistics` = ?, `bind_domain` = ?, `comment_status` = ?, `reply_setting` = ?, `default_module` = ?, `attachment_limit` = ?, `attachment_size` = ?, `sync_member` = ?, `remote` = ? WHERE ( `uniacid` = ? );"
	ImsUniSettingsSelectSql                 = "SELECT `uniacid`, `passport`, `oauth`, `jsauth_acid`, `notify`, `creditnames`, `creditbehaviors`, `welcome`, `default`, `default_message`, `payment`, `stat`, `default_site`, `sync`, `recharge`, `tplnotice`, `grouplevel`, `mcplugin`, `exchange_enable`, `coupon_type`, `menuset`, `statistics`, `bind_domain`, `comment_status`, `reply_setting`, `default_module`, `attachment_limit`, `attachment_size`, `sync_member`, `remote` FROM `ims_uni_settings` WHERE ( `uniacid` = ? ) ORDER BY `uniacid` DESC LIMIT 0, 1;"
	ImsUniVerifycodeInsertSql               = "INSERT INTO `ims_uni_verifycode` ( `uniacid`, `receiver`, `verifycode`, `total`, `createtime`, `failed_count` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsUniVerifycodeDeleteSql               = "DELETE FROM `ims_uni_verifycode` WHERE ( `id` = ? );"
	ImsUniVerifycodeUpdateSql               = "UPDATE `ims_uni_verifycode` SET `uniacid` = ?, `receiver` = ?, `verifycode` = ?, `total` = ?, `createtime` = ?, `failed_count` = ? WHERE ( `id` = ? );"
	ImsUniVerifycodeSelectSql               = "SELECT `id`, `uniacid`, `receiver`, `verifycode`, `total`, `createtime`, `failed_count` FROM `ims_uni_verifycode` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUserapiCacheInsertSql                = "INSERT INTO `ims_userapi_cache` ( `key`, `content`, `lastupdate` ) VALUES ( ?, ?, ? );"
	ImsUserapiCacheDeleteSql                = "DELETE FROM `ims_userapi_cache` WHERE ( `id` = ? );"
	ImsUserapiCacheUpdateSql                = "UPDATE `ims_userapi_cache` SET `key` = ?, `content` = ?, `lastupdate` = ? WHERE ( `id` = ? );"
	ImsUserapiCacheSelectSql                = "SELECT `id`, `key`, `content`, `lastupdate` FROM `ims_userapi_cache` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUserapiReplyInsertSql                = "INSERT INTO `ims_userapi_reply` ( `rid`, `description`, `apiurl`, `token`, `default_text`, `cachetime` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsUserapiReplyDeleteSql                = "DELETE FROM `ims_userapi_reply` WHERE ( `id` = ? );"
	ImsUserapiReplyUpdateSql                = "UPDATE `ims_userapi_reply` SET `rid` = ?, `description` = ?, `apiurl` = ?, `token` = ?, `default_text` = ?, `cachetime` = ? WHERE ( `id` = ? );"
	ImsUserapiReplySelectSql                = "SELECT `id`, `rid`, `description`, `apiurl`, `token`, `default_text`, `cachetime` FROM `ims_userapi_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersInsertSql                       = "INSERT INTO `ims_users` ( `owner_uid`, `groupid`, `founder_groupid`, `username`, `password`, `salt`, `type`, `status`, `joindate`, `joinip`, `lastvisit`, `lastip`, `remark`, `starttime`, `endtime`, `register_type`, `openid`, `welcome_link`, `notice_setting`, `is_bind` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsUsersDeleteSql                       = "DELETE FROM `ims_users` WHERE ( `uid` = ? );"
	ImsUsersUpdateSql                       = "UPDATE `ims_users` SET `owner_uid` = ?, `groupid` = ?, `founder_groupid` = ?, `username` = ?, `password` = ?, `salt` = ?, `type` = ?, `status` = ?, `joindate` = ?, `joinip` = ?, `lastvisit` = ?, `lastip` = ?, `remark` = ?, `starttime` = ?, `endtime` = ?, `register_type` = ?, `openid` = ?, `welcome_link` = ?, `notice_setting` = ?, `is_bind` = ? WHERE ( `uid` = ? );"
	ImsUsersSelectSql                       = "SELECT `uid`, `owner_uid`, `groupid`, `founder_groupid`, `username`, `password`, `salt`, `type`, `status`, `joindate`, `joinip`, `lastvisit`, `lastip`, `remark`, `starttime`, `endtime`, `register_type`, `openid`, `welcome_link`, `notice_setting`, `is_bind` FROM `ims_users` WHERE ( `uid` = ? ) ORDER BY `uid` DESC LIMIT 0, 1;"
	ImsUsersBindInsertSql                   = "INSERT INTO `ims_users_bind` ( `uid`, `bind_sign`, `third_type`, `third_nickname` ) VALUES ( ?, ?, ?, ? );"
	ImsUsersBindDeleteSql                   = "DELETE FROM `ims_users_bind` WHERE ( `id` = ? );"
	ImsUsersBindUpdateSql                   = "UPDATE `ims_users_bind` SET `uid` = ?, `bind_sign` = ?, `third_type` = ?, `third_nickname` = ? WHERE ( `id` = ? );"
	ImsUsersBindSelectSql                   = "SELECT `id`, `uid`, `bind_sign`, `third_type`, `third_nickname` FROM `ims_users_bind` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersCreateGroupInsertSql            = "INSERT INTO `ims_users_create_group` ( `group_name`, `maxaccount`, `maxwxapp`, `maxwebapp`, `maxphoneapp`, `maxxzapp`, `maxaliapp`, `createtime`, `maxbaiduapp`, `maxtoutiaoapp` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsUsersCreateGroupDeleteSql            = "DELETE FROM `ims_users_create_group` WHERE ( `id` = ? );"
	ImsUsersCreateGroupUpdateSql            = "UPDATE `ims_users_create_group` SET `group_name` = ?, `maxaccount` = ?, `maxwxapp` = ?, `maxwebapp` = ?, `maxphoneapp` = ?, `maxxzapp` = ?, `maxaliapp` = ?, `createtime` = ?, `maxbaiduapp` = ?, `maxtoutiaoapp` = ? WHERE ( `id` = ? );"
	ImsUsersCreateGroupSelectSql            = "SELECT `id`, `group_name`, `maxaccount`, `maxwxapp`, `maxwebapp`, `maxphoneapp`, `maxxzapp`, `maxaliapp`, `createtime`, `maxbaiduapp`, `maxtoutiaoapp` FROM `ims_users_create_group` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersExtraGroupInsertSql             = "INSERT INTO `ims_users_extra_group` ( `uid`, `uni_group_id`, `create_group_id` ) VALUES ( ?, ?, ? );"
	ImsUsersExtraGroupDeleteSql             = "DELETE FROM `ims_users_extra_group` WHERE ( `id` = ? );"
	ImsUsersExtraGroupUpdateSql             = "UPDATE `ims_users_extra_group` SET `uid` = ?, `uni_group_id` = ?, `create_group_id` = ? WHERE ( `id` = ? );"
	ImsUsersExtraGroupSelectSql             = "SELECT `id`, `uid`, `uni_group_id`, `create_group_id` FROM `ims_users_extra_group` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersExtraLimitInsertSql             = "INSERT INTO `ims_users_extra_limit` ( `uid`, `maxaccount`, `maxwxapp`, `maxwebapp`, `maxphoneapp`, `maxxzapp`, `maxaliapp`, `timelimit`, `maxbaiduapp`, `maxtoutiaoapp` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsUsersExtraLimitDeleteSql             = "DELETE FROM `ims_users_extra_limit` WHERE ( `id` = ? );"
	ImsUsersExtraLimitUpdateSql             = "UPDATE `ims_users_extra_limit` SET `uid` = ?, `maxaccount` = ?, `maxwxapp` = ?, `maxwebapp` = ?, `maxphoneapp` = ?, `maxxzapp` = ?, `maxaliapp` = ?, `timelimit` = ?, `maxbaiduapp` = ?, `maxtoutiaoapp` = ? WHERE ( `id` = ? );"
	ImsUsersExtraLimitSelectSql             = "SELECT `id`, `uid`, `maxaccount`, `maxwxapp`, `maxwebapp`, `maxphoneapp`, `maxxzapp`, `maxaliapp`, `timelimit`, `maxbaiduapp`, `maxtoutiaoapp` FROM `ims_users_extra_limit` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersExtraModulesInsertSql           = "INSERT INTO `ims_users_extra_modules` ( `uid`, `module_name`, `support` ) VALUES ( ?, ?, ? );"
	ImsUsersExtraModulesDeleteSql           = "DELETE FROM `ims_users_extra_modules` WHERE ( `id` = ? );"
	ImsUsersExtraModulesUpdateSql           = "UPDATE `ims_users_extra_modules` SET `uid` = ?, `module_name` = ?, `support` = ? WHERE ( `id` = ? );"
	ImsUsersExtraModulesSelectSql           = "SELECT `id`, `uid`, `module_name`, `support` FROM `ims_users_extra_modules` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersExtraTemplatesInsertSql         = "INSERT INTO `ims_users_extra_templates` ( `uid`, `template_id` ) VALUES ( ?, ? );"
	ImsUsersExtraTemplatesDeleteSql         = "DELETE FROM `ims_users_extra_templates` WHERE ( `id` = ? );"
	ImsUsersExtraTemplatesUpdateSql         = "UPDATE `ims_users_extra_templates` SET `uid` = ?, `template_id` = ? WHERE ( `id` = ? );"
	ImsUsersExtraTemplatesSelectSql         = "SELECT `id`, `uid`, `template_id` FROM `ims_users_extra_templates` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersFailedLoginInsertSql            = "INSERT INTO `ims_users_failed_login` ( `ip`, `username`, `count`, `lastupdate` ) VALUES ( ?, ?, ?, ? );"
	ImsUsersFailedLoginDeleteSql            = "DELETE FROM `ims_users_failed_login` WHERE ( `id` = ? );"
	ImsUsersFailedLoginUpdateSql            = "UPDATE `ims_users_failed_login` SET `ip` = ?, `username` = ?, `count` = ?, `lastupdate` = ? WHERE ( `id` = ? );"
	ImsUsersFailedLoginSelectSql            = "SELECT `id`, `ip`, `username`, `count`, `lastupdate` FROM `ims_users_failed_login` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersFounderGroupInsertSql           = "INSERT INTO `ims_users_founder_group` ( `name`, `package`, `maxaccount`, `timelimit`, `maxwxapp`, `maxwebapp`, `maxphoneapp`, `maxxzapp`, `maxaliapp`, `maxbaiduapp`, `maxtoutiaoapp` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsUsersFounderGroupDeleteSql           = "DELETE FROM `ims_users_founder_group` WHERE ( `id` = ? );"
	ImsUsersFounderGroupUpdateSql           = "UPDATE `ims_users_founder_group` SET `name` = ?, `package` = ?, `maxaccount` = ?, `timelimit` = ?, `maxwxapp` = ?, `maxwebapp` = ?, `maxphoneapp` = ?, `maxxzapp` = ?, `maxaliapp` = ?, `maxbaiduapp` = ?, `maxtoutiaoapp` = ? WHERE ( `id` = ? );"
	ImsUsersFounderGroupSelectSql           = "SELECT `id`, `name`, `package`, `maxaccount`, `timelimit`, `maxwxapp`, `maxwebapp`, `maxphoneapp`, `maxxzapp`, `maxaliapp`, `maxbaiduapp`, `maxtoutiaoapp` FROM `ims_users_founder_group` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersFounderOwnCreateGroupsInsertSql = "INSERT INTO `ims_users_founder_own_create_groups` ( `founder_uid`, `create_group_id` ) VALUES ( ?, ? );"
	ImsUsersFounderOwnCreateGroupsDeleteSql = "DELETE FROM `ims_users_founder_own_create_groups` WHERE ( `id` = ? );"
	ImsUsersFounderOwnCreateGroupsUpdateSql = "UPDATE `ims_users_founder_own_create_groups` SET `founder_uid` = ?, `create_group_id` = ? WHERE ( `id` = ? );"
	ImsUsersFounderOwnCreateGroupsSelectSql = "SELECT `id`, `founder_uid`, `create_group_id` FROM `ims_users_founder_own_create_groups` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersFounderOwnUniGroupsInsertSql    = "INSERT INTO `ims_users_founder_own_uni_groups` ( `founder_uid`, `uni_group_id` ) VALUES ( ?, ? );"
	ImsUsersFounderOwnUniGroupsDeleteSql    = "DELETE FROM `ims_users_founder_own_uni_groups` WHERE ( `id` = ? );"
	ImsUsersFounderOwnUniGroupsUpdateSql    = "UPDATE `ims_users_founder_own_uni_groups` SET `founder_uid` = ?, `uni_group_id` = ? WHERE ( `id` = ? );"
	ImsUsersFounderOwnUniGroupsSelectSql    = "SELECT `id`, `founder_uid`, `uni_group_id` FROM `ims_users_founder_own_uni_groups` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersFounderOwnUsersInsertSql        = "INSERT INTO `ims_users_founder_own_users` ( `uid`, `founder_uid` ) VALUES ( ?, ? );"
	ImsUsersFounderOwnUsersDeleteSql        = "DELETE FROM `ims_users_founder_own_users` WHERE ( `id` = ? );"
	ImsUsersFounderOwnUsersUpdateSql        = "UPDATE `ims_users_founder_own_users` SET `uid` = ?, `founder_uid` = ? WHERE ( `id` = ? );"
	ImsUsersFounderOwnUsersSelectSql        = "SELECT `id`, `uid`, `founder_uid` FROM `ims_users_founder_own_users` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersFounderOwnUsersGroupsInsertSql  = "INSERT INTO `ims_users_founder_own_users_groups` ( `founder_uid`, `users_group_id` ) VALUES ( ?, ? );"
	ImsUsersFounderOwnUsersGroupsDeleteSql  = "DELETE FROM `ims_users_founder_own_users_groups` WHERE ( `id` = ? );"
	ImsUsersFounderOwnUsersGroupsUpdateSql  = "UPDATE `ims_users_founder_own_users_groups` SET `founder_uid` = ?, `users_group_id` = ? WHERE ( `id` = ? );"
	ImsUsersFounderOwnUsersGroupsSelectSql  = "SELECT `id`, `founder_uid`, `users_group_id` FROM `ims_users_founder_own_users_groups` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersGroupInsertSql                  = "INSERT INTO `ims_users_group` ( `owner_uid`, `name`, `package`, `maxaccount`, `timelimit`, `maxwxapp`, `maxwebapp`, `maxphoneapp`, `maxxzapp`, `maxaliapp`, `maxbaiduapp`, `maxtoutiaoapp` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsUsersGroupDeleteSql                  = "DELETE FROM `ims_users_group` WHERE ( `id` = ? );"
	ImsUsersGroupUpdateSql                  = "UPDATE `ims_users_group` SET `owner_uid` = ?, `name` = ?, `package` = ?, `maxaccount` = ?, `timelimit` = ?, `maxwxapp` = ?, `maxwebapp` = ?, `maxphoneapp` = ?, `maxxzapp` = ?, `maxaliapp` = ?, `maxbaiduapp` = ?, `maxtoutiaoapp` = ? WHERE ( `id` = ? );"
	ImsUsersGroupSelectSql                  = "SELECT `id`, `owner_uid`, `name`, `package`, `maxaccount`, `timelimit`, `maxwxapp`, `maxwebapp`, `maxphoneapp`, `maxxzapp`, `maxaliapp`, `maxbaiduapp`, `maxtoutiaoapp` FROM `ims_users_group` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersInvitationInsertSql             = "INSERT INTO `ims_users_invitation` ( `code`, `fromuid`, `inviteuid`, `createtime` ) VALUES ( ?, ?, ?, ? );"
	ImsUsersInvitationDeleteSql             = "DELETE FROM `ims_users_invitation` WHERE ( `id` = ? );"
	ImsUsersInvitationUpdateSql             = "UPDATE `ims_users_invitation` SET `code` = ?, `fromuid` = ?, `inviteuid` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsUsersInvitationSelectSql             = "SELECT `id`, `code`, `fromuid`, `inviteuid`, `createtime` FROM `ims_users_invitation` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersLastuseInsertSql                = "INSERT INTO `ims_users_lastuse` ( `uid`, `uniacid`, `modulename`, `type` ) VALUES ( ?, ?, ?, ? );"
	ImsUsersLastuseDeleteSql                = "DELETE FROM `ims_users_lastuse` WHERE ( `id` = ? );"
	ImsUsersLastuseUpdateSql                = "UPDATE `ims_users_lastuse` SET `uid` = ?, `uniacid` = ?, `modulename` = ?, `type` = ? WHERE ( `id` = ? );"
	ImsUsersLastuseSelectSql                = "SELECT `id`, `uid`, `uniacid`, `modulename`, `type` FROM `ims_users_lastuse` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersLoginLogsInsertSql              = "INSERT INTO `ims_users_login_logs` ( `uid`, `ip`, `city`, `login_at`, `createtime` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsUsersLoginLogsDeleteSql              = "DELETE FROM `ims_users_login_logs` WHERE ( `id` = ? );"
	ImsUsersLoginLogsUpdateSql              = "UPDATE `ims_users_login_logs` SET `uid` = ?, `ip` = ?, `city` = ?, `login_at` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsUsersLoginLogsSelectSql              = "SELECT `id`, `uid`, `ip`, `city`, `login_at`, `createtime` FROM `ims_users_login_logs` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersOperateHistoryInsertSql         = "INSERT INTO `ims_users_operate_history` ( `type`, `uid`, `uniacid`, `module_name`, `createtime` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsUsersOperateHistoryDeleteSql         = "DELETE FROM `ims_users_operate_history` WHERE ( `id` = ? );"
	ImsUsersOperateHistoryUpdateSql         = "UPDATE `ims_users_operate_history` SET `type` = ?, `uid` = ?, `uniacid` = ?, `module_name` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsUsersOperateHistorySelectSql         = "SELECT `id`, `type`, `uid`, `uniacid`, `module_name`, `createtime` FROM `ims_users_operate_history` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersOperateStarInsertSql            = "INSERT INTO `ims_users_operate_star` ( `type`, `uid`, `uniacid`, `module_name`, `rank`, `createtime` ) VALUES ( ?, ?, ?, ?, ?, ? );"
	ImsUsersOperateStarDeleteSql            = "DELETE FROM `ims_users_operate_star` WHERE ( `id` = ? );"
	ImsUsersOperateStarUpdateSql            = "UPDATE `ims_users_operate_star` SET `type` = ?, `uid` = ?, `uniacid` = ?, `module_name` = ?, `rank` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsUsersOperateStarSelectSql            = "SELECT `id`, `type`, `uid`, `uniacid`, `module_name`, `rank`, `createtime` FROM `ims_users_operate_star` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersPermissionInsertSql             = "INSERT INTO `ims_users_permission` ( `uniacid`, `uid`, `type`, `permission`, `url`, `modules`, `templates` ) VALUES ( ?, ?, ?, ?, ?, ?, ? );"
	ImsUsersPermissionDeleteSql             = "DELETE FROM `ims_users_permission` WHERE ( `id` = ? );"
	ImsUsersPermissionUpdateSql             = "UPDATE `ims_users_permission` SET `uniacid` = ?, `uid` = ?, `type` = ?, `permission` = ?, `url` = ?, `modules` = ?, `templates` = ? WHERE ( `id` = ? );"
	ImsUsersPermissionSelectSql             = "SELECT `id`, `uniacid`, `uid`, `type`, `permission`, `url`, `modules`, `templates` FROM `ims_users_permission` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsUsersProfileInsertSql                = "INSERT INTO `ims_users_profile` ( `uid`, `createtime`, `edittime`, `realname`, `nickname`, `avatar`, `qq`, `mobile`, `fakeid`, `vip`, `gender`, `birthyear`, `birthmonth`, `birthday`, `constellation`, `zodiac`, `telephone`, `idcard`, `studentid`, `grade`, `address`, `zipcode`, `nationality`, `resideprovince`, `residecity`, `residedist`, `graduateschool`, `company`, `education`, `occupation`, `position`, `revenue`, `affectivestatus`, `lookingfor`, `bloodtype`, `height`, `weight`, `alipay`, `msn`, `email`, `taobao`, `site`, `bio`, `interest`, `workerid`, `is_send_mobile_status`, `send_expire_status` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsUsersProfileDeleteSql                = "DELETE FROM `ims_users_profile` WHERE ( `id` = ? );"
	ImsUsersProfileUpdateSql                = "UPDATE `ims_users_profile` SET `uid` = ?, `createtime` = ?, `edittime` = ?, `realname` = ?, `nickname` = ?, `avatar` = ?, `qq` = ?, `mobile` = ?, `fakeid` = ?, `vip` = ?, `gender` = ?, `birthyear` = ?, `birthmonth` = ?, `birthday` = ?, `constellation` = ?, `zodiac` = ?, `telephone` = ?, `idcard` = ?, `studentid` = ?, `grade` = ?, `address` = ?, `zipcode` = ?, `nationality` = ?, `resideprovince` = ?, `residecity` = ?, `residedist` = ?, `graduateschool` = ?, `company` = ?, `education` = ?, `occupation` = ?, `position` = ?, `revenue` = ?, `affectivestatus` = ?, `lookingfor` = ?, `bloodtype` = ?, `height` = ?, `weight` = ?, `alipay` = ?, `msn` = ?, `email` = ?, `taobao` = ?, `site` = ?, `bio` = ?, `interest` = ?, `workerid` = ?, `is_send_mobile_status` = ?, `send_expire_status` = ? WHERE ( `id` = ? );"
	ImsUsersProfileSelectSql                = "SELECT `id`, `uid`, `createtime`, `edittime`, `realname`, `nickname`, `avatar`, `qq`, `mobile`, `fakeid`, `vip`, `gender`, `birthyear`, `birthmonth`, `birthday`, `constellation`, `zodiac`, `telephone`, `idcard`, `studentid`, `grade`, `address`, `zipcode`, `nationality`, `resideprovince`, `residecity`, `residedist`, `graduateschool`, `company`, `education`, `occupation`, `position`, `revenue`, `affectivestatus`, `lookingfor`, `bloodtype`, `height`, `weight`, `alipay`, `msn`, `email`, `taobao`, `site`, `bio`, `interest`, `workerid`, `is_send_mobile_status`, `send_expire_status` FROM `ims_users_profile` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVideoReplyInsertSql                  = "INSERT INTO `ims_video_reply` ( `rid`, `title`, `description`, `mediaid`, `createtime` ) VALUES ( ?, ?, ?, ?, ? );"
	ImsVideoReplyDeleteSql                  = "DELETE FROM `ims_video_reply` WHERE ( `id` = ? );"
	ImsVideoReplyUpdateSql                  = "UPDATE `ims_video_reply` SET `rid` = ?, `title` = ?, `description` = ?, `mediaid` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsVideoReplySelectSql                  = "SELECT `id`, `rid`, `title`, `description`, `mediaid`, `createtime` FROM `ims_video_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVoiceReplyInsertSql                  = "INSERT INTO `ims_voice_reply` ( `rid`, `title`, `mediaid`, `createtime` ) VALUES ( ?, ?, ?, ? );"
	ImsVoiceReplyDeleteSql                  = "DELETE FROM `ims_voice_reply` WHERE ( `id` = ? );"
	ImsVoiceReplyUpdateSql                  = "UPDATE `ims_voice_reply` SET `rid` = ?, `title` = ?, `mediaid` = ?, `createtime` = ? WHERE ( `id` = ? );"
	ImsVoiceReplySelectSql                  = "SELECT `id`, `rid`, `title`, `mediaid`, `createtime` FROM `ims_voice_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanInsertSql                       = "INSERT INTO `ims_vx_fan` ( `pid`, `account_id`, `account_category`, `openid`, `group_id`, `union_id`, `mobile`, `category`, `category_msg`, `subscribe`, `nickname`, `sex`, `language`, `city`, `province`, `country`, `avatar`, `subscribe_at`, `state`, `state_msg`, `created_at`, `updated_at`, `note`, `recommend_qr_code` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanDeleteSql                       = "DELETE FROM `ims_vx_fan` WHERE ( `id` = ? );"
	ImsVxFanUpdateSql                       = "UPDATE `ims_vx_fan` SET `pid` = ?, `account_id` = ?, `account_category` = ?, `openid` = ?, `group_id` = ?, `union_id` = ?, `mobile` = ?, `category` = ?, `category_msg` = ?, `subscribe` = ?, `nickname` = ?, `sex` = ?, `language` = ?, `city` = ?, `province` = ?, `country` = ?, `avatar` = ?, `subscribe_at` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ?, `recommend_qr_code` = ? WHERE ( `id` = ? );"
	ImsVxFanSelectSql                       = "SELECT `id`, `pid`, `account_id`, `account_category`, `openid`, `group_id`, `union_id`, `mobile`, `category`, `category_msg`, `subscribe`, `nickname`, `sex`, `language`, `city`, `province`, `country`, `avatar`, `subscribe_at`, `state`, `state_msg`, `created_at`, `updated_at`, `note`, `recommend_qr_code` FROM `ims_vx_fan` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanAccountInsertSql                = "INSERT INTO `ims_vx_fan_account` ( `uid`, `balance`, `prepare`, `withdraw`, `last_withdraw_at`, `last_withdraw_day`, `state`, `state_msg`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanAccountDeleteSql                = "DELETE FROM `ims_vx_fan_account` WHERE ( `id` = ? );"
	ImsVxFanAccountUpdateSql                = "UPDATE `ims_vx_fan_account` SET `uid` = ?, `balance` = ?, `prepare` = ?, `withdraw` = ?, `last_withdraw_at` = ?, `last_withdraw_day` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanAccountSelectSql                = "SELECT `id`, `uid`, `balance`, `prepare`, `withdraw`, `last_withdraw_at`, `last_withdraw_day`, `state`, `state_msg`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_account` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanAddressInsertSql                = "INSERT INTO `ims_vx_fan_address` ( `uid`, `consignee`, `mobile`, `address1`, `address2`, `address3`, `address4`, `postal_code`, `is_default`, `state`, `state_msg`, `created_at`, `updated_at`, `deleted_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanAddressDeleteSql                = "DELETE FROM `ims_vx_fan_address` WHERE ( `id` = ? );"
	ImsVxFanAddressUpdateSql                = "UPDATE `ims_vx_fan_address` SET `uid` = ?, `consignee` = ?, `mobile` = ?, `address1` = ?, `address2` = ?, `address3` = ?, `address4` = ?, `postal_code` = ?, `is_default` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `deleted_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanAddressSelectSql                = "SELECT `id`, `uid`, `consignee`, `mobile`, `address1`, `address2`, `address3`, `address4`, `postal_code`, `is_default`, `state`, `state_msg`, `created_at`, `updated_at`, `deleted_at`, `note` FROM `ims_vx_fan_address` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanBalanceLogInsertSql             = "INSERT INTO `ims_vx_fan_balance_log` ( `uid`, `val1`, `val2`, `val3`, `category`, `category_msg`, `order_id`, `order_no`, `state`, `state_msg`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanBalanceLogDeleteSql             = "DELETE FROM `ims_vx_fan_balance_log` WHERE ( `id` = ? );"
	ImsVxFanBalanceLogUpdateSql             = "UPDATE `ims_vx_fan_balance_log` SET `uid` = ?, `val1` = ?, `val2` = ?, `val3` = ?, `category` = ?, `category_msg` = ?, `order_id` = ?, `order_no` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanBalanceLogSelectSql             = "SELECT `id`, `uid`, `val1`, `val2`, `val3`, `category`, `category_msg`, `order_id`, `order_no`, `state`, `state_msg`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_balance_log` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanGoodsInsertSql                  = "INSERT INTO `ims_vx_fan_goods` ( `account_id`, `avatar`, `name`, `title`, `title_desc`, `label`, `details`, `video`, `price_nominal`, `price`, `postage`, `stock`, `sold`, `serial`, `category`, `category_msg`, `uid`, `delivery_address`, `recommend`, `state`, `state_msg`, `created_at`, `updated_at`, `deleted_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanGoodsDeleteSql                  = "DELETE FROM `ims_vx_fan_goods` WHERE ( `id` = ? );"
	ImsVxFanGoodsUpdateSql                  = "UPDATE `ims_vx_fan_goods` SET `account_id` = ?, `avatar` = ?, `name` = ?, `title` = ?, `title_desc` = ?, `label` = ?, `details` = ?, `video` = ?, `price_nominal` = ?, `price` = ?, `postage` = ?, `stock` = ?, `sold` = ?, `serial` = ?, `category` = ?, `category_msg` = ?, `uid` = ?, `delivery_address` = ?, `recommend` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `deleted_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanGoodsSelectSql                  = "SELECT `id`, `account_id`, `avatar`, `name`, `title`, `title_desc`, `label`, `details`, `video`, `price_nominal`, `price`, `postage`, `stock`, `sold`, `serial`, `category`, `category_msg`, `uid`, `delivery_address`, `recommend`, `state`, `state_msg`, `created_at`, `updated_at`, `deleted_at`, `note` FROM `ims_vx_fan_goods` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanGoodsRuleInsertSql              = "INSERT INTO `ims_vx_fan_goods_rule` ( `goods_id`, `attributes`, `serial`, `avatar`, `details`, `price_nominal`, `price`, `stock`, `sold`, `state`, `state_msg`, `created_at`, `updated_at`, `deleted_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanGoodsRuleDeleteSql              = "DELETE FROM `ims_vx_fan_goods_rule` WHERE ( `id` = ? );"
	ImsVxFanGoodsRuleUpdateSql              = "UPDATE `ims_vx_fan_goods_rule` SET `goods_id` = ?, `attributes` = ?, `serial` = ?, `avatar` = ?, `details` = ?, `price_nominal` = ?, `price` = ?, `stock` = ?, `sold` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `deleted_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanGoodsRuleSelectSql              = "SELECT `id`, `goods_id`, `attributes`, `serial`, `avatar`, `details`, `price_nominal`, `price`, `stock`, `sold`, `state`, `state_msg`, `created_at`, `updated_at`, `deleted_at`, `note` FROM `ims_vx_fan_goods_rule` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanOrderBagQuotaCardInsertSql      = "INSERT INTO `ims_vx_fan_order_bag_quota_card` ( `uid`, `account_id`, `goods_id`, `goods_num`, `goods_price`, `goods_title`, `order_no`, `amount`, `pay_amount`, `pay_type`, `pay_state`, `pay_no`, `pay_details`, `goods_details`, `user_details`, `state`, `state1`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanOrderBagQuotaCardDeleteSql      = "DELETE FROM `ims_vx_fan_order_bag_quota_card` WHERE ( `id` = ? );"
	ImsVxFanOrderBagQuotaCardUpdateSql      = "UPDATE `ims_vx_fan_order_bag_quota_card` SET `uid` = ?, `account_id` = ?, `goods_id` = ?, `goods_num` = ?, `goods_price` = ?, `goods_title` = ?, `order_no` = ?, `amount` = ?, `pay_amount` = ?, `pay_type` = ?, `pay_state` = ?, `pay_no` = ?, `pay_details` = ?, `goods_details` = ?, `user_details` = ?, `state` = ?, `state1` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanOrderBagQuotaCardSelectSql      = "SELECT `id`, `uid`, `account_id`, `goods_id`, `goods_num`, `goods_price`, `goods_title`, `order_no`, `amount`, `pay_amount`, `pay_type`, `pay_state`, `pay_no`, `pay_details`, `goods_details`, `user_details`, `state`, `state1`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_order_bag_quota_card` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanOrderGoodsInsertSql             = "INSERT INTO `ims_vx_fan_order_goods` ( `uid`, `account_id`, `goods_id`, `goods_num`, `goods_price`, `order_no`, `amount`, `discount_amount`, `discount`, `pay_amount`, `pay_type`, `pay_state`, `pay_no`, `pay_details`, `user_details`, `goods_details`, `address_details`, `address_id`, `address`, `state`, `state1`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanOrderGoodsDeleteSql             = "DELETE FROM `ims_vx_fan_order_goods` WHERE ( `id` = ? );"
	ImsVxFanOrderGoodsUpdateSql             = "UPDATE `ims_vx_fan_order_goods` SET `uid` = ?, `account_id` = ?, `goods_id` = ?, `goods_num` = ?, `goods_price` = ?, `order_no` = ?, `amount` = ?, `discount_amount` = ?, `discount` = ?, `pay_amount` = ?, `pay_type` = ?, `pay_state` = ?, `pay_no` = ?, `pay_details` = ?, `user_details` = ?, `goods_details` = ?, `address_details` = ?, `address_id` = ?, `address` = ?, `state` = ?, `state1` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanOrderGoodsSelectSql             = "SELECT `id`, `uid`, `account_id`, `goods_id`, `goods_num`, `goods_price`, `order_no`, `amount`, `discount_amount`, `discount`, `pay_amount`, `pay_type`, `pay_state`, `pay_no`, `pay_details`, `user_details`, `goods_details`, `address_details`, `address_id`, `address`, `state`, `state1`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_order_goods` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanOrderRightCardInsertSql         = "INSERT INTO `ims_vx_fan_order_right_card` ( `uid`, `account_id`, `goods_id`, `goods_num`, `goods_price`, `goods_title`, `order_no`, `amount`, `pay_amount`, `pay_type`, `pay_state`, `pay_no`, `pay_details`, `goods_details`, `user_details`, `state`, `state1`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanOrderRightCardDeleteSql         = "DELETE FROM `ims_vx_fan_order_right_card` WHERE ( `id` = ? );"
	ImsVxFanOrderRightCardUpdateSql         = "UPDATE `ims_vx_fan_order_right_card` SET `uid` = ?, `account_id` = ?, `goods_id` = ?, `goods_num` = ?, `goods_price` = ?, `goods_title` = ?, `order_no` = ?, `amount` = ?, `pay_amount` = ?, `pay_type` = ?, `pay_state` = ?, `pay_no` = ?, `pay_details` = ?, `goods_details` = ?, `user_details` = ?, `state` = ?, `state1` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanOrderRightCardSelectSql         = "SELECT `id`, `uid`, `account_id`, `goods_id`, `goods_num`, `goods_price`, `goods_title`, `order_no`, `amount`, `pay_amount`, `pay_type`, `pay_state`, `pay_no`, `pay_details`, `goods_details`, `user_details`, `state`, `state1`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_order_right_card` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanPrepareLogInsertSql             = "INSERT INTO `ims_vx_fan_prepare_log` ( `uid`, `val1`, `val2`, `val3`, `category`, `category_msg`, `order_id`, `order_no`, `state`, `state_msg`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanPrepareLogDeleteSql             = "DELETE FROM `ims_vx_fan_prepare_log` WHERE ( `id` = ? );"
	ImsVxFanPrepareLogUpdateSql             = "UPDATE `ims_vx_fan_prepare_log` SET `uid` = ?, `val1` = ?, `val2` = ?, `val3` = ?, `category` = ?, `category_msg` = ?, `order_id` = ?, `order_no` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanPrepareLogSelectSql             = "SELECT `id`, `uid`, `val1`, `val2`, `val3`, `category`, `category_msg`, `order_id`, `order_no`, `state`, `state_msg`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_prepare_log` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanRightCardInsertSql              = "INSERT INTO `ims_vx_fan_right_card` ( `account_id`, `name`, `money`, `avatar`, `title`, `details`, `multiple`, `vip`, `discount`, `serial`, `state`, `state_msg`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanRightCardDeleteSql              = "DELETE FROM `ims_vx_fan_right_card` WHERE ( `id` = ? );"
	ImsVxFanRightCardUpdateSql              = "UPDATE `ims_vx_fan_right_card` SET `account_id` = ?, `name` = ?, `money` = ?, `avatar` = ?, `title` = ?, `details` = ?, `multiple` = ?, `vip` = ?, `discount` = ?, `serial` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanRightCardSelectSql              = "SELECT `id`, `account_id`, `name`, `money`, `avatar`, `title`, `details`, `multiple`, `vip`, `discount`, `serial`, `state`, `state_msg`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_right_card` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanRightCardUpgradeInsertSql       = "INSERT INTO `ims_vx_fan_right_card_upgrade` ( `uid`, `upgrade_ratio`, `from_uid`, `right_card_amount`, `state`, `state_msg`, `created_at`, `updated_at`, `note`, `cards` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanRightCardUpgradeDeleteSql       = "DELETE FROM `ims_vx_fan_right_card_upgrade` WHERE ( `id` = ? );"
	ImsVxFanRightCardUpgradeUpdateSql       = "UPDATE `ims_vx_fan_right_card_upgrade` SET `uid` = ?, `upgrade_ratio` = ?, `from_uid` = ?, `right_card_amount` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ?, `cards` = ? WHERE ( `id` = ? );"
	ImsVxFanRightCardUpgradeSelectSql       = "SELECT `id`, `uid`, `upgrade_ratio`, `from_uid`, `right_card_amount`, `state`, `state_msg`, `created_at`, `updated_at`, `note`, `cards` FROM `ims_vx_fan_right_card_upgrade` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanRightCardUsedInsertSql          = "INSERT INTO `ims_vx_fan_right_card_used` ( `cid`, `uid`, `serial`, `name`, `money`, `multiple`, `amount`, `days`, `amount_remain`, `days_remain`, `day_ratio`, `day_money`, `amount_achieved`, `amount_achieved_days`, `discount`, `state`, `state_msg`, `created_at`, `updated_at`, `note`, `tips`, `card` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanRightCardUsedDeleteSql          = "DELETE FROM `ims_vx_fan_right_card_used` WHERE ( `id` = ? );"
	ImsVxFanRightCardUsedUpdateSql          = "UPDATE `ims_vx_fan_right_card_used` SET `cid` = ?, `uid` = ?, `serial` = ?, `name` = ?, `money` = ?, `multiple` = ?, `amount` = ?, `days` = ?, `amount_remain` = ?, `days_remain` = ?, `day_ratio` = ?, `day_money` = ?, `amount_achieved` = ?, `amount_achieved_days` = ?, `discount` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ?, `tips` = ?, `card` = ? WHERE ( `id` = ? );"
	ImsVxFanRightCardUsedSelectSql          = "SELECT `id`, `cid`, `uid`, `serial`, `name`, `money`, `multiple`, `amount`, `days`, `amount_remain`, `days_remain`, `day_ratio`, `day_money`, `amount_achieved`, `amount_achieved_days`, `discount`, `state`, `state_msg`, `created_at`, `updated_at`, `note`, `tips`, `card` FROM `ims_vx_fan_right_card_used` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanRightCardUsingInsertSql         = "INSERT INTO `ims_vx_fan_right_card_using` ( `cid`, `uid`, `serial`, `name`, `money`, `multiple`, `amount`, `days`, `amount_remain`, `days_remain`, `day_ratio`, `day_money`, `amount_achieved`, `amount_achieved_days`, `discount`, `state`, `state_msg`, `created_at`, `updated_at`, `note`, `tips`, `card` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanRightCardUsingDeleteSql         = "DELETE FROM `ims_vx_fan_right_card_using` WHERE ( `id` = ? );"
	ImsVxFanRightCardUsingUpdateSql         = "UPDATE `ims_vx_fan_right_card_using` SET `cid` = ?, `uid` = ?, `serial` = ?, `name` = ?, `money` = ?, `multiple` = ?, `amount` = ?, `days` = ?, `amount_remain` = ?, `days_remain` = ?, `day_ratio` = ?, `day_money` = ?, `amount_achieved` = ?, `amount_achieved_days` = ?, `discount` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ?, `tips` = ?, `card` = ? WHERE ( `id` = ? );"
	ImsVxFanRightCardUsingSelectSql         = "SELECT `id`, `cid`, `uid`, `serial`, `name`, `money`, `multiple`, `amount`, `days`, `amount_remain`, `days_remain`, `day_ratio`, `day_money`, `amount_achieved`, `amount_achieved_days`, `discount`, `state`, `state_msg`, `created_at`, `updated_at`, `note`, `tips`, `card` FROM `ims_vx_fan_right_card_using` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanSignInsertSql                   = "INSERT INTO `ims_vx_fan_sign` ( `uid`, `val1`, `val2`, `val3`, `day`, `continued_days`, `state`, `state_msg`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanSignDeleteSql                   = "DELETE FROM `ims_vx_fan_sign` WHERE ( `id` = ? );"
	ImsVxFanSignUpdateSql                   = "UPDATE `ims_vx_fan_sign` SET `uid` = ?, `val1` = ?, `val2` = ?, `val3` = ?, `day` = ?, `continued_days` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanSignSelectSql                   = "SELECT `id`, `uid`, `val1`, `val2`, `val3`, `day`, `continued_days`, `state`, `state_msg`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_sign` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanSignAdvertisingInsertSql        = "INSERT INTO `ims_vx_fan_sign_advertising` ( `account_id`, `name`, `avatar`, `video`, `details`, `duration`, `category`, `category_msg`, `uid`, `state`, `state_msg`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanSignAdvertisingDeleteSql        = "DELETE FROM `ims_vx_fan_sign_advertising` WHERE ( `id` = ? );"
	ImsVxFanSignAdvertisingUpdateSql        = "UPDATE `ims_vx_fan_sign_advertising` SET `account_id` = ?, `name` = ?, `avatar` = ?, `video` = ?, `details` = ?, `duration` = ?, `category` = ?, `category_msg` = ?, `uid` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanSignAdvertisingSelectSql        = "SELECT `id`, `account_id`, `name`, `avatar`, `video`, `details`, `duration`, `category`, `category_msg`, `uid`, `state`, `state_msg`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_sign_advertising` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanVipInsertSql                    = "INSERT INTO `ims_vx_fan_vip` ( `uid`, `vip`, `vip_msg`, `state`, `state_msg`, `created_at`, `updated_at`, `expired_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanVipDeleteSql                    = "DELETE FROM `ims_vx_fan_vip` WHERE ( `id` = ? );"
	ImsVxFanVipUpdateSql                    = "UPDATE `ims_vx_fan_vip` SET `uid` = ?, `vip` = ?, `vip_msg` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `expired_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanVipSelectSql                    = "SELECT `id`, `uid`, `vip`, `vip_msg`, `state`, `state_msg`, `created_at`, `updated_at`, `expired_at`, `note` FROM `ims_vx_fan_vip` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxFanWithdrawLogInsertSql            = "INSERT INTO `ims_vx_fan_withdraw_log` ( `uid`, `val1`, `val2`, `val3`, `category`, `category_msg`, `order_id`, `order_no`, `state`, `state_msg`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxFanWithdrawLogDeleteSql            = "DELETE FROM `ims_vx_fan_withdraw_log` WHERE ( `id` = ? );"
	ImsVxFanWithdrawLogUpdateSql            = "UPDATE `ims_vx_fan_withdraw_log` SET `uid` = ?, `val1` = ?, `val2` = ?, `val3` = ?, `category` = ?, `category_msg` = ?, `order_id` = ?, `order_no` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxFanWithdrawLogSelectSql            = "SELECT `id`, `uid`, `val1`, `val2`, `val3`, `category`, `category_msg`, `order_id`, `order_no`, `state`, `state_msg`, `created_at`, `updated_at`, `note` FROM `ims_vx_fan_withdraw_log` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxRedBagInsertSql                    = "INSERT INTO `ims_vx_red_bag` ( `uid`, `account_id`, `account_app_id`, `mch_id`, `openid`, `amount`, `send_list_id`, `mch_no`, `return_code`, `return_msg`, `result_code`, `err_code`, `err_code_des`, `day`, `state`, `state_msg`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxRedBagDeleteSql                    = "DELETE FROM `ims_vx_red_bag` WHERE ( `id` = ? );"
	ImsVxRedBagUpdateSql                    = "UPDATE `ims_vx_red_bag` SET `uid` = ?, `account_id` = ?, `account_app_id` = ?, `mch_id` = ?, `openid` = ?, `amount` = ?, `send_list_id` = ?, `mch_no` = ?, `return_code` = ?, `return_msg` = ?, `result_code` = ?, `err_code` = ?, `err_code_des` = ?, `day` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxRedBagSelectSql                    = "SELECT `id`, `uid`, `account_id`, `account_app_id`, `mch_id`, `openid`, `amount`, `send_list_id`, `mch_no`, `return_code`, `return_msg`, `result_code`, `err_code`, `err_code_des`, `day`, `state`, `state_msg`, `created_at`, `updated_at`, `note` FROM `ims_vx_red_bag` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsVxRedBagQuotaCardInsertSql           = "INSERT INTO `ims_vx_red_bag_quota_card` ( `account_id`, `name`, `avatar`, `title`, `details`, `money`, `prepare`, `state`, `state_msg`, `created_at`, `updated_at`, `note` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsVxRedBagQuotaCardDeleteSql           = "DELETE FROM `ims_vx_red_bag_quota_card` WHERE ( `id` = ? );"
	ImsVxRedBagQuotaCardUpdateSql           = "UPDATE `ims_vx_red_bag_quota_card` SET `account_id` = ?, `name` = ?, `avatar` = ?, `title` = ?, `details` = ?, `money` = ?, `prepare` = ?, `state` = ?, `state_msg` = ?, `created_at` = ?, `updated_at` = ?, `note` = ? WHERE ( `id` = ? );"
	ImsVxRedBagQuotaCardSelectSql           = "SELECT `id`, `account_id`, `name`, `avatar`, `title`, `details`, `money`, `prepare`, `state`, `state_msg`, `created_at`, `updated_at`, `note` FROM `ims_vx_red_bag_quota_card` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsWechatAttachmentInsertSql            = "INSERT INTO `ims_wechat_attachment` ( `uniacid`, `acid`, `uid`, `filename`, `attachment`, `media_id`, `width`, `height`, `type`, `model`, `tag`, `createtime`, `module_upload_dir`, `group_id` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsWechatAttachmentDeleteSql            = "DELETE FROM `ims_wechat_attachment` WHERE ( `id` = ? );"
	ImsWechatAttachmentUpdateSql            = "UPDATE `ims_wechat_attachment` SET `uniacid` = ?, `acid` = ?, `uid` = ?, `filename` = ?, `attachment` = ?, `media_id` = ?, `width` = ?, `height` = ?, `type` = ?, `model` = ?, `tag` = ?, `createtime` = ?, `module_upload_dir` = ?, `group_id` = ? WHERE ( `id` = ? );"
	ImsWechatAttachmentSelectSql            = "SELECT `id`, `uniacid`, `acid`, `uid`, `filename`, `attachment`, `media_id`, `width`, `height`, `type`, `model`, `tag`, `createtime`, `module_upload_dir`, `group_id` FROM `ims_wechat_attachment` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsWechatNewsInsertSql                  = "INSERT INTO `ims_wechat_news` ( `uniacid`, `attach_id`, `thumb_media_id`, `thumb_url`, `title`, `author`, `digest`, `content`, `content_source_url`, `show_cover_pic`, `url`, `displayorder`, `need_open_comment`, `only_fans_can_comment` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsWechatNewsDeleteSql                  = "DELETE FROM `ims_wechat_news` WHERE ( `id` = ? );"
	ImsWechatNewsUpdateSql                  = "UPDATE `ims_wechat_news` SET `uniacid` = ?, `attach_id` = ?, `thumb_media_id` = ?, `thumb_url` = ?, `title` = ?, `author` = ?, `digest` = ?, `content` = ?, `content_source_url` = ?, `show_cover_pic` = ?, `url` = ?, `displayorder` = ?, `need_open_comment` = ?, `only_fans_can_comment` = ? WHERE ( `id` = ? );"
	ImsWechatNewsSelectSql                  = "SELECT `id`, `uniacid`, `attach_id`, `thumb_media_id`, `thumb_url`, `title`, `author`, `digest`, `content`, `content_source_url`, `show_cover_pic`, `url`, `displayorder`, `need_open_comment`, `only_fans_can_comment` FROM `ims_wechat_news` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsWxappGeneralAnalysisInsertSql        = "INSERT INTO `ims_wxapp_general_analysis` ( `uniacid`, `session_cnt`, `visit_pv`, `visit_uv`, `visit_uv_new`, `type`, `stay_time_uv`, `stay_time_session`, `visit_depth`, `ref_date` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsWxappGeneralAnalysisDeleteSql        = "DELETE FROM `ims_wxapp_general_analysis` WHERE ( `id` = ? );"
	ImsWxappGeneralAnalysisUpdateSql        = "UPDATE `ims_wxapp_general_analysis` SET `uniacid` = ?, `session_cnt` = ?, `visit_pv` = ?, `visit_uv` = ?, `visit_uv_new` = ?, `type` = ?, `stay_time_uv` = ?, `stay_time_session` = ?, `visit_depth` = ?, `ref_date` = ? WHERE ( `id` = ? );"
	ImsWxappGeneralAnalysisSelectSql        = "SELECT `id`, `uniacid`, `session_cnt`, `visit_pv`, `visit_uv`, `visit_uv_new`, `type`, `stay_time_uv`, `stay_time_session`, `visit_depth`, `ref_date` FROM `ims_wxapp_general_analysis` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsWxappRegisterVersionInsertSql        = "INSERT INTO `ims_wxapp_register_version` ( `uniacid`, `version_id`, `auditid`, `version`, `description`, `status`, `reason`, `upload_time`, `audit_info`, `submit_info`, `developer` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsWxappRegisterVersionDeleteSql        = "DELETE FROM `ims_wxapp_register_version` WHERE ( `id` = ? );"
	ImsWxappRegisterVersionUpdateSql        = "UPDATE `ims_wxapp_register_version` SET `uniacid` = ?, `version_id` = ?, `auditid` = ?, `version` = ?, `description` = ?, `status` = ?, `reason` = ?, `upload_time` = ?, `audit_info` = ?, `submit_info` = ?, `developer` = ? WHERE ( `id` = ? );"
	ImsWxappRegisterVersionSelectSql        = "SELECT `id`, `uniacid`, `version_id`, `auditid`, `version`, `description`, `status`, `reason`, `upload_time`, `audit_info`, `submit_info`, `developer` FROM `ims_wxapp_register_version` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsWxappUndocodeauditLogInsertSql       = "INSERT INTO `ims_wxapp_undocodeaudit_log` ( `uniacid`, `version_id`, `auditid`, `revoke_time` ) VALUES ( ?, ?, ?, ? );"
	ImsWxappUndocodeauditLogDeleteSql       = "DELETE FROM `ims_wxapp_undocodeaudit_log` WHERE ( `id` = ? );"
	ImsWxappUndocodeauditLogUpdateSql       = "UPDATE `ims_wxapp_undocodeaudit_log` SET `uniacid` = ?, `version_id` = ?, `auditid` = ?, `revoke_time` = ? WHERE ( `id` = ? );"
	ImsWxappUndocodeauditLogSelectSql       = "SELECT `id`, `uniacid`, `version_id`, `auditid`, `revoke_time` FROM `ims_wxapp_undocodeaudit_log` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsWxappVersionsInsertSql               = "INSERT INTO `ims_wxapp_versions` ( `uniacid`, `multiid`, `version`, `description`, `modules`, `design_method`, `template`, `quickmenu`, `createtime`, `type`, `entry_id`, `appjson`, `default_appjson`, `use_default`, `last_modules`, `tominiprogram`, `upload_time` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsWxappVersionsDeleteSql               = "DELETE FROM `ims_wxapp_versions` WHERE ( `id` = ? );"
	ImsWxappVersionsUpdateSql               = "UPDATE `ims_wxapp_versions` SET `uniacid` = ?, `multiid` = ?, `version` = ?, `description` = ?, `modules` = ?, `design_method` = ?, `template` = ?, `quickmenu` = ?, `createtime` = ?, `type` = ?, `entry_id` = ?, `appjson` = ?, `default_appjson` = ?, `use_default` = ?, `last_modules` = ?, `tominiprogram` = ?, `upload_time` = ? WHERE ( `id` = ? );"
	ImsWxappVersionsSelectSql               = "SELECT `id`, `uniacid`, `multiid`, `version`, `description`, `modules`, `design_method`, `template`, `quickmenu`, `createtime`, `type`, `entry_id`, `appjson`, `default_appjson`, `use_default`, `last_modules`, `tominiprogram`, `upload_time` FROM `ims_wxapp_versions` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
	ImsWxcardReplyInsertSql                 = "INSERT INTO `ims_wxcard_reply` ( `rid`, `title`, `card_id`, `cid`, `brand_name`, `logo_url`, `success`, `error` ) VALUES ( ?, ?, ?, ?, ?, ?, ?, ? );"
	ImsWxcardReplyDeleteSql                 = "DELETE FROM `ims_wxcard_reply` WHERE ( `id` = ? );"
	ImsWxcardReplyUpdateSql                 = "UPDATE `ims_wxcard_reply` SET `rid` = ?, `title` = ?, `card_id` = ?, `cid` = ?, `brand_name` = ?, `logo_url` = ?, `success` = ?, `error` = ? WHERE ( `id` = ? );"
	ImsWxcardReplySelectSql                 = "SELECT `id`, `rid`, `title`, `card_id`, `cid`, `brand_name`, `logo_url`, `success`, `error` FROM `ims_wxcard_reply` WHERE ( `id` = ? ) ORDER BY `id` DESC LIMIT 0, 1;"
)

func ImsAccountScan(rows *sql.Rows) (s *ImsAccount, err error) {
	if rows.Next() {
		s = &ImsAccount{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Hash,
			&s.Type,
			&s.Isconnect,
			&s.Isdeleted,
			&s.Endtime,
			&s.SendAccountExpireStatus,
			&s.SendApiExpireStatus,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAccountScanAll(rows *sql.Rows) (ss []*ImsAccount, err error) {
	for rows.Next() {
		s := &ImsAccount{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Hash,
			&s.Type,
			&s.Isconnect,
			&s.Isdeleted,
			&s.Endtime,
			&s.SendAccountExpireStatus,
			&s.SendApiExpireStatus,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAccountAliappScan(rows *sql.Rows) (s *ImsAccountAliapp, err error) {
	if rows.Next() {
		s = &ImsAccountAliapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Level,
			&s.Name,
			&s.Description,
			&s.Key,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAccountAliappScanAll(rows *sql.Rows) (ss []*ImsAccountAliapp, err error) {
	for rows.Next() {
		s := &ImsAccountAliapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Level,
			&s.Name,
			&s.Description,
			&s.Key,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAccountBaiduappScan(rows *sql.Rows) (s *ImsAccountBaiduapp, err error) {
	if rows.Next() {
		s = &ImsAccountBaiduapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
			&s.Appid,
			&s.Key,
			&s.Secret,
			&s.Description,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAccountBaiduappScanAll(rows *sql.Rows) (ss []*ImsAccountBaiduapp, err error) {
	for rows.Next() {
		s := &ImsAccountBaiduapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
			&s.Appid,
			&s.Key,
			&s.Secret,
			&s.Description,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAccountPhoneappScan(rows *sql.Rows) (s *ImsAccountPhoneapp, err error) {
	if rows.Next() {
		s = &ImsAccountPhoneapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAccountPhoneappScanAll(rows *sql.Rows) (ss []*ImsAccountPhoneapp, err error) {
	for rows.Next() {
		s := &ImsAccountPhoneapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAccountToutiaoappScan(rows *sql.Rows) (s *ImsAccountToutiaoapp, err error) {
	if rows.Next() {
		s = &ImsAccountToutiaoapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
			&s.Appid,
			&s.Key,
			&s.Secret,
			&s.Description,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAccountToutiaoappScanAll(rows *sql.Rows) (ss []*ImsAccountToutiaoapp, err error) {
	for rows.Next() {
		s := &ImsAccountToutiaoapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
			&s.Appid,
			&s.Key,
			&s.Secret,
			&s.Description,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAccountWebappScan(rows *sql.Rows) (s *ImsAccountWebapp, err error) {
	if rows.Next() {
		s = &ImsAccountWebapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAccountWebappScanAll(rows *sql.Rows) (ss []*ImsAccountWebapp, err error) {
	for rows.Next() {
		s := &ImsAccountWebapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAccountWechatsScan(rows *sql.Rows) (s *ImsAccountWechats, err error) {
	if rows.Next() {
		s = &ImsAccountWechats{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Token,
			&s.Encodingaeskey,
			&s.Level,
			&s.Name,
			&s.Account,
			&s.Original,
			&s.Signature,
			&s.Country,
			&s.Province,
			&s.City,
			&s.Username,
			&s.Password,
			&s.Lastupdate,
			&s.Key,
			&s.Secret,
			&s.Styleid,
			&s.Subscribeurl,
			&s.AuthRefreshToken,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAccountWechatsScanAll(rows *sql.Rows) (ss []*ImsAccountWechats, err error) {
	for rows.Next() {
		s := &ImsAccountWechats{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Token,
			&s.Encodingaeskey,
			&s.Level,
			&s.Name,
			&s.Account,
			&s.Original,
			&s.Signature,
			&s.Country,
			&s.Province,
			&s.City,
			&s.Username,
			&s.Password,
			&s.Lastupdate,
			&s.Key,
			&s.Secret,
			&s.Styleid,
			&s.Subscribeurl,
			&s.AuthRefreshToken,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAccountWxappScan(rows *sql.Rows) (s *ImsAccountWxapp, err error) {
	if rows.Next() {
		s = &ImsAccountWxapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Token,
			&s.Encodingaeskey,
			&s.Level,
			&s.Account,
			&s.Original,
			&s.Key,
			&s.Secret,
			&s.Name,
			&s.Appdomain,
			&s.AuthRefreshToken,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAccountWxappScanAll(rows *sql.Rows) (ss []*ImsAccountWxapp, err error) {
	for rows.Next() {
		s := &ImsAccountWxapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Token,
			&s.Encodingaeskey,
			&s.Level,
			&s.Account,
			&s.Original,
			&s.Key,
			&s.Secret,
			&s.Name,
			&s.Appdomain,
			&s.AuthRefreshToken,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAccountXzappScan(rows *sql.Rows) (s *ImsAccountXzapp, err error) {
	if rows.Next() {
		s = &ImsAccountXzapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
			&s.Original,
			&s.Lastupdate,
			&s.Styleid,
			&s.Createtime,
			&s.Token,
			&s.Encodingaeskey,
			&s.XzappId,
			&s.Level,
			&s.Key,
			&s.Secret,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAccountXzappScanAll(rows *sql.Rows) (ss []*ImsAccountXzapp, err error) {
	for rows.Next() {
		s := &ImsAccountXzapp{}
		err = rows.Scan(
			&s.Acid,
			&s.Uniacid,
			&s.Name,
			&s.Original,
			&s.Lastupdate,
			&s.Styleid,
			&s.Createtime,
			&s.Token,
			&s.Encodingaeskey,
			&s.XzappId,
			&s.Level,
			&s.Key,
			&s.Secret,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsActivityClerksScan(rows *sql.Rows) (s *ImsActivityClerks, err error) {
	if rows.Next() {
		s = &ImsActivityClerks{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Storeid,
			&s.Name,
			&s.Password,
			&s.Mobile,
			&s.Openid,
			&s.Nickname,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsActivityClerksScanAll(rows *sql.Rows) (ss []*ImsActivityClerks, err error) {
	for rows.Next() {
		s := &ImsActivityClerks{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Storeid,
			&s.Name,
			&s.Password,
			&s.Mobile,
			&s.Openid,
			&s.Nickname,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsActivityClerkMenuScan(rows *sql.Rows) (s *ImsActivityClerkMenu, err error) {
	if rows.Next() {
		s = &ImsActivityClerkMenu{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Displayorder,
			&s.Pid,
			&s.GroupName,
			&s.Title,
			&s.Icon,
			&s.Url,
			&s.Type,
			&s.Permission,
			&s.System,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsActivityClerkMenuScanAll(rows *sql.Rows) (ss []*ImsActivityClerkMenu, err error) {
	for rows.Next() {
		s := &ImsActivityClerkMenu{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Displayorder,
			&s.Pid,
			&s.GroupName,
			&s.Title,
			&s.Icon,
			&s.Url,
			&s.Type,
			&s.Permission,
			&s.System,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAddressScan(rows *sql.Rows) (s *ImsAddress, err error) {
	if rows.Next() {
		s = &ImsAddress{}
		err = rows.Scan(
			&s.Id,
			&s.Pid,
			&s.Name,
			&s.Level,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAddressScanAll(rows *sql.Rows) (ss []*ImsAddress, err error) {
	for rows.Next() {
		s := &ImsAddress{}
		err = rows.Scan(
			&s.Id,
			&s.Pid,
			&s.Name,
			&s.Level,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsArticleCategoryScan(rows *sql.Rows) (s *ImsArticleCategory, err error) {
	if rows.Next() {
		s = &ImsArticleCategory{}
		err = rows.Scan(
			&s.Id,
			&s.Title,
			&s.Displayorder,
			&s.Type,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsArticleCategoryScanAll(rows *sql.Rows) (ss []*ImsArticleCategory, err error) {
	for rows.Next() {
		s := &ImsArticleCategory{}
		err = rows.Scan(
			&s.Id,
			&s.Title,
			&s.Displayorder,
			&s.Type,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsArticleCommentScan(rows *sql.Rows) (s *ImsArticleComment, err error) {
	if rows.Next() {
		s = &ImsArticleComment{}
		err = rows.Scan(
			&s.Id,
			&s.Articleid,
			&s.Parentid,
			&s.Uid,
			&s.Content,
			&s.IsLike,
			&s.IsReply,
			&s.LikeNum,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsArticleCommentScanAll(rows *sql.Rows) (ss []*ImsArticleComment, err error) {
	for rows.Next() {
		s := &ImsArticleComment{}
		err = rows.Scan(
			&s.Id,
			&s.Articleid,
			&s.Parentid,
			&s.Uid,
			&s.Content,
			&s.IsLike,
			&s.IsReply,
			&s.LikeNum,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsArticleNewsScan(rows *sql.Rows) (s *ImsArticleNews, err error) {
	if rows.Next() {
		s = &ImsArticleNews{}
		err = rows.Scan(
			&s.Id,
			&s.Cateid,
			&s.Title,
			&s.Content,
			&s.Thumb,
			&s.Source,
			&s.Author,
			&s.Displayorder,
			&s.IsDisplay,
			&s.IsShowHome,
			&s.Createtime,
			&s.Click,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsArticleNewsScanAll(rows *sql.Rows) (ss []*ImsArticleNews, err error) {
	for rows.Next() {
		s := &ImsArticleNews{}
		err = rows.Scan(
			&s.Id,
			&s.Cateid,
			&s.Title,
			&s.Content,
			&s.Thumb,
			&s.Source,
			&s.Author,
			&s.Displayorder,
			&s.IsDisplay,
			&s.IsShowHome,
			&s.Createtime,
			&s.Click,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsArticleNoticeScan(rows *sql.Rows) (s *ImsArticleNotice, err error) {
	if rows.Next() {
		s = &ImsArticleNotice{}
		err = rows.Scan(
			&s.Id,
			&s.Cateid,
			&s.Title,
			&s.Content,
			&s.Displayorder,
			&s.IsDisplay,
			&s.IsShowHome,
			&s.Createtime,
			&s.Click,
			&s.Style,
			&s.Group,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsArticleNoticeScanAll(rows *sql.Rows) (ss []*ImsArticleNotice, err error) {
	for rows.Next() {
		s := &ImsArticleNotice{}
		err = rows.Scan(
			&s.Id,
			&s.Cateid,
			&s.Title,
			&s.Content,
			&s.Displayorder,
			&s.IsDisplay,
			&s.IsShowHome,
			&s.Createtime,
			&s.Click,
			&s.Style,
			&s.Group,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsArticleUnreadNoticeScan(rows *sql.Rows) (s *ImsArticleUnreadNotice, err error) {
	if rows.Next() {
		s = &ImsArticleUnreadNotice{}
		err = rows.Scan(
			&s.Id,
			&s.NoticeId,
			&s.Uid,
			&s.IsNew,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsArticleUnreadNoticeScanAll(rows *sql.Rows) (ss []*ImsArticleUnreadNotice, err error) {
	for rows.Next() {
		s := &ImsArticleUnreadNotice{}
		err = rows.Scan(
			&s.Id,
			&s.NoticeId,
			&s.Uid,
			&s.IsNew,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsAttachmentGroupScan(rows *sql.Rows) (s *ImsAttachmentGroup, err error) {
	if rows.Next() {
		s = &ImsAttachmentGroup{}
		err = rows.Scan(
			&s.Id,
			&s.Pid,
			&s.Name,
			&s.Uniacid,
			&s.Uid,
			&s.Type,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsAttachmentGroupScanAll(rows *sql.Rows) (ss []*ImsAttachmentGroup, err error) {
	for rows.Next() {
		s := &ImsAttachmentGroup{}
		err = rows.Scan(
			&s.Id,
			&s.Pid,
			&s.Name,
			&s.Uniacid,
			&s.Uid,
			&s.Type,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsBasicReplyScan(rows *sql.Rows) (s *ImsBasicReply, err error) {
	if rows.Next() {
		s = &ImsBasicReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Content,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsBasicReplyScanAll(rows *sql.Rows) (ss []*ImsBasicReply, err error) {
	for rows.Next() {
		s := &ImsBasicReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Content,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsBusinessScan(rows *sql.Rows) (s *ImsBusiness, err error) {
	if rows.Next() {
		s = &ImsBusiness{}
		err = rows.Scan(
			&s.Id,
			&s.Weid,
			&s.Title,
			&s.Thumb,
			&s.Content,
			&s.Phone,
			&s.Qq,
			&s.Province,
			&s.City,
			&s.Dist,
			&s.Address,
			&s.Lng,
			&s.Lat,
			&s.Industry1,
			&s.Industry2,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsBusinessScanAll(rows *sql.Rows) (ss []*ImsBusiness, err error) {
	for rows.Next() {
		s := &ImsBusiness{}
		err = rows.Scan(
			&s.Id,
			&s.Weid,
			&s.Title,
			&s.Thumb,
			&s.Content,
			&s.Phone,
			&s.Qq,
			&s.Province,
			&s.City,
			&s.Dist,
			&s.Address,
			&s.Lng,
			&s.Lat,
			&s.Industry1,
			&s.Industry2,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreAttachmentScan(rows *sql.Rows) (s *ImsCoreAttachment, err error) {
	if rows.Next() {
		s = &ImsCoreAttachment{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Filename,
			&s.Attachment,
			&s.Type,
			&s.Createtime,
			&s.ModuleUploadDir,
			&s.GroupId,
			&s.Displayorder,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreAttachmentScanAll(rows *sql.Rows) (ss []*ImsCoreAttachment, err error) {
	for rows.Next() {
		s := &ImsCoreAttachment{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Filename,
			&s.Attachment,
			&s.Type,
			&s.Createtime,
			&s.ModuleUploadDir,
			&s.GroupId,
			&s.Displayorder,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreCacheScan(rows *sql.Rows) (s *ImsCoreCache, err error) {
	if rows.Next() {
		s = &ImsCoreCache{}
		err = rows.Scan(
			&s.Key,
			&s.Value,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreCacheScanAll(rows *sql.Rows) (ss []*ImsCoreCache, err error) {
	for rows.Next() {
		s := &ImsCoreCache{}
		err = rows.Scan(
			&s.Key,
			&s.Value,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreCronScan(rows *sql.Rows) (s *ImsCoreCron, err error) {
	if rows.Next() {
		s = &ImsCoreCron{}
		err = rows.Scan(
			&s.Id,
			&s.Cloudid,
			&s.Module,
			&s.Uniacid,
			&s.Type,
			&s.Name,
			&s.Filename,
			&s.Lastruntime,
			&s.Nextruntime,
			&s.Weekday,
			&s.Day,
			&s.Hour,
			&s.Minute,
			&s.Extra,
			&s.Status,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreCronScanAll(rows *sql.Rows) (ss []*ImsCoreCron, err error) {
	for rows.Next() {
		s := &ImsCoreCron{}
		err = rows.Scan(
			&s.Id,
			&s.Cloudid,
			&s.Module,
			&s.Uniacid,
			&s.Type,
			&s.Name,
			&s.Filename,
			&s.Lastruntime,
			&s.Nextruntime,
			&s.Weekday,
			&s.Day,
			&s.Hour,
			&s.Minute,
			&s.Extra,
			&s.Status,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreCronRecordScan(rows *sql.Rows) (s *ImsCoreCronRecord, err error) {
	if rows.Next() {
		s = &ImsCoreCronRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Module,
			&s.Type,
			&s.Tid,
			&s.Note,
			&s.Tag,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreCronRecordScanAll(rows *sql.Rows) (ss []*ImsCoreCronRecord, err error) {
	for rows.Next() {
		s := &ImsCoreCronRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Module,
			&s.Type,
			&s.Tid,
			&s.Note,
			&s.Tag,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreJobScan(rows *sql.Rows) (s *ImsCoreJob, err error) {
	if rows.Next() {
		s = &ImsCoreJob{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Uniacid,
			&s.Payload,
			&s.Status,
			&s.Title,
			&s.Handled,
			&s.Total,
			&s.Createtime,
			&s.Updatetime,
			&s.Endtime,
			&s.Uid,
			&s.Isdeleted,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreJobScanAll(rows *sql.Rows) (ss []*ImsCoreJob, err error) {
	for rows.Next() {
		s := &ImsCoreJob{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Uniacid,
			&s.Payload,
			&s.Status,
			&s.Title,
			&s.Handled,
			&s.Total,
			&s.Createtime,
			&s.Updatetime,
			&s.Endtime,
			&s.Uid,
			&s.Isdeleted,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreMenuScan(rows *sql.Rows) (s *ImsCoreMenu, err error) {
	if rows.Next() {
		s = &ImsCoreMenu{}
		err = rows.Scan(
			&s.Id,
			&s.Pid,
			&s.Title,
			&s.Name,
			&s.Url,
			&s.AppendTitle,
			&s.AppendUrl,
			&s.Displayorder,
			&s.Type,
			&s.IsDisplay,
			&s.IsSystem,
			&s.PermissionName,
			&s.GroupName,
			&s.Icon,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreMenuScanAll(rows *sql.Rows) (ss []*ImsCoreMenu, err error) {
	for rows.Next() {
		s := &ImsCoreMenu{}
		err = rows.Scan(
			&s.Id,
			&s.Pid,
			&s.Title,
			&s.Name,
			&s.Url,
			&s.AppendTitle,
			&s.AppendUrl,
			&s.Displayorder,
			&s.Type,
			&s.IsDisplay,
			&s.IsSystem,
			&s.PermissionName,
			&s.GroupName,
			&s.Icon,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreMenuShortcutScan(rows *sql.Rows) (s *ImsCoreMenuShortcut, err error) {
	if rows.Next() {
		s = &ImsCoreMenuShortcut{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Uniacid,
			&s.Modulename,
			&s.Displayorder,
			&s.Position,
			&s.Updatetime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreMenuShortcutScanAll(rows *sql.Rows) (ss []*ImsCoreMenuShortcut, err error) {
	for rows.Next() {
		s := &ImsCoreMenuShortcut{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Uniacid,
			&s.Modulename,
			&s.Displayorder,
			&s.Position,
			&s.Updatetime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCorePaylogScan(rows *sql.Rows) (s *ImsCorePaylog, err error) {
	if rows.Next() {
		s = &ImsCorePaylog{}
		err = rows.Scan(
			&s.Plid,
			&s.Type,
			&s.Uniacid,
			&s.Acid,
			&s.Openid,
			&s.Uniontid,
			&s.Tid,
			&s.Fee,
			&s.Status,
			&s.Module,
			&s.Tag,
			&s.IsUsecard,
			&s.CardType,
			&s.CardId,
			&s.CardFee,
			&s.EncryptCode,
			&s.IsWish,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCorePaylogScanAll(rows *sql.Rows) (ss []*ImsCorePaylog, err error) {
	for rows.Next() {
		s := &ImsCorePaylog{}
		err = rows.Scan(
			&s.Plid,
			&s.Type,
			&s.Uniacid,
			&s.Acid,
			&s.Openid,
			&s.Uniontid,
			&s.Tid,
			&s.Fee,
			&s.Status,
			&s.Module,
			&s.Tag,
			&s.IsUsecard,
			&s.CardType,
			&s.CardId,
			&s.CardFee,
			&s.EncryptCode,
			&s.IsWish,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCorePerformanceScan(rows *sql.Rows) (s *ImsCorePerformance, err error) {
	if rows.Next() {
		s = &ImsCorePerformance{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Runtime,
			&s.Runurl,
			&s.Runsql,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCorePerformanceScanAll(rows *sql.Rows) (ss []*ImsCorePerformance, err error) {
	for rows.Next() {
		s := &ImsCorePerformance{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Runtime,
			&s.Runurl,
			&s.Runsql,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreQueueScan(rows *sql.Rows) (s *ImsCoreQueue, err error) {
	if rows.Next() {
		s = &ImsCoreQueue{}
		err = rows.Scan(
			&s.Qid,
			&s.Uniacid,
			&s.Acid,
			&s.Message,
			&s.Params,
			&s.Keyword,
			&s.Response,
			&s.Module,
			&s.Type,
			&s.Dateline,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreQueueScanAll(rows *sql.Rows) (ss []*ImsCoreQueue, err error) {
	for rows.Next() {
		s := &ImsCoreQueue{}
		err = rows.Scan(
			&s.Qid,
			&s.Uniacid,
			&s.Acid,
			&s.Message,
			&s.Params,
			&s.Keyword,
			&s.Response,
			&s.Module,
			&s.Type,
			&s.Dateline,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreRefundlogScan(rows *sql.Rows) (s *ImsCoreRefundlog, err error) {
	if rows.Next() {
		s = &ImsCoreRefundlog{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.RefundUniontid,
			&s.Reason,
			&s.Uniontid,
			&s.Fee,
			&s.Status,
			&s.IsWish,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreRefundlogScanAll(rows *sql.Rows) (ss []*ImsCoreRefundlog, err error) {
	for rows.Next() {
		s := &ImsCoreRefundlog{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.RefundUniontid,
			&s.Reason,
			&s.Uniontid,
			&s.Fee,
			&s.Status,
			&s.IsWish,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreResourceScan(rows *sql.Rows) (s *ImsCoreResource, err error) {
	if rows.Next() {
		s = &ImsCoreResource{}
		err = rows.Scan(
			&s.Mid,
			&s.Uniacid,
			&s.MediaId,
			&s.Trunk,
			&s.Type,
			&s.Dateline,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreResourceScanAll(rows *sql.Rows) (ss []*ImsCoreResource, err error) {
	for rows.Next() {
		s := &ImsCoreResource{}
		err = rows.Scan(
			&s.Mid,
			&s.Uniacid,
			&s.MediaId,
			&s.Trunk,
			&s.Type,
			&s.Dateline,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreSendsmsLogScan(rows *sql.Rows) (s *ImsCoreSendsmsLog, err error) {
	if rows.Next() {
		s = &ImsCoreSendsmsLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Mobile,
			&s.Content,
			&s.Result,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreSendsmsLogScanAll(rows *sql.Rows) (ss []*ImsCoreSendsmsLog, err error) {
	for rows.Next() {
		s := &ImsCoreSendsmsLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Mobile,
			&s.Content,
			&s.Result,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreSessionsScan(rows *sql.Rows) (s *ImsCoreSessions, err error) {
	if rows.Next() {
		s = &ImsCoreSessions{}
		err = rows.Scan(
			&s.Sid,
			&s.Uniacid,
			&s.Openid,
			&s.Data,
			&s.Expiretime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreSessionsScanAll(rows *sql.Rows) (ss []*ImsCoreSessions, err error) {
	for rows.Next() {
		s := &ImsCoreSessions{}
		err = rows.Scan(
			&s.Sid,
			&s.Uniacid,
			&s.Openid,
			&s.Data,
			&s.Expiretime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoreSettingsScan(rows *sql.Rows) (s *ImsCoreSettings, err error) {
	if rows.Next() {
		s = &ImsCoreSettings{}
		err = rows.Scan(
			&s.Key,
			&s.Value,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoreSettingsScanAll(rows *sql.Rows) (ss []*ImsCoreSettings, err error) {
	for rows.Next() {
		s := &ImsCoreSettings{}
		err = rows.Scan(
			&s.Key,
			&s.Value,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCouponLocationScan(rows *sql.Rows) (s *ImsCouponLocation, err error) {
	if rows.Next() {
		s = &ImsCouponLocation{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Sid,
			&s.LocationId,
			&s.BusinessName,
			&s.BranchName,
			&s.Category,
			&s.Province,
			&s.City,
			&s.District,
			&s.Address,
			&s.Longitude,
			&s.Latitude,
			&s.Telephone,
			&s.PhotoList,
			&s.AvgPrice,
			&s.OpenTime,
			&s.Recommend,
			&s.Special,
			&s.Introduction,
			&s.OffsetType,
			&s.Status,
			&s.Message,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCouponLocationScanAll(rows *sql.Rows) (ss []*ImsCouponLocation, err error) {
	for rows.Next() {
		s := &ImsCouponLocation{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Sid,
			&s.LocationId,
			&s.BusinessName,
			&s.BranchName,
			&s.Category,
			&s.Province,
			&s.City,
			&s.District,
			&s.Address,
			&s.Longitude,
			&s.Latitude,
			&s.Telephone,
			&s.PhotoList,
			&s.AvgPrice,
			&s.OpenTime,
			&s.Recommend,
			&s.Special,
			&s.Introduction,
			&s.OffsetType,
			&s.Status,
			&s.Message,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCoverReplyScan(rows *sql.Rows) (s *ImsCoverReply, err error) {
	if rows.Next() {
		s = &ImsCoverReply{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Rid,
			&s.Module,
			&s.Do,
			&s.Title,
			&s.Description,
			&s.Thumb,
			&s.Url,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCoverReplyScanAll(rows *sql.Rows) (ss []*ImsCoverReply, err error) {
	for rows.Next() {
		s := &ImsCoverReply{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Rid,
			&s.Module,
			&s.Do,
			&s.Title,
			&s.Description,
			&s.Thumb,
			&s.Url,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsCustomReplyScan(rows *sql.Rows) (s *ImsCustomReply, err error) {
	if rows.Next() {
		s = &ImsCustomReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Start1,
			&s.End1,
			&s.Start2,
			&s.End2,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsCustomReplyScanAll(rows *sql.Rows) (ss []*ImsCustomReply, err error) {
	for rows.Next() {
		s := &ImsCustomReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Start1,
			&s.End1,
			&s.Start2,
			&s.End2,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsImagesReplyScan(rows *sql.Rows) (s *ImsImagesReply, err error) {
	if rows.Next() {
		s = &ImsImagesReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.Description,
			&s.Mediaid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsImagesReplyScanAll(rows *sql.Rows) (ss []*ImsImagesReply, err error) {
	for rows.Next() {
		s := &ImsImagesReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.Description,
			&s.Mediaid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcCashRecordScan(rows *sql.Rows) (s *ImsMcCashRecord, err error) {
	if rows.Next() {
		s = &ImsMcCashRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.ClerkId,
			&s.StoreId,
			&s.ClerkType,
			&s.Fee,
			&s.FinalFee,
			&s.Credit1,
			&s.Credit1Fee,
			&s.Credit2,
			&s.Cash,
			&s.ReturnCash,
			&s.FinalCash,
			&s.Remark,
			&s.Createtime,
			&s.TradeType,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcCashRecordScanAll(rows *sql.Rows) (ss []*ImsMcCashRecord, err error) {
	for rows.Next() {
		s := &ImsMcCashRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.ClerkId,
			&s.StoreId,
			&s.ClerkType,
			&s.Fee,
			&s.FinalFee,
			&s.Credit1,
			&s.Credit1Fee,
			&s.Credit2,
			&s.Cash,
			&s.ReturnCash,
			&s.FinalCash,
			&s.Remark,
			&s.Createtime,
			&s.TradeType,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcChatsRecordScan(rows *sql.Rows) (s *ImsMcChatsRecord, err error) {
	if rows.Next() {
		s = &ImsMcChatsRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Flag,
			&s.Openid,
			&s.Msgtype,
			&s.Content,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcChatsRecordScanAll(rows *sql.Rows) (ss []*ImsMcChatsRecord, err error) {
	for rows.Next() {
		s := &ImsMcChatsRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Flag,
			&s.Openid,
			&s.Msgtype,
			&s.Content,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcCreditsRechargeScan(rows *sql.Rows) (s *ImsMcCreditsRecharge, err error) {
	if rows.Next() {
		s = &ImsMcCreditsRecharge{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Openid,
			&s.Tid,
			&s.Transid,
			&s.Fee,
			&s.Type,
			&s.Tag,
			&s.Status,
			&s.Createtime,
			&s.Backtype,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcCreditsRechargeScanAll(rows *sql.Rows) (ss []*ImsMcCreditsRecharge, err error) {
	for rows.Next() {
		s := &ImsMcCreditsRecharge{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Openid,
			&s.Tid,
			&s.Transid,
			&s.Fee,
			&s.Type,
			&s.Tag,
			&s.Status,
			&s.Createtime,
			&s.Backtype,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcCreditsRecordScan(rows *sql.Rows) (s *ImsMcCreditsRecord, err error) {
	if rows.Next() {
		s = &ImsMcCreditsRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Uniacid,
			&s.Credittype,
			&s.Num,
			&s.Operator,
			&s.Module,
			&s.ClerkId,
			&s.StoreId,
			&s.ClerkType,
			&s.Createtime,
			&s.Remark,
			&s.RealUniacid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcCreditsRecordScanAll(rows *sql.Rows) (ss []*ImsMcCreditsRecord, err error) {
	for rows.Next() {
		s := &ImsMcCreditsRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Uniacid,
			&s.Credittype,
			&s.Num,
			&s.Operator,
			&s.Module,
			&s.ClerkId,
			&s.StoreId,
			&s.ClerkType,
			&s.Createtime,
			&s.Remark,
			&s.RealUniacid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcFansGroupsScan(rows *sql.Rows) (s *ImsMcFansGroups, err error) {
	if rows.Next() {
		s = &ImsMcFansGroups{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Groups,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcFansGroupsScanAll(rows *sql.Rows) (ss []*ImsMcFansGroups, err error) {
	for rows.Next() {
		s := &ImsMcFansGroups{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Groups,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcFansTagScan(rows *sql.Rows) (s *ImsMcFansTag, err error) {
	if rows.Next() {
		s = &ImsMcFansTag{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Fanid,
			&s.Openid,
			&s.Subscribe,
			&s.Nickname,
			&s.Sex,
			&s.Language,
			&s.City,
			&s.Province,
			&s.Country,
			&s.Headimgurl,
			&s.SubscribeTime,
			&s.Unionid,
			&s.Remark,
			&s.Groupid,
			&s.TagidList,
			&s.SubscribeScene,
			&s.QrSceneStr,
			&s.QrScene,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcFansTagScanAll(rows *sql.Rows) (ss []*ImsMcFansTag, err error) {
	for rows.Next() {
		s := &ImsMcFansTag{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Fanid,
			&s.Openid,
			&s.Subscribe,
			&s.Nickname,
			&s.Sex,
			&s.Language,
			&s.City,
			&s.Province,
			&s.Country,
			&s.Headimgurl,
			&s.SubscribeTime,
			&s.Unionid,
			&s.Remark,
			&s.Groupid,
			&s.TagidList,
			&s.SubscribeScene,
			&s.QrSceneStr,
			&s.QrScene,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcFansTagMappingScan(rows *sql.Rows) (s *ImsMcFansTagMapping, err error) {
	if rows.Next() {
		s = &ImsMcFansTagMapping{}
		err = rows.Scan(
			&s.Id,
			&s.Fanid,
			&s.Tagid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcFansTagMappingScanAll(rows *sql.Rows) (ss []*ImsMcFansTagMapping, err error) {
	for rows.Next() {
		s := &ImsMcFansTagMapping{}
		err = rows.Scan(
			&s.Id,
			&s.Fanid,
			&s.Tagid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcGroupsScan(rows *sql.Rows) (s *ImsMcGroups, err error) {
	if rows.Next() {
		s = &ImsMcGroups{}
		err = rows.Scan(
			&s.Groupid,
			&s.Uniacid,
			&s.Title,
			&s.Credit,
			&s.Isdefault,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcGroupsScanAll(rows *sql.Rows) (ss []*ImsMcGroups, err error) {
	for rows.Next() {
		s := &ImsMcGroups{}
		err = rows.Scan(
			&s.Groupid,
			&s.Uniacid,
			&s.Title,
			&s.Credit,
			&s.Isdefault,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcHandselScan(rows *sql.Rows) (s *ImsMcHandsel, err error) {
	if rows.Next() {
		s = &ImsMcHandsel{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Touid,
			&s.Fromuid,
			&s.Module,
			&s.Sign,
			&s.Action,
			&s.CreditValue,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcHandselScanAll(rows *sql.Rows) (ss []*ImsMcHandsel, err error) {
	for rows.Next() {
		s := &ImsMcHandsel{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Touid,
			&s.Fromuid,
			&s.Module,
			&s.Sign,
			&s.Action,
			&s.CreditValue,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcMappingFansScan(rows *sql.Rows) (s *ImsMcMappingFans, err error) {
	if rows.Next() {
		s = &ImsMcMappingFans{}
		err = rows.Scan(
			&s.Fanid,
			&s.Acid,
			&s.Uniacid,
			&s.Uid,
			&s.Openid,
			&s.Nickname,
			&s.Groupid,
			&s.Salt,
			&s.Follow,
			&s.Followtime,
			&s.Unfollowtime,
			&s.Tag,
			&s.Updatetime,
			&s.Unionid,
			&s.UserFrom,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcMappingFansScanAll(rows *sql.Rows) (ss []*ImsMcMappingFans, err error) {
	for rows.Next() {
		s := &ImsMcMappingFans{}
		err = rows.Scan(
			&s.Fanid,
			&s.Acid,
			&s.Uniacid,
			&s.Uid,
			&s.Openid,
			&s.Nickname,
			&s.Groupid,
			&s.Salt,
			&s.Follow,
			&s.Followtime,
			&s.Unfollowtime,
			&s.Tag,
			&s.Updatetime,
			&s.Unionid,
			&s.UserFrom,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcMassRecordScan(rows *sql.Rows) (s *ImsMcMassRecord, err error) {
	if rows.Next() {
		s = &ImsMcMassRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Groupname,
			&s.Fansnum,
			&s.Msgtype,
			&s.Content,
			&s.Group,
			&s.AttachId,
			&s.MediaId,
			&s.Type,
			&s.Status,
			&s.CronId,
			&s.Sendtime,
			&s.Finalsendtime,
			&s.Createtime,
			&s.MsgId,
			&s.MsgDataId,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcMassRecordScanAll(rows *sql.Rows) (ss []*ImsMcMassRecord, err error) {
	for rows.Next() {
		s := &ImsMcMassRecord{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Groupname,
			&s.Fansnum,
			&s.Msgtype,
			&s.Content,
			&s.Group,
			&s.AttachId,
			&s.MediaId,
			&s.Type,
			&s.Status,
			&s.CronId,
			&s.Sendtime,
			&s.Finalsendtime,
			&s.Createtime,
			&s.MsgId,
			&s.MsgDataId,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcMembersScan(rows *sql.Rows) (s *ImsMcMembers, err error) {
	if rows.Next() {
		s = &ImsMcMembers{}
		err = rows.Scan(
			&s.Uid,
			&s.Uniacid,
			&s.Mobile,
			&s.Email,
			&s.Password,
			&s.Salt,
			&s.Groupid,
			&s.Credit1,
			&s.Credit2,
			&s.Credit3,
			&s.Credit4,
			&s.Credit5,
			&s.Credit6,
			&s.Createtime,
			&s.Realname,
			&s.Nickname,
			&s.Avatar,
			&s.Qq,
			&s.Vip,
			&s.Gender,
			&s.Birthyear,
			&s.Birthmonth,
			&s.Birthday,
			&s.Constellation,
			&s.Zodiac,
			&s.Telephone,
			&s.Idcard,
			&s.Studentid,
			&s.Grade,
			&s.Address,
			&s.Zipcode,
			&s.Nationality,
			&s.Resideprovince,
			&s.Residecity,
			&s.Residedist,
			&s.Graduateschool,
			&s.Company,
			&s.Education,
			&s.Occupation,
			&s.Position,
			&s.Revenue,
			&s.Affectivestatus,
			&s.Lookingfor,
			&s.Bloodtype,
			&s.Height,
			&s.Weight,
			&s.Alipay,
			&s.Msn,
			&s.Taobao,
			&s.Site,
			&s.Bio,
			&s.Interest,
			&s.PayPassword,
			&s.UserFrom,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcMembersScanAll(rows *sql.Rows) (ss []*ImsMcMembers, err error) {
	for rows.Next() {
		s := &ImsMcMembers{}
		err = rows.Scan(
			&s.Uid,
			&s.Uniacid,
			&s.Mobile,
			&s.Email,
			&s.Password,
			&s.Salt,
			&s.Groupid,
			&s.Credit1,
			&s.Credit2,
			&s.Credit3,
			&s.Credit4,
			&s.Credit5,
			&s.Credit6,
			&s.Createtime,
			&s.Realname,
			&s.Nickname,
			&s.Avatar,
			&s.Qq,
			&s.Vip,
			&s.Gender,
			&s.Birthyear,
			&s.Birthmonth,
			&s.Birthday,
			&s.Constellation,
			&s.Zodiac,
			&s.Telephone,
			&s.Idcard,
			&s.Studentid,
			&s.Grade,
			&s.Address,
			&s.Zipcode,
			&s.Nationality,
			&s.Resideprovince,
			&s.Residecity,
			&s.Residedist,
			&s.Graduateschool,
			&s.Company,
			&s.Education,
			&s.Occupation,
			&s.Position,
			&s.Revenue,
			&s.Affectivestatus,
			&s.Lookingfor,
			&s.Bloodtype,
			&s.Height,
			&s.Weight,
			&s.Alipay,
			&s.Msn,
			&s.Taobao,
			&s.Site,
			&s.Bio,
			&s.Interest,
			&s.PayPassword,
			&s.UserFrom,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcMemberAddressScan(rows *sql.Rows) (s *ImsMcMemberAddress, err error) {
	if rows.Next() {
		s = &ImsMcMemberAddress{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Username,
			&s.Mobile,
			&s.Zipcode,
			&s.Province,
			&s.City,
			&s.District,
			&s.Address,
			&s.Isdefault,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcMemberAddressScanAll(rows *sql.Rows) (ss []*ImsMcMemberAddress, err error) {
	for rows.Next() {
		s := &ImsMcMemberAddress{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Username,
			&s.Mobile,
			&s.Zipcode,
			&s.Province,
			&s.City,
			&s.District,
			&s.Address,
			&s.Isdefault,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcMemberFieldsScan(rows *sql.Rows) (s *ImsMcMemberFields, err error) {
	if rows.Next() {
		s = &ImsMcMemberFields{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Fieldid,
			&s.Title,
			&s.Available,
			&s.Displayorder,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcMemberFieldsScanAll(rows *sql.Rows) (ss []*ImsMcMemberFields, err error) {
	for rows.Next() {
		s := &ImsMcMemberFields{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Fieldid,
			&s.Title,
			&s.Available,
			&s.Displayorder,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcMemberPropertyScan(rows *sql.Rows) (s *ImsMcMemberProperty, err error) {
	if rows.Next() {
		s = &ImsMcMemberProperty{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Property,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcMemberPropertyScanAll(rows *sql.Rows) (ss []*ImsMcMemberProperty, err error) {
	for rows.Next() {
		s := &ImsMcMemberProperty{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Property,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMcOauthFansScan(rows *sql.Rows) (s *ImsMcOauthFans, err error) {
	if rows.Next() {
		s = &ImsMcOauthFans{}
		err = rows.Scan(
			&s.Id,
			&s.OauthOpenid,
			&s.Acid,
			&s.Uid,
			&s.Openid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMcOauthFansScanAll(rows *sql.Rows) (ss []*ImsMcOauthFans, err error) {
	for rows.Next() {
		s := &ImsMcOauthFans{}
		err = rows.Scan(
			&s.Id,
			&s.OauthOpenid,
			&s.Acid,
			&s.Uid,
			&s.Openid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMenuEventScan(rows *sql.Rows) (s *ImsMenuEvent, err error) {
	if rows.Next() {
		s = &ImsMenuEvent{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Keyword,
			&s.Type,
			&s.Picmd5,
			&s.Openid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMenuEventScanAll(rows *sql.Rows) (ss []*ImsMenuEvent, err error) {
	for rows.Next() {
		s := &ImsMenuEvent{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Keyword,
			&s.Type,
			&s.Picmd5,
			&s.Openid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMessageNoticeLogScan(rows *sql.Rows) (s *ImsMessageNoticeLog, err error) {
	if rows.Next() {
		s = &ImsMessageNoticeLog{}
		err = rows.Scan(
			&s.Id,
			&s.Message,
			&s.IsRead,
			&s.Uid,
			&s.Sign,
			&s.Type,
			&s.Status,
			&s.CreateTime,
			&s.EndTime,
			&s.Url,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMessageNoticeLogScanAll(rows *sql.Rows) (ss []*ImsMessageNoticeLog, err error) {
	for rows.Next() {
		s := &ImsMessageNoticeLog{}
		err = rows.Scan(
			&s.Id,
			&s.Message,
			&s.IsRead,
			&s.Uid,
			&s.Sign,
			&s.Type,
			&s.Status,
			&s.CreateTime,
			&s.EndTime,
			&s.Url,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMobilenumberScan(rows *sql.Rows) (s *ImsMobilenumber, err error) {
	if rows.Next() {
		s = &ImsMobilenumber{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Enabled,
			&s.Dateline,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMobilenumberScanAll(rows *sql.Rows) (ss []*ImsMobilenumber, err error) {
	for rows.Next() {
		s := &ImsMobilenumber{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Enabled,
			&s.Dateline,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsModulesScan(rows *sql.Rows) (s *ImsModules, err error) {
	if rows.Next() {
		s = &ImsModules{}
		err = rows.Scan(
			&s.Mid,
			&s.Name,
			&s.ApplicationType,
			&s.Type,
			&s.Title,
			&s.Version,
			&s.Ability,
			&s.Description,
			&s.Author,
			&s.Url,
			&s.Settings,
			&s.Subscribes,
			&s.Handles,
			&s.Isrulefields,
			&s.Issystem,
			&s.Target,
			&s.Iscard,
			&s.Permissions,
			&s.TitleInitial,
			&s.WxappSupport,
			&s.WelcomeSupport,
			&s.OauthType,
			&s.WebappSupport,
			&s.PhoneappSupport,
			&s.AccountSupport,
			&s.XzappSupport,
			&s.AliappSupport,
			&s.Logo,
			&s.BaiduappSupport,
			&s.ToutiaoappSupport,
			&s.From,
			&s.CloudRecord,
			&s.Sections,
			&s.Label,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsModulesScanAll(rows *sql.Rows) (ss []*ImsModules, err error) {
	for rows.Next() {
		s := &ImsModules{}
		err = rows.Scan(
			&s.Mid,
			&s.Name,
			&s.ApplicationType,
			&s.Type,
			&s.Title,
			&s.Version,
			&s.Ability,
			&s.Description,
			&s.Author,
			&s.Url,
			&s.Settings,
			&s.Subscribes,
			&s.Handles,
			&s.Isrulefields,
			&s.Issystem,
			&s.Target,
			&s.Iscard,
			&s.Permissions,
			&s.TitleInitial,
			&s.WxappSupport,
			&s.WelcomeSupport,
			&s.OauthType,
			&s.WebappSupport,
			&s.PhoneappSupport,
			&s.AccountSupport,
			&s.XzappSupport,
			&s.AliappSupport,
			&s.Logo,
			&s.BaiduappSupport,
			&s.ToutiaoappSupport,
			&s.From,
			&s.CloudRecord,
			&s.Sections,
			&s.Label,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsModulesBindingsScan(rows *sql.Rows) (s *ImsModulesBindings, err error) {
	if rows.Next() {
		s = &ImsModulesBindings{}
		err = rows.Scan(
			&s.Eid,
			&s.Module,
			&s.Entry,
			&s.Call,
			&s.Title,
			&s.Do,
			&s.State,
			&s.Direct,
			&s.Url,
			&s.Icon,
			&s.Displayorder,
			&s.Multilevel,
			&s.Parent,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsModulesBindingsScanAll(rows *sql.Rows) (ss []*ImsModulesBindings, err error) {
	for rows.Next() {
		s := &ImsModulesBindings{}
		err = rows.Scan(
			&s.Eid,
			&s.Module,
			&s.Entry,
			&s.Call,
			&s.Title,
			&s.Do,
			&s.State,
			&s.Direct,
			&s.Url,
			&s.Icon,
			&s.Displayorder,
			&s.Multilevel,
			&s.Parent,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsModulesCloudScan(rows *sql.Rows) (s *ImsModulesCloud, err error) {
	if rows.Next() {
		s = &ImsModulesCloud{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.ApplicationType,
			&s.Title,
			&s.TitleInitial,
			&s.Logo,
			&s.Version,
			&s.InstallStatus,
			&s.AccountSupport,
			&s.WxappSupport,
			&s.WebappSupport,
			&s.PhoneappSupport,
			&s.WelcomeSupport,
			&s.MainModuleName,
			&s.MainModuleLogo,
			&s.HasNewVersion,
			&s.HasNewBranch,
			&s.IsBan,
			&s.Lastupdatetime,
			&s.XzappSupport,
			&s.CloudId,
			&s.AliappSupport,
			&s.BaiduappSupport,
			&s.ToutiaoappSupport,
			&s.Buytime,
			&s.ModuleStatus,
			&s.Label,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsModulesCloudScanAll(rows *sql.Rows) (ss []*ImsModulesCloud, err error) {
	for rows.Next() {
		s := &ImsModulesCloud{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.ApplicationType,
			&s.Title,
			&s.TitleInitial,
			&s.Logo,
			&s.Version,
			&s.InstallStatus,
			&s.AccountSupport,
			&s.WxappSupport,
			&s.WebappSupport,
			&s.PhoneappSupport,
			&s.WelcomeSupport,
			&s.MainModuleName,
			&s.MainModuleLogo,
			&s.HasNewVersion,
			&s.HasNewBranch,
			&s.IsBan,
			&s.Lastupdatetime,
			&s.XzappSupport,
			&s.CloudId,
			&s.AliappSupport,
			&s.BaiduappSupport,
			&s.ToutiaoappSupport,
			&s.Buytime,
			&s.ModuleStatus,
			&s.Label,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsModulesIgnoreScan(rows *sql.Rows) (s *ImsModulesIgnore, err error) {
	if rows.Next() {
		s = &ImsModulesIgnore{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.Version,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsModulesIgnoreScanAll(rows *sql.Rows) (ss []*ImsModulesIgnore, err error) {
	for rows.Next() {
		s := &ImsModulesIgnore{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.Version,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsModulesPluginScan(rows *sql.Rows) (s *ImsModulesPlugin, err error) {
	if rows.Next() {
		s = &ImsModulesPlugin{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.MainModule,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsModulesPluginScanAll(rows *sql.Rows) (ss []*ImsModulesPlugin, err error) {
	for rows.Next() {
		s := &ImsModulesPlugin{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.MainModule,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsModulesPluginRankScan(rows *sql.Rows) (s *ImsModulesPluginRank, err error) {
	if rows.Next() {
		s = &ImsModulesPluginRank{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Rank,
			&s.PluginName,
			&s.MainModuleName,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsModulesPluginRankScanAll(rows *sql.Rows) (ss []*ImsModulesPluginRank, err error) {
	for rows.Next() {
		s := &ImsModulesPluginRank{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Rank,
			&s.PluginName,
			&s.MainModuleName,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsModulesRankScan(rows *sql.Rows) (s *ImsModulesRank, err error) {
	if rows.Next() {
		s = &ImsModulesRank{}
		err = rows.Scan(
			&s.Id,
			&s.ModuleName,
			&s.Uid,
			&s.Rank,
			&s.Uniacid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsModulesRankScanAll(rows *sql.Rows) (ss []*ImsModulesRank, err error) {
	for rows.Next() {
		s := &ImsModulesRank{}
		err = rows.Scan(
			&s.Id,
			&s.ModuleName,
			&s.Uid,
			&s.Rank,
			&s.Uniacid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsModulesRecycleScan(rows *sql.Rows) (s *ImsModulesRecycle, err error) {
	if rows.Next() {
		s = &ImsModulesRecycle{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.Type,
			&s.AccountSupport,
			&s.WxappSupport,
			&s.WelcomeSupport,
			&s.WebappSupport,
			&s.PhoneappSupport,
			&s.XzappSupport,
			&s.AliappSupport,
			&s.BaiduappSupport,
			&s.ToutiaoappSupport,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsModulesRecycleScanAll(rows *sql.Rows) (ss []*ImsModulesRecycle, err error) {
	for rows.Next() {
		s := &ImsModulesRecycle{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.Type,
			&s.AccountSupport,
			&s.WxappSupport,
			&s.WelcomeSupport,
			&s.WebappSupport,
			&s.PhoneappSupport,
			&s.XzappSupport,
			&s.AliappSupport,
			&s.BaiduappSupport,
			&s.ToutiaoappSupport,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsMusicReplyScan(rows *sql.Rows) (s *ImsMusicReply, err error) {
	if rows.Next() {
		s = &ImsMusicReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.Description,
			&s.Url,
			&s.Hqurl,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsMusicReplyScanAll(rows *sql.Rows) (ss []*ImsMusicReply, err error) {
	for rows.Next() {
		s := &ImsMusicReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.Description,
			&s.Url,
			&s.Hqurl,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsNewsReplyScan(rows *sql.Rows) (s *ImsNewsReply, err error) {
	if rows.Next() {
		s = &ImsNewsReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.ParentId,
			&s.Title,
			&s.Author,
			&s.Description,
			&s.Thumb,
			&s.Content,
			&s.Url,
			&s.Displayorder,
			&s.Incontent,
			&s.Createtime,
			&s.MediaId,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsNewsReplyScanAll(rows *sql.Rows) (ss []*ImsNewsReply, err error) {
	for rows.Next() {
		s := &ImsNewsReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.ParentId,
			&s.Title,
			&s.Author,
			&s.Description,
			&s.Thumb,
			&s.Content,
			&s.Url,
			&s.Displayorder,
			&s.Incontent,
			&s.Createtime,
			&s.MediaId,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsPhoneappVersionsScan(rows *sql.Rows) (s *ImsPhoneappVersions, err error) {
	if rows.Next() {
		s = &ImsPhoneappVersions{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Version,
			&s.Description,
			&s.Modules,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsPhoneappVersionsScanAll(rows *sql.Rows) (ss []*ImsPhoneappVersions, err error) {
	for rows.Next() {
		s := &ImsPhoneappVersions{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Version,
			&s.Description,
			&s.Modules,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsProfileFieldsScan(rows *sql.Rows) (s *ImsProfileFields, err error) {
	if rows.Next() {
		s = &ImsProfileFields{}
		err = rows.Scan(
			&s.Id,
			&s.Field,
			&s.Available,
			&s.Title,
			&s.Description,
			&s.Displayorder,
			&s.Required,
			&s.Unchangeable,
			&s.Showinregister,
			&s.FieldLength,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsProfileFieldsScanAll(rows *sql.Rows) (ss []*ImsProfileFields, err error) {
	for rows.Next() {
		s := &ImsProfileFields{}
		err = rows.Scan(
			&s.Id,
			&s.Field,
			&s.Available,
			&s.Title,
			&s.Description,
			&s.Displayorder,
			&s.Required,
			&s.Unchangeable,
			&s.Showinregister,
			&s.FieldLength,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsQrcodeScan(rows *sql.Rows) (s *ImsQrcode, err error) {
	if rows.Next() {
		s = &ImsQrcode{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Type,
			&s.Extra,
			&s.Qrcid,
			&s.SceneStr,
			&s.Name,
			&s.Keyword,
			&s.Model,
			&s.Ticket,
			&s.Url,
			&s.Expire,
			&s.Subnum,
			&s.Createtime,
			&s.Status,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsQrcodeScanAll(rows *sql.Rows) (ss []*ImsQrcode, err error) {
	for rows.Next() {
		s := &ImsQrcode{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Type,
			&s.Extra,
			&s.Qrcid,
			&s.SceneStr,
			&s.Name,
			&s.Keyword,
			&s.Model,
			&s.Ticket,
			&s.Url,
			&s.Expire,
			&s.Subnum,
			&s.Createtime,
			&s.Status,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsQrcodeStatScan(rows *sql.Rows) (s *ImsQrcodeStat, err error) {
	if rows.Next() {
		s = &ImsQrcodeStat{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Qid,
			&s.Openid,
			&s.Type,
			&s.Qrcid,
			&s.SceneStr,
			&s.Name,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsQrcodeStatScanAll(rows *sql.Rows) (ss []*ImsQrcodeStat, err error) {
	for rows.Next() {
		s := &ImsQrcodeStat{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Qid,
			&s.Openid,
			&s.Type,
			&s.Qrcid,
			&s.SceneStr,
			&s.Name,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsRuleScan(rows *sql.Rows) (s *ImsRule, err error) {
	if rows.Next() {
		s = &ImsRule{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Name,
			&s.Module,
			&s.Displayorder,
			&s.Status,
			&s.Containtype,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsRuleScanAll(rows *sql.Rows) (ss []*ImsRule, err error) {
	for rows.Next() {
		s := &ImsRule{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Name,
			&s.Module,
			&s.Displayorder,
			&s.Status,
			&s.Containtype,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsRuleKeywordScan(rows *sql.Rows) (s *ImsRuleKeyword, err error) {
	if rows.Next() {
		s = &ImsRuleKeyword{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Uniacid,
			&s.Module,
			&s.Content,
			&s.Type,
			&s.Displayorder,
			&s.Status,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsRuleKeywordScanAll(rows *sql.Rows) (ss []*ImsRuleKeyword, err error) {
	for rows.Next() {
		s := &ImsRuleKeyword{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Uniacid,
			&s.Module,
			&s.Content,
			&s.Type,
			&s.Displayorder,
			&s.Status,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteArticleScan(rows *sql.Rows) (s *ImsSiteArticle, err error) {
	if rows.Next() {
		s = &ImsSiteArticle{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Rid,
			&s.Kid,
			&s.Iscommend,
			&s.Ishot,
			&s.Pcate,
			&s.Ccate,
			&s.Template,
			&s.Title,
			&s.Description,
			&s.Content,
			&s.Thumb,
			&s.Incontent,
			&s.Source,
			&s.Author,
			&s.Displayorder,
			&s.Linkurl,
			&s.Createtime,
			&s.Edittime,
			&s.Click,
			&s.Type,
			&s.Credit,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteArticleScanAll(rows *sql.Rows) (ss []*ImsSiteArticle, err error) {
	for rows.Next() {
		s := &ImsSiteArticle{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Rid,
			&s.Kid,
			&s.Iscommend,
			&s.Ishot,
			&s.Pcate,
			&s.Ccate,
			&s.Template,
			&s.Title,
			&s.Description,
			&s.Content,
			&s.Thumb,
			&s.Incontent,
			&s.Source,
			&s.Author,
			&s.Displayorder,
			&s.Linkurl,
			&s.Createtime,
			&s.Edittime,
			&s.Click,
			&s.Type,
			&s.Credit,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteArticleCommentScan(rows *sql.Rows) (s *ImsSiteArticleComment, err error) {
	if rows.Next() {
		s = &ImsSiteArticleComment{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Articleid,
			&s.Parentid,
			&s.Uid,
			&s.Openid,
			&s.Content,
			&s.IsRead,
			&s.Iscomment,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteArticleCommentScanAll(rows *sql.Rows) (ss []*ImsSiteArticleComment, err error) {
	for rows.Next() {
		s := &ImsSiteArticleComment{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Articleid,
			&s.Parentid,
			&s.Uid,
			&s.Openid,
			&s.Content,
			&s.IsRead,
			&s.Iscomment,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteCategoryScan(rows *sql.Rows) (s *ImsSiteCategory, err error) {
	if rows.Next() {
		s = &ImsSiteCategory{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Nid,
			&s.Name,
			&s.Parentid,
			&s.Displayorder,
			&s.Enabled,
			&s.Icon,
			&s.Description,
			&s.Styleid,
			&s.Linkurl,
			&s.Ishomepage,
			&s.Icontype,
			&s.Css,
			&s.Multiid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteCategoryScanAll(rows *sql.Rows) (ss []*ImsSiteCategory, err error) {
	for rows.Next() {
		s := &ImsSiteCategory{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Nid,
			&s.Name,
			&s.Parentid,
			&s.Displayorder,
			&s.Enabled,
			&s.Icon,
			&s.Description,
			&s.Styleid,
			&s.Linkurl,
			&s.Ishomepage,
			&s.Icontype,
			&s.Css,
			&s.Multiid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteMultiScan(rows *sql.Rows) (s *ImsSiteMulti, err error) {
	if rows.Next() {
		s = &ImsSiteMulti{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Title,
			&s.Styleid,
			&s.SiteInfo,
			&s.Status,
			&s.Bindhost,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteMultiScanAll(rows *sql.Rows) (ss []*ImsSiteMulti, err error) {
	for rows.Next() {
		s := &ImsSiteMulti{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Title,
			&s.Styleid,
			&s.SiteInfo,
			&s.Status,
			&s.Bindhost,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteNavScan(rows *sql.Rows) (s *ImsSiteNav, err error) {
	if rows.Next() {
		s = &ImsSiteNav{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Section,
			&s.Module,
			&s.Displayorder,
			&s.Name,
			&s.Description,
			&s.Position,
			&s.Url,
			&s.Icon,
			&s.Css,
			&s.Status,
			&s.Categoryid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteNavScanAll(rows *sql.Rows) (ss []*ImsSiteNav, err error) {
	for rows.Next() {
		s := &ImsSiteNav{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Section,
			&s.Module,
			&s.Displayorder,
			&s.Name,
			&s.Description,
			&s.Position,
			&s.Url,
			&s.Icon,
			&s.Css,
			&s.Status,
			&s.Categoryid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSitePageScan(rows *sql.Rows) (s *ImsSitePage, err error) {
	if rows.Next() {
		s = &ImsSitePage{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Title,
			&s.Description,
			&s.Params,
			&s.Html,
			&s.Multipage,
			&s.Type,
			&s.Status,
			&s.Createtime,
			&s.Goodnum,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSitePageScanAll(rows *sql.Rows) (ss []*ImsSitePage, err error) {
	for rows.Next() {
		s := &ImsSitePage{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Title,
			&s.Description,
			&s.Params,
			&s.Html,
			&s.Multipage,
			&s.Type,
			&s.Status,
			&s.Createtime,
			&s.Goodnum,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteSlideScan(rows *sql.Rows) (s *ImsSiteSlide, err error) {
	if rows.Next() {
		s = &ImsSiteSlide{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Title,
			&s.Url,
			&s.Thumb,
			&s.Displayorder,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteSlideScanAll(rows *sql.Rows) (ss []*ImsSiteSlide, err error) {
	for rows.Next() {
		s := &ImsSiteSlide{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Title,
			&s.Url,
			&s.Thumb,
			&s.Displayorder,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteStoreCashLogScan(rows *sql.Rows) (s *ImsSiteStoreCashLog, err error) {
	if rows.Next() {
		s = &ImsSiteStoreCashLog{}
		err = rows.Scan(
			&s.Id,
			&s.FounderUid,
			&s.Number,
			&s.Amount,
			&s.Status,
			&s.CreateTime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteStoreCashLogScanAll(rows *sql.Rows) (ss []*ImsSiteStoreCashLog, err error) {
	for rows.Next() {
		s := &ImsSiteStoreCashLog{}
		err = rows.Scan(
			&s.Id,
			&s.FounderUid,
			&s.Number,
			&s.Amount,
			&s.Status,
			&s.CreateTime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteStoreCashOrderScan(rows *sql.Rows) (s *ImsSiteStoreCashOrder, err error) {
	if rows.Next() {
		s = &ImsSiteStoreCashOrder{}
		err = rows.Scan(
			&s.Id,
			&s.Number,
			&s.FounderUid,
			&s.OrderId,
			&s.GoodsId,
			&s.OrderAmount,
			&s.CashLogId,
			&s.Status,
			&s.CreateTime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteStoreCashOrderScanAll(rows *sql.Rows) (ss []*ImsSiteStoreCashOrder, err error) {
	for rows.Next() {
		s := &ImsSiteStoreCashOrder{}
		err = rows.Scan(
			&s.Id,
			&s.Number,
			&s.FounderUid,
			&s.OrderId,
			&s.GoodsId,
			&s.OrderAmount,
			&s.CashLogId,
			&s.Status,
			&s.CreateTime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteStoreCreateAccountScan(rows *sql.Rows) (s *ImsSiteStoreCreateAccount, err error) {
	if rows.Next() {
		s = &ImsSiteStoreCreateAccount{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Uniacid,
			&s.Type,
			&s.Endtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteStoreCreateAccountScanAll(rows *sql.Rows) (ss []*ImsSiteStoreCreateAccount, err error) {
	for rows.Next() {
		s := &ImsSiteStoreCreateAccount{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Uniacid,
			&s.Type,
			&s.Endtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteStoreGoodsScan(rows *sql.Rows) (s *ImsSiteStoreGoods, err error) {
	if rows.Next() {
		s = &ImsSiteStoreGoods{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Title,
			&s.Module,
			&s.AccountNum,
			&s.WxappNum,
			&s.Price,
			&s.Unit,
			&s.Slide,
			&s.CategoryId,
			&s.TitleInitial,
			&s.Status,
			&s.Createtime,
			&s.Synopsis,
			&s.Description,
			&s.ModuleGroup,
			&s.ApiNum,
			&s.UserGroupPrice,
			&s.UserGroup,
			&s.AccountGroup,
			&s.IsWish,
			&s.Logo,
			&s.PlatformNum,
			&s.AliappNum,
			&s.BaiduappNum,
			&s.PhoneappNum,
			&s.ToutiaoappNum,
			&s.WebappNum,
			&s.XzappNum,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteStoreGoodsScanAll(rows *sql.Rows) (ss []*ImsSiteStoreGoods, err error) {
	for rows.Next() {
		s := &ImsSiteStoreGoods{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Title,
			&s.Module,
			&s.AccountNum,
			&s.WxappNum,
			&s.Price,
			&s.Unit,
			&s.Slide,
			&s.CategoryId,
			&s.TitleInitial,
			&s.Status,
			&s.Createtime,
			&s.Synopsis,
			&s.Description,
			&s.ModuleGroup,
			&s.ApiNum,
			&s.UserGroupPrice,
			&s.UserGroup,
			&s.AccountGroup,
			&s.IsWish,
			&s.Logo,
			&s.PlatformNum,
			&s.AliappNum,
			&s.BaiduappNum,
			&s.PhoneappNum,
			&s.ToutiaoappNum,
			&s.WebappNum,
			&s.XzappNum,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteStoreGoodsCloudScan(rows *sql.Rows) (s *ImsSiteStoreGoodsCloud, err error) {
	if rows.Next() {
		s = &ImsSiteStoreGoodsCloud{}
		err = rows.Scan(
			&s.Id,
			&s.CloudId,
			&s.Name,
			&s.Title,
			&s.Logo,
			&s.WishBranch,
			&s.IsEdited,
			&s.Isdeleted,
			&s.Branchs,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteStoreGoodsCloudScanAll(rows *sql.Rows) (ss []*ImsSiteStoreGoodsCloud, err error) {
	for rows.Next() {
		s := &ImsSiteStoreGoodsCloud{}
		err = rows.Scan(
			&s.Id,
			&s.CloudId,
			&s.Name,
			&s.Title,
			&s.Logo,
			&s.WishBranch,
			&s.IsEdited,
			&s.Isdeleted,
			&s.Branchs,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteStoreOrderScan(rows *sql.Rows) (s *ImsSiteStoreOrder, err error) {
	if rows.Next() {
		s = &ImsSiteStoreOrder{}
		err = rows.Scan(
			&s.Id,
			&s.Orderid,
			&s.Goodsid,
			&s.Duration,
			&s.Buyer,
			&s.Buyerid,
			&s.Amount,
			&s.Type,
			&s.Changeprice,
			&s.Createtime,
			&s.Uniacid,
			&s.Endtime,
			&s.Wxapp,
			&s.IsWish,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteStoreOrderScanAll(rows *sql.Rows) (ss []*ImsSiteStoreOrder, err error) {
	for rows.Next() {
		s := &ImsSiteStoreOrder{}
		err = rows.Scan(
			&s.Id,
			&s.Orderid,
			&s.Goodsid,
			&s.Duration,
			&s.Buyer,
			&s.Buyerid,
			&s.Amount,
			&s.Type,
			&s.Changeprice,
			&s.Createtime,
			&s.Uniacid,
			&s.Endtime,
			&s.Wxapp,
			&s.IsWish,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteStylesScan(rows *sql.Rows) (s *ImsSiteStyles, err error) {
	if rows.Next() {
		s = &ImsSiteStyles{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Templateid,
			&s.Name,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteStylesScanAll(rows *sql.Rows) (ss []*ImsSiteStyles, err error) {
	for rows.Next() {
		s := &ImsSiteStyles{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Templateid,
			&s.Name,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteStylesVarsScan(rows *sql.Rows) (s *ImsSiteStylesVars, err error) {
	if rows.Next() {
		s = &ImsSiteStylesVars{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Templateid,
			&s.Styleid,
			&s.Variable,
			&s.Content,
			&s.Description,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteStylesVarsScanAll(rows *sql.Rows) (ss []*ImsSiteStylesVars, err error) {
	for rows.Next() {
		s := &ImsSiteStylesVars{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Templateid,
			&s.Styleid,
			&s.Variable,
			&s.Content,
			&s.Description,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSiteTemplatesScan(rows *sql.Rows) (s *ImsSiteTemplates, err error) {
	if rows.Next() {
		s = &ImsSiteTemplates{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.Title,
			&s.Version,
			&s.Description,
			&s.Author,
			&s.Url,
			&s.Type,
			&s.Sections,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSiteTemplatesScanAll(rows *sql.Rows) (ss []*ImsSiteTemplates, err error) {
	for rows.Next() {
		s := &ImsSiteTemplates{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.Title,
			&s.Version,
			&s.Description,
			&s.Author,
			&s.Url,
			&s.Type,
			&s.Sections,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsStatFansScan(rows *sql.Rows) (s *ImsStatFans, err error) {
	if rows.Next() {
		s = &ImsStatFans{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.New,
			&s.Cancel,
			&s.Cumulate,
			&s.Date,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsStatFansScanAll(rows *sql.Rows) (ss []*ImsStatFans, err error) {
	for rows.Next() {
		s := &ImsStatFans{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.New,
			&s.Cancel,
			&s.Cumulate,
			&s.Date,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsStatKeywordScan(rows *sql.Rows) (s *ImsStatKeyword, err error) {
	if rows.Next() {
		s = &ImsStatKeyword{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Rid,
			&s.Kid,
			&s.Hit,
			&s.Lastupdate,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsStatKeywordScanAll(rows *sql.Rows) (ss []*ImsStatKeyword, err error) {
	for rows.Next() {
		s := &ImsStatKeyword{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Rid,
			&s.Kid,
			&s.Hit,
			&s.Lastupdate,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsStatMsgHistoryScan(rows *sql.Rows) (s *ImsStatMsgHistory, err error) {
	if rows.Next() {
		s = &ImsStatMsgHistory{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Rid,
			&s.Kid,
			&s.FromUser,
			&s.Module,
			&s.Message,
			&s.Type,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsStatMsgHistoryScanAll(rows *sql.Rows) (ss []*ImsStatMsgHistory, err error) {
	for rows.Next() {
		s := &ImsStatMsgHistory{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Rid,
			&s.Kid,
			&s.FromUser,
			&s.Module,
			&s.Message,
			&s.Type,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsStatRuleScan(rows *sql.Rows) (s *ImsStatRule, err error) {
	if rows.Next() {
		s = &ImsStatRule{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Rid,
			&s.Hit,
			&s.Lastupdate,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsStatRuleScanAll(rows *sql.Rows) (ss []*ImsStatRule, err error) {
	for rows.Next() {
		s := &ImsStatRule{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Rid,
			&s.Hit,
			&s.Lastupdate,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsStatVisitScan(rows *sql.Rows) (s *ImsStatVisit, err error) {
	if rows.Next() {
		s = &ImsStatVisit{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Module,
			&s.Count,
			&s.Date,
			&s.Type,
			&s.IpCount,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsStatVisitScanAll(rows *sql.Rows) (ss []*ImsStatVisit, err error) {
	for rows.Next() {
		s := &ImsStatVisit{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Module,
			&s.Count,
			&s.Date,
			&s.Type,
			&s.IpCount,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsStatVisitIpScan(rows *sql.Rows) (s *ImsStatVisitIp, err error) {
	if rows.Next() {
		s = &ImsStatVisitIp{}
		err = rows.Scan(
			&s.Id,
			&s.Ip,
			&s.Uniacid,
			&s.Type,
			&s.Module,
			&s.Date,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsStatVisitIpScanAll(rows *sql.Rows) (ss []*ImsStatVisitIp, err error) {
	for rows.Next() {
		s := &ImsStatVisitIp{}
		err = rows.Scan(
			&s.Id,
			&s.Ip,
			&s.Uniacid,
			&s.Type,
			&s.Module,
			&s.Date,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSystemStatVisitScan(rows *sql.Rows) (s *ImsSystemStatVisit, err error) {
	if rows.Next() {
		s = &ImsSystemStatVisit{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Modulename,
			&s.Uid,
			&s.Displayorder,
			&s.Createtime,
			&s.Updatetime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSystemStatVisitScanAll(rows *sql.Rows) (ss []*ImsSystemStatVisit, err error) {
	for rows.Next() {
		s := &ImsSystemStatVisit{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Modulename,
			&s.Uid,
			&s.Displayorder,
			&s.Createtime,
			&s.Updatetime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsSystemWelcomeBinddomainScan(rows *sql.Rows) (s *ImsSystemWelcomeBinddomain, err error) {
	if rows.Next() {
		s = &ImsSystemWelcomeBinddomain{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.ModuleName,
			&s.Domain,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsSystemWelcomeBinddomainScanAll(rows *sql.Rows) (ss []*ImsSystemWelcomeBinddomain, err error) {
	for rows.Next() {
		s := &ImsSystemWelcomeBinddomain{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.ModuleName,
			&s.Domain,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniAccountScan(rows *sql.Rows) (s *ImsUniAccount, err error) {
	if rows.Next() {
		s = &ImsUniAccount{}
		err = rows.Scan(
			&s.Uniacid,
			&s.Groupid,
			&s.Name,
			&s.Description,
			&s.DefaultAcid,
			&s.Rank,
			&s.TitleInitial,
			&s.Createtime,
			&s.Logo,
			&s.Qrcode,
			&s.CreateUid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniAccountScanAll(rows *sql.Rows) (ss []*ImsUniAccount, err error) {
	for rows.Next() {
		s := &ImsUniAccount{}
		err = rows.Scan(
			&s.Uniacid,
			&s.Groupid,
			&s.Name,
			&s.Description,
			&s.DefaultAcid,
			&s.Rank,
			&s.TitleInitial,
			&s.Createtime,
			&s.Logo,
			&s.Qrcode,
			&s.CreateUid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniAccountExtraModulesScan(rows *sql.Rows) (s *ImsUniAccountExtraModules, err error) {
	if rows.Next() {
		s = &ImsUniAccountExtraModules{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Modules,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniAccountExtraModulesScanAll(rows *sql.Rows) (ss []*ImsUniAccountExtraModules, err error) {
	for rows.Next() {
		s := &ImsUniAccountExtraModules{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Modules,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniAccountGroupScan(rows *sql.Rows) (s *ImsUniAccountGroup, err error) {
	if rows.Next() {
		s = &ImsUniAccountGroup{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Groupid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniAccountGroupScanAll(rows *sql.Rows) (ss []*ImsUniAccountGroup, err error) {
	for rows.Next() {
		s := &ImsUniAccountGroup{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Groupid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniAccountMenusScan(rows *sql.Rows) (s *ImsUniAccountMenus, err error) {
	if rows.Next() {
		s = &ImsUniAccountMenus{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Menuid,
			&s.Type,
			&s.Title,
			&s.Sex,
			&s.GroupId,
			&s.ClientPlatformType,
			&s.Area,
			&s.Data,
			&s.Status,
			&s.Createtime,
			&s.Isdeleted,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniAccountMenusScanAll(rows *sql.Rows) (ss []*ImsUniAccountMenus, err error) {
	for rows.Next() {
		s := &ImsUniAccountMenus{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Menuid,
			&s.Type,
			&s.Title,
			&s.Sex,
			&s.GroupId,
			&s.ClientPlatformType,
			&s.Area,
			&s.Data,
			&s.Status,
			&s.Createtime,
			&s.Isdeleted,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniAccountModulesScan(rows *sql.Rows) (s *ImsUniAccountModules, err error) {
	if rows.Next() {
		s = &ImsUniAccountModules{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Module,
			&s.Enabled,
			&s.Settings,
			&s.Shortcut,
			&s.Displayorder,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniAccountModulesScanAll(rows *sql.Rows) (ss []*ImsUniAccountModules, err error) {
	for rows.Next() {
		s := &ImsUniAccountModules{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Module,
			&s.Enabled,
			&s.Settings,
			&s.Shortcut,
			&s.Displayorder,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniAccountModulesShortcutScan(rows *sql.Rows) (s *ImsUniAccountModulesShortcut, err error) {
	if rows.Next() {
		s = &ImsUniAccountModulesShortcut{}
		err = rows.Scan(
			&s.Id,
			&s.Title,
			&s.Url,
			&s.Icon,
			&s.Uniacid,
			&s.VersionId,
			&s.ModuleName,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniAccountModulesShortcutScanAll(rows *sql.Rows) (ss []*ImsUniAccountModulesShortcut, err error) {
	for rows.Next() {
		s := &ImsUniAccountModulesShortcut{}
		err = rows.Scan(
			&s.Id,
			&s.Title,
			&s.Url,
			&s.Icon,
			&s.Uniacid,
			&s.VersionId,
			&s.ModuleName,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniAccountUsersScan(rows *sql.Rows) (s *ImsUniAccountUsers, err error) {
	if rows.Next() {
		s = &ImsUniAccountUsers{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Role,
			&s.Rank,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniAccountUsersScanAll(rows *sql.Rows) (ss []*ImsUniAccountUsers, err error) {
	for rows.Next() {
		s := &ImsUniAccountUsers{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Role,
			&s.Rank,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniGroupScan(rows *sql.Rows) (s *ImsUniGroup, err error) {
	if rows.Next() {
		s = &ImsUniGroup{}
		err = rows.Scan(
			&s.Id,
			&s.OwnerUid,
			&s.Name,
			&s.Modules,
			&s.Templates,
			&s.Uniacid,
			&s.Uid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniGroupScanAll(rows *sql.Rows) (ss []*ImsUniGroup, err error) {
	for rows.Next() {
		s := &ImsUniGroup{}
		err = rows.Scan(
			&s.Id,
			&s.OwnerUid,
			&s.Name,
			&s.Modules,
			&s.Templates,
			&s.Uniacid,
			&s.Uid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniLinkUniacidScan(rows *sql.Rows) (s *ImsUniLinkUniacid, err error) {
	if rows.Next() {
		s = &ImsUniLinkUniacid{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.LinkUniacid,
			&s.VersionId,
			&s.ModuleName,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniLinkUniacidScanAll(rows *sql.Rows) (ss []*ImsUniLinkUniacid, err error) {
	for rows.Next() {
		s := &ImsUniLinkUniacid{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.LinkUniacid,
			&s.VersionId,
			&s.ModuleName,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniModulesScan(rows *sql.Rows) (s *ImsUniModules, err error) {
	if rows.Next() {
		s = &ImsUniModules{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.ModuleName,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniModulesScanAll(rows *sql.Rows) (ss []*ImsUniModules, err error) {
	for rows.Next() {
		s := &ImsUniModules{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.ModuleName,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniSettingsScan(rows *sql.Rows) (s *ImsUniSettings, err error) {
	if rows.Next() {
		s = &ImsUniSettings{}
		err = rows.Scan(
			&s.Uniacid,
			&s.Passport,
			&s.Oauth,
			&s.JsauthAcid,
			&s.Notify,
			&s.Creditnames,
			&s.Creditbehaviors,
			&s.Welcome,
			&s.Default,
			&s.DefaultMessage,
			&s.Payment,
			&s.Stat,
			&s.DefaultSite,
			&s.Sync,
			&s.Recharge,
			&s.Tplnotice,
			&s.Grouplevel,
			&s.Mcplugin,
			&s.ExchangeEnable,
			&s.CouponType,
			&s.Menuset,
			&s.Statistics,
			&s.BindDomain,
			&s.CommentStatus,
			&s.ReplySetting,
			&s.DefaultModule,
			&s.AttachmentLimit,
			&s.AttachmentSize,
			&s.SyncMember,
			&s.Remote,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniSettingsScanAll(rows *sql.Rows) (ss []*ImsUniSettings, err error) {
	for rows.Next() {
		s := &ImsUniSettings{}
		err = rows.Scan(
			&s.Uniacid,
			&s.Passport,
			&s.Oauth,
			&s.JsauthAcid,
			&s.Notify,
			&s.Creditnames,
			&s.Creditbehaviors,
			&s.Welcome,
			&s.Default,
			&s.DefaultMessage,
			&s.Payment,
			&s.Stat,
			&s.DefaultSite,
			&s.Sync,
			&s.Recharge,
			&s.Tplnotice,
			&s.Grouplevel,
			&s.Mcplugin,
			&s.ExchangeEnable,
			&s.CouponType,
			&s.Menuset,
			&s.Statistics,
			&s.BindDomain,
			&s.CommentStatus,
			&s.ReplySetting,
			&s.DefaultModule,
			&s.AttachmentLimit,
			&s.AttachmentSize,
			&s.SyncMember,
			&s.Remote,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUniVerifycodeScan(rows *sql.Rows) (s *ImsUniVerifycode, err error) {
	if rows.Next() {
		s = &ImsUniVerifycode{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Receiver,
			&s.Verifycode,
			&s.Total,
			&s.Createtime,
			&s.FailedCount,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUniVerifycodeScanAll(rows *sql.Rows) (ss []*ImsUniVerifycode, err error) {
	for rows.Next() {
		s := &ImsUniVerifycode{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Receiver,
			&s.Verifycode,
			&s.Total,
			&s.Createtime,
			&s.FailedCount,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUserapiCacheScan(rows *sql.Rows) (s *ImsUserapiCache, err error) {
	if rows.Next() {
		s = &ImsUserapiCache{}
		err = rows.Scan(
			&s.Id,
			&s.Key,
			&s.Content,
			&s.Lastupdate,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUserapiCacheScanAll(rows *sql.Rows) (ss []*ImsUserapiCache, err error) {
	for rows.Next() {
		s := &ImsUserapiCache{}
		err = rows.Scan(
			&s.Id,
			&s.Key,
			&s.Content,
			&s.Lastupdate,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUserapiReplyScan(rows *sql.Rows) (s *ImsUserapiReply, err error) {
	if rows.Next() {
		s = &ImsUserapiReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Description,
			&s.Apiurl,
			&s.Token,
			&s.DefaultText,
			&s.Cachetime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUserapiReplyScanAll(rows *sql.Rows) (ss []*ImsUserapiReply, err error) {
	for rows.Next() {
		s := &ImsUserapiReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Description,
			&s.Apiurl,
			&s.Token,
			&s.DefaultText,
			&s.Cachetime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersScan(rows *sql.Rows) (s *ImsUsers, err error) {
	if rows.Next() {
		s = &ImsUsers{}
		err = rows.Scan(
			&s.Uid,
			&s.OwnerUid,
			&s.Groupid,
			&s.FounderGroupid,
			&s.Username,
			&s.Password,
			&s.Salt,
			&s.Type,
			&s.Status,
			&s.Joindate,
			&s.Joinip,
			&s.Lastvisit,
			&s.Lastip,
			&s.Remark,
			&s.Starttime,
			&s.Endtime,
			&s.RegisterType,
			&s.Openid,
			&s.WelcomeLink,
			&s.NoticeSetting,
			&s.IsBind,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersScanAll(rows *sql.Rows) (ss []*ImsUsers, err error) {
	for rows.Next() {
		s := &ImsUsers{}
		err = rows.Scan(
			&s.Uid,
			&s.OwnerUid,
			&s.Groupid,
			&s.FounderGroupid,
			&s.Username,
			&s.Password,
			&s.Salt,
			&s.Type,
			&s.Status,
			&s.Joindate,
			&s.Joinip,
			&s.Lastvisit,
			&s.Lastip,
			&s.Remark,
			&s.Starttime,
			&s.Endtime,
			&s.RegisterType,
			&s.Openid,
			&s.WelcomeLink,
			&s.NoticeSetting,
			&s.IsBind,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersBindScan(rows *sql.Rows) (s *ImsUsersBind, err error) {
	if rows.Next() {
		s = &ImsUsersBind{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.BindSign,
			&s.ThirdType,
			&s.ThirdNickname,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersBindScanAll(rows *sql.Rows) (ss []*ImsUsersBind, err error) {
	for rows.Next() {
		s := &ImsUsersBind{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.BindSign,
			&s.ThirdType,
			&s.ThirdNickname,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersCreateGroupScan(rows *sql.Rows) (s *ImsUsersCreateGroup, err error) {
	if rows.Next() {
		s = &ImsUsersCreateGroup{}
		err = rows.Scan(
			&s.Id,
			&s.GroupName,
			&s.Maxaccount,
			&s.Maxwxapp,
			&s.Maxwebapp,
			&s.Maxphoneapp,
			&s.Maxxzapp,
			&s.Maxaliapp,
			&s.Createtime,
			&s.Maxbaiduapp,
			&s.Maxtoutiaoapp,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersCreateGroupScanAll(rows *sql.Rows) (ss []*ImsUsersCreateGroup, err error) {
	for rows.Next() {
		s := &ImsUsersCreateGroup{}
		err = rows.Scan(
			&s.Id,
			&s.GroupName,
			&s.Maxaccount,
			&s.Maxwxapp,
			&s.Maxwebapp,
			&s.Maxphoneapp,
			&s.Maxxzapp,
			&s.Maxaliapp,
			&s.Createtime,
			&s.Maxbaiduapp,
			&s.Maxtoutiaoapp,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersExtraGroupScan(rows *sql.Rows) (s *ImsUsersExtraGroup, err error) {
	if rows.Next() {
		s = &ImsUsersExtraGroup{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.UniGroupId,
			&s.CreateGroupId,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersExtraGroupScanAll(rows *sql.Rows) (ss []*ImsUsersExtraGroup, err error) {
	for rows.Next() {
		s := &ImsUsersExtraGroup{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.UniGroupId,
			&s.CreateGroupId,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersExtraLimitScan(rows *sql.Rows) (s *ImsUsersExtraLimit, err error) {
	if rows.Next() {
		s = &ImsUsersExtraLimit{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Maxaccount,
			&s.Maxwxapp,
			&s.Maxwebapp,
			&s.Maxphoneapp,
			&s.Maxxzapp,
			&s.Maxaliapp,
			&s.Timelimit,
			&s.Maxbaiduapp,
			&s.Maxtoutiaoapp,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersExtraLimitScanAll(rows *sql.Rows) (ss []*ImsUsersExtraLimit, err error) {
	for rows.Next() {
		s := &ImsUsersExtraLimit{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Maxaccount,
			&s.Maxwxapp,
			&s.Maxwebapp,
			&s.Maxphoneapp,
			&s.Maxxzapp,
			&s.Maxaliapp,
			&s.Timelimit,
			&s.Maxbaiduapp,
			&s.Maxtoutiaoapp,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersExtraModulesScan(rows *sql.Rows) (s *ImsUsersExtraModules, err error) {
	if rows.Next() {
		s = &ImsUsersExtraModules{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.ModuleName,
			&s.Support,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersExtraModulesScanAll(rows *sql.Rows) (ss []*ImsUsersExtraModules, err error) {
	for rows.Next() {
		s := &ImsUsersExtraModules{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.ModuleName,
			&s.Support,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersExtraTemplatesScan(rows *sql.Rows) (s *ImsUsersExtraTemplates, err error) {
	if rows.Next() {
		s = &ImsUsersExtraTemplates{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.TemplateId,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersExtraTemplatesScanAll(rows *sql.Rows) (ss []*ImsUsersExtraTemplates, err error) {
	for rows.Next() {
		s := &ImsUsersExtraTemplates{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.TemplateId,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersFailedLoginScan(rows *sql.Rows) (s *ImsUsersFailedLogin, err error) {
	if rows.Next() {
		s = &ImsUsersFailedLogin{}
		err = rows.Scan(
			&s.Id,
			&s.Ip,
			&s.Username,
			&s.Count,
			&s.Lastupdate,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersFailedLoginScanAll(rows *sql.Rows) (ss []*ImsUsersFailedLogin, err error) {
	for rows.Next() {
		s := &ImsUsersFailedLogin{}
		err = rows.Scan(
			&s.Id,
			&s.Ip,
			&s.Username,
			&s.Count,
			&s.Lastupdate,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersFounderGroupScan(rows *sql.Rows) (s *ImsUsersFounderGroup, err error) {
	if rows.Next() {
		s = &ImsUsersFounderGroup{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.Package,
			&s.Maxaccount,
			&s.Timelimit,
			&s.Maxwxapp,
			&s.Maxwebapp,
			&s.Maxphoneapp,
			&s.Maxxzapp,
			&s.Maxaliapp,
			&s.Maxbaiduapp,
			&s.Maxtoutiaoapp,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersFounderGroupScanAll(rows *sql.Rows) (ss []*ImsUsersFounderGroup, err error) {
	for rows.Next() {
		s := &ImsUsersFounderGroup{}
		err = rows.Scan(
			&s.Id,
			&s.Name,
			&s.Package,
			&s.Maxaccount,
			&s.Timelimit,
			&s.Maxwxapp,
			&s.Maxwebapp,
			&s.Maxphoneapp,
			&s.Maxxzapp,
			&s.Maxaliapp,
			&s.Maxbaiduapp,
			&s.Maxtoutiaoapp,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersFounderOwnCreateGroupsScan(rows *sql.Rows) (s *ImsUsersFounderOwnCreateGroups, err error) {
	if rows.Next() {
		s = &ImsUsersFounderOwnCreateGroups{}
		err = rows.Scan(
			&s.Id,
			&s.FounderUid,
			&s.CreateGroupId,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersFounderOwnCreateGroupsScanAll(rows *sql.Rows) (ss []*ImsUsersFounderOwnCreateGroups, err error) {
	for rows.Next() {
		s := &ImsUsersFounderOwnCreateGroups{}
		err = rows.Scan(
			&s.Id,
			&s.FounderUid,
			&s.CreateGroupId,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersFounderOwnUniGroupsScan(rows *sql.Rows) (s *ImsUsersFounderOwnUniGroups, err error) {
	if rows.Next() {
		s = &ImsUsersFounderOwnUniGroups{}
		err = rows.Scan(
			&s.Id,
			&s.FounderUid,
			&s.UniGroupId,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersFounderOwnUniGroupsScanAll(rows *sql.Rows) (ss []*ImsUsersFounderOwnUniGroups, err error) {
	for rows.Next() {
		s := &ImsUsersFounderOwnUniGroups{}
		err = rows.Scan(
			&s.Id,
			&s.FounderUid,
			&s.UniGroupId,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersFounderOwnUsersScan(rows *sql.Rows) (s *ImsUsersFounderOwnUsers, err error) {
	if rows.Next() {
		s = &ImsUsersFounderOwnUsers{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.FounderUid,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersFounderOwnUsersScanAll(rows *sql.Rows) (ss []*ImsUsersFounderOwnUsers, err error) {
	for rows.Next() {
		s := &ImsUsersFounderOwnUsers{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.FounderUid,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersFounderOwnUsersGroupsScan(rows *sql.Rows) (s *ImsUsersFounderOwnUsersGroups, err error) {
	if rows.Next() {
		s = &ImsUsersFounderOwnUsersGroups{}
		err = rows.Scan(
			&s.Id,
			&s.FounderUid,
			&s.UsersGroupId,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersFounderOwnUsersGroupsScanAll(rows *sql.Rows) (ss []*ImsUsersFounderOwnUsersGroups, err error) {
	for rows.Next() {
		s := &ImsUsersFounderOwnUsersGroups{}
		err = rows.Scan(
			&s.Id,
			&s.FounderUid,
			&s.UsersGroupId,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersGroupScan(rows *sql.Rows) (s *ImsUsersGroup, err error) {
	if rows.Next() {
		s = &ImsUsersGroup{}
		err = rows.Scan(
			&s.Id,
			&s.OwnerUid,
			&s.Name,
			&s.Package,
			&s.Maxaccount,
			&s.Timelimit,
			&s.Maxwxapp,
			&s.Maxwebapp,
			&s.Maxphoneapp,
			&s.Maxxzapp,
			&s.Maxaliapp,
			&s.Maxbaiduapp,
			&s.Maxtoutiaoapp,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersGroupScanAll(rows *sql.Rows) (ss []*ImsUsersGroup, err error) {
	for rows.Next() {
		s := &ImsUsersGroup{}
		err = rows.Scan(
			&s.Id,
			&s.OwnerUid,
			&s.Name,
			&s.Package,
			&s.Maxaccount,
			&s.Timelimit,
			&s.Maxwxapp,
			&s.Maxwebapp,
			&s.Maxphoneapp,
			&s.Maxxzapp,
			&s.Maxaliapp,
			&s.Maxbaiduapp,
			&s.Maxtoutiaoapp,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersInvitationScan(rows *sql.Rows) (s *ImsUsersInvitation, err error) {
	if rows.Next() {
		s = &ImsUsersInvitation{}
		err = rows.Scan(
			&s.Id,
			&s.Code,
			&s.Fromuid,
			&s.Inviteuid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersInvitationScanAll(rows *sql.Rows) (ss []*ImsUsersInvitation, err error) {
	for rows.Next() {
		s := &ImsUsersInvitation{}
		err = rows.Scan(
			&s.Id,
			&s.Code,
			&s.Fromuid,
			&s.Inviteuid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersLastuseScan(rows *sql.Rows) (s *ImsUsersLastuse, err error) {
	if rows.Next() {
		s = &ImsUsersLastuse{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Uniacid,
			&s.Modulename,
			&s.Type,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersLastuseScanAll(rows *sql.Rows) (ss []*ImsUsersLastuse, err error) {
	for rows.Next() {
		s := &ImsUsersLastuse{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Uniacid,
			&s.Modulename,
			&s.Type,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersLoginLogsScan(rows *sql.Rows) (s *ImsUsersLoginLogs, err error) {
	if rows.Next() {
		s = &ImsUsersLoginLogs{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Ip,
			&s.City,
			&s.LoginAt,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersLoginLogsScanAll(rows *sql.Rows) (ss []*ImsUsersLoginLogs, err error) {
	for rows.Next() {
		s := &ImsUsersLoginLogs{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Ip,
			&s.City,
			&s.LoginAt,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersOperateHistoryScan(rows *sql.Rows) (s *ImsUsersOperateHistory, err error) {
	if rows.Next() {
		s = &ImsUsersOperateHistory{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Uid,
			&s.Uniacid,
			&s.ModuleName,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersOperateHistoryScanAll(rows *sql.Rows) (ss []*ImsUsersOperateHistory, err error) {
	for rows.Next() {
		s := &ImsUsersOperateHistory{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Uid,
			&s.Uniacid,
			&s.ModuleName,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersOperateStarScan(rows *sql.Rows) (s *ImsUsersOperateStar, err error) {
	if rows.Next() {
		s = &ImsUsersOperateStar{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Uid,
			&s.Uniacid,
			&s.ModuleName,
			&s.Rank,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersOperateStarScanAll(rows *sql.Rows) (ss []*ImsUsersOperateStar, err error) {
	for rows.Next() {
		s := &ImsUsersOperateStar{}
		err = rows.Scan(
			&s.Id,
			&s.Type,
			&s.Uid,
			&s.Uniacid,
			&s.ModuleName,
			&s.Rank,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersPermissionScan(rows *sql.Rows) (s *ImsUsersPermission, err error) {
	if rows.Next() {
		s = &ImsUsersPermission{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Type,
			&s.Permission,
			&s.Url,
			&s.Modules,
			&s.Templates,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersPermissionScanAll(rows *sql.Rows) (ss []*ImsUsersPermission, err error) {
	for rows.Next() {
		s := &ImsUsersPermission{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Uid,
			&s.Type,
			&s.Permission,
			&s.Url,
			&s.Modules,
			&s.Templates,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsUsersProfileScan(rows *sql.Rows) (s *ImsUsersProfile, err error) {
	if rows.Next() {
		s = &ImsUsersProfile{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Createtime,
			&s.Edittime,
			&s.Realname,
			&s.Nickname,
			&s.Avatar,
			&s.Qq,
			&s.Mobile,
			&s.Fakeid,
			&s.Vip,
			&s.Gender,
			&s.Birthyear,
			&s.Birthmonth,
			&s.Birthday,
			&s.Constellation,
			&s.Zodiac,
			&s.Telephone,
			&s.Idcard,
			&s.Studentid,
			&s.Grade,
			&s.Address,
			&s.Zipcode,
			&s.Nationality,
			&s.Resideprovince,
			&s.Residecity,
			&s.Residedist,
			&s.Graduateschool,
			&s.Company,
			&s.Education,
			&s.Occupation,
			&s.Position,
			&s.Revenue,
			&s.Affectivestatus,
			&s.Lookingfor,
			&s.Bloodtype,
			&s.Height,
			&s.Weight,
			&s.Alipay,
			&s.Msn,
			&s.Email,
			&s.Taobao,
			&s.Site,
			&s.Bio,
			&s.Interest,
			&s.Workerid,
			&s.IsSendMobileStatus,
			&s.SendExpireStatus,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsUsersProfileScanAll(rows *sql.Rows) (ss []*ImsUsersProfile, err error) {
	for rows.Next() {
		s := &ImsUsersProfile{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Createtime,
			&s.Edittime,
			&s.Realname,
			&s.Nickname,
			&s.Avatar,
			&s.Qq,
			&s.Mobile,
			&s.Fakeid,
			&s.Vip,
			&s.Gender,
			&s.Birthyear,
			&s.Birthmonth,
			&s.Birthday,
			&s.Constellation,
			&s.Zodiac,
			&s.Telephone,
			&s.Idcard,
			&s.Studentid,
			&s.Grade,
			&s.Address,
			&s.Zipcode,
			&s.Nationality,
			&s.Resideprovince,
			&s.Residecity,
			&s.Residedist,
			&s.Graduateschool,
			&s.Company,
			&s.Education,
			&s.Occupation,
			&s.Position,
			&s.Revenue,
			&s.Affectivestatus,
			&s.Lookingfor,
			&s.Bloodtype,
			&s.Height,
			&s.Weight,
			&s.Alipay,
			&s.Msn,
			&s.Email,
			&s.Taobao,
			&s.Site,
			&s.Bio,
			&s.Interest,
			&s.Workerid,
			&s.IsSendMobileStatus,
			&s.SendExpireStatus,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVideoReplyScan(rows *sql.Rows) (s *ImsVideoReply, err error) {
	if rows.Next() {
		s = &ImsVideoReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.Description,
			&s.Mediaid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVideoReplyScanAll(rows *sql.Rows) (ss []*ImsVideoReply, err error) {
	for rows.Next() {
		s := &ImsVideoReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.Description,
			&s.Mediaid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVoiceReplyScan(rows *sql.Rows) (s *ImsVoiceReply, err error) {
	if rows.Next() {
		s = &ImsVoiceReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.Mediaid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVoiceReplyScanAll(rows *sql.Rows) (ss []*ImsVoiceReply, err error) {
	for rows.Next() {
		s := &ImsVoiceReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.Mediaid,
			&s.Createtime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanScan(rows *sql.Rows) (s *ImsVxFan, err error) {
	if rows.Next() {
		s = &ImsVxFan{}
		err = rows.Scan(
			&s.Id,
			&s.Pid,
			&s.AccountId,
			&s.AccountCategory,
			&s.Openid,
			&s.GroupId,
			&s.UnionId,
			&s.Mobile,
			&s.Category,
			&s.CategoryMsg,
			&s.Subscribe,
			&s.Nickname,
			&s.Sex,
			&s.Language,
			&s.City,
			&s.Province,
			&s.Country,
			&s.Avatar,
			&s.SubscribeAt,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
			&s.RecommendQrCode,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanScanAll(rows *sql.Rows) (ss []*ImsVxFan, err error) {
	for rows.Next() {
		s := &ImsVxFan{}
		err = rows.Scan(
			&s.Id,
			&s.Pid,
			&s.AccountId,
			&s.AccountCategory,
			&s.Openid,
			&s.GroupId,
			&s.UnionId,
			&s.Mobile,
			&s.Category,
			&s.CategoryMsg,
			&s.Subscribe,
			&s.Nickname,
			&s.Sex,
			&s.Language,
			&s.City,
			&s.Province,
			&s.Country,
			&s.Avatar,
			&s.SubscribeAt,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
			&s.RecommendQrCode,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanAccountScan(rows *sql.Rows) (s *ImsVxFanAccount, err error) {
	if rows.Next() {
		s = &ImsVxFanAccount{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Balance,
			&s.Prepare,
			&s.Withdraw,
			&s.LastWithdrawAt,
			&s.LastWithdrawDay,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanAccountScanAll(rows *sql.Rows) (ss []*ImsVxFanAccount, err error) {
	for rows.Next() {
		s := &ImsVxFanAccount{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Balance,
			&s.Prepare,
			&s.Withdraw,
			&s.LastWithdrawAt,
			&s.LastWithdrawDay,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanAddressScan(rows *sql.Rows) (s *ImsVxFanAddress, err error) {
	if rows.Next() {
		s = &ImsVxFanAddress{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Consignee,
			&s.Mobile,
			&s.Address1,
			&s.Address2,
			&s.Address3,
			&s.Address4,
			&s.PostalCode,
			&s.IsDefault,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.DeletedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanAddressScanAll(rows *sql.Rows) (ss []*ImsVxFanAddress, err error) {
	for rows.Next() {
		s := &ImsVxFanAddress{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Consignee,
			&s.Mobile,
			&s.Address1,
			&s.Address2,
			&s.Address3,
			&s.Address4,
			&s.PostalCode,
			&s.IsDefault,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.DeletedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanBalanceLogScan(rows *sql.Rows) (s *ImsVxFanBalanceLog, err error) {
	if rows.Next() {
		s = &ImsVxFanBalanceLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Val1,
			&s.Val2,
			&s.Val3,
			&s.Category,
			&s.CategoryMsg,
			&s.OrderId,
			&s.OrderNo,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanBalanceLogScanAll(rows *sql.Rows) (ss []*ImsVxFanBalanceLog, err error) {
	for rows.Next() {
		s := &ImsVxFanBalanceLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Val1,
			&s.Val2,
			&s.Val3,
			&s.Category,
			&s.CategoryMsg,
			&s.OrderId,
			&s.OrderNo,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanGoodsScan(rows *sql.Rows) (s *ImsVxFanGoods, err error) {
	if rows.Next() {
		s = &ImsVxFanGoods{}
		err = rows.Scan(
			&s.Id,
			&s.AccountId,
			&s.Avatar,
			&s.Name,
			&s.Title,
			&s.TitleDesc,
			&s.Label,
			&s.Details,
			&s.Video,
			&s.PriceNominal,
			&s.Price,
			&s.Postage,
			&s.Stock,
			&s.Sold,
			&s.Serial,
			&s.Category,
			&s.CategoryMsg,
			&s.Uid,
			&s.DeliveryAddress,
			&s.Recommend,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.DeletedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanGoodsScanAll(rows *sql.Rows) (ss []*ImsVxFanGoods, err error) {
	for rows.Next() {
		s := &ImsVxFanGoods{}
		err = rows.Scan(
			&s.Id,
			&s.AccountId,
			&s.Avatar,
			&s.Name,
			&s.Title,
			&s.TitleDesc,
			&s.Label,
			&s.Details,
			&s.Video,
			&s.PriceNominal,
			&s.Price,
			&s.Postage,
			&s.Stock,
			&s.Sold,
			&s.Serial,
			&s.Category,
			&s.CategoryMsg,
			&s.Uid,
			&s.DeliveryAddress,
			&s.Recommend,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.DeletedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanGoodsRuleScan(rows *sql.Rows) (s *ImsVxFanGoodsRule, err error) {
	if rows.Next() {
		s = &ImsVxFanGoodsRule{}
		err = rows.Scan(
			&s.Id,
			&s.GoodsId,
			&s.Attributes,
			&s.Serial,
			&s.Avatar,
			&s.Details,
			&s.PriceNominal,
			&s.Price,
			&s.Stock,
			&s.Sold,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.DeletedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanGoodsRuleScanAll(rows *sql.Rows) (ss []*ImsVxFanGoodsRule, err error) {
	for rows.Next() {
		s := &ImsVxFanGoodsRule{}
		err = rows.Scan(
			&s.Id,
			&s.GoodsId,
			&s.Attributes,
			&s.Serial,
			&s.Avatar,
			&s.Details,
			&s.PriceNominal,
			&s.Price,
			&s.Stock,
			&s.Sold,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.DeletedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanOrderBagQuotaCardScan(rows *sql.Rows) (s *ImsVxFanOrderBagQuotaCard, err error) {
	if rows.Next() {
		s = &ImsVxFanOrderBagQuotaCard{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.AccountId,
			&s.GoodsId,
			&s.GoodsNum,
			&s.GoodsPrice,
			&s.GoodsTitle,
			&s.OrderNo,
			&s.Amount,
			&s.PayAmount,
			&s.PayType,
			&s.PayState,
			&s.PayNo,
			&s.PayDetails,
			&s.GoodsDetails,
			&s.UserDetails,
			&s.State,
			&s.State1,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanOrderBagQuotaCardScanAll(rows *sql.Rows) (ss []*ImsVxFanOrderBagQuotaCard, err error) {
	for rows.Next() {
		s := &ImsVxFanOrderBagQuotaCard{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.AccountId,
			&s.GoodsId,
			&s.GoodsNum,
			&s.GoodsPrice,
			&s.GoodsTitle,
			&s.OrderNo,
			&s.Amount,
			&s.PayAmount,
			&s.PayType,
			&s.PayState,
			&s.PayNo,
			&s.PayDetails,
			&s.GoodsDetails,
			&s.UserDetails,
			&s.State,
			&s.State1,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanOrderGoodsScan(rows *sql.Rows) (s *ImsVxFanOrderGoods, err error) {
	if rows.Next() {
		s = &ImsVxFanOrderGoods{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.AccountId,
			&s.GoodsId,
			&s.GoodsNum,
			&s.GoodsPrice,
			&s.OrderNo,
			&s.Amount,
			&s.DiscountAmount,
			&s.Discount,
			&s.PayAmount,
			&s.PayType,
			&s.PayState,
			&s.PayNo,
			&s.PayDetails,
			&s.UserDetails,
			&s.GoodsDetails,
			&s.AddressDetails,
			&s.AddressId,
			&s.Address,
			&s.State,
			&s.State1,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanOrderGoodsScanAll(rows *sql.Rows) (ss []*ImsVxFanOrderGoods, err error) {
	for rows.Next() {
		s := &ImsVxFanOrderGoods{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.AccountId,
			&s.GoodsId,
			&s.GoodsNum,
			&s.GoodsPrice,
			&s.OrderNo,
			&s.Amount,
			&s.DiscountAmount,
			&s.Discount,
			&s.PayAmount,
			&s.PayType,
			&s.PayState,
			&s.PayNo,
			&s.PayDetails,
			&s.UserDetails,
			&s.GoodsDetails,
			&s.AddressDetails,
			&s.AddressId,
			&s.Address,
			&s.State,
			&s.State1,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanOrderRightCardScan(rows *sql.Rows) (s *ImsVxFanOrderRightCard, err error) {
	if rows.Next() {
		s = &ImsVxFanOrderRightCard{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.AccountId,
			&s.GoodsId,
			&s.GoodsNum,
			&s.GoodsPrice,
			&s.GoodsTitle,
			&s.OrderNo,
			&s.Amount,
			&s.PayAmount,
			&s.PayType,
			&s.PayState,
			&s.PayNo,
			&s.PayDetails,
			&s.GoodsDetails,
			&s.UserDetails,
			&s.State,
			&s.State1,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanOrderRightCardScanAll(rows *sql.Rows) (ss []*ImsVxFanOrderRightCard, err error) {
	for rows.Next() {
		s := &ImsVxFanOrderRightCard{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.AccountId,
			&s.GoodsId,
			&s.GoodsNum,
			&s.GoodsPrice,
			&s.GoodsTitle,
			&s.OrderNo,
			&s.Amount,
			&s.PayAmount,
			&s.PayType,
			&s.PayState,
			&s.PayNo,
			&s.PayDetails,
			&s.GoodsDetails,
			&s.UserDetails,
			&s.State,
			&s.State1,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanPrepareLogScan(rows *sql.Rows) (s *ImsVxFanPrepareLog, err error) {
	if rows.Next() {
		s = &ImsVxFanPrepareLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Val1,
			&s.Val2,
			&s.Val3,
			&s.Category,
			&s.CategoryMsg,
			&s.OrderId,
			&s.OrderNo,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanPrepareLogScanAll(rows *sql.Rows) (ss []*ImsVxFanPrepareLog, err error) {
	for rows.Next() {
		s := &ImsVxFanPrepareLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Val1,
			&s.Val2,
			&s.Val3,
			&s.Category,
			&s.CategoryMsg,
			&s.OrderId,
			&s.OrderNo,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanRightCardScan(rows *sql.Rows) (s *ImsVxFanRightCard, err error) {
	if rows.Next() {
		s = &ImsVxFanRightCard{}
		err = rows.Scan(
			&s.Id,
			&s.AccountId,
			&s.Name,
			&s.Money,
			&s.Avatar,
			&s.Title,
			&s.Details,
			&s.Multiple,
			&s.Vip,
			&s.Discount,
			&s.Serial,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanRightCardScanAll(rows *sql.Rows) (ss []*ImsVxFanRightCard, err error) {
	for rows.Next() {
		s := &ImsVxFanRightCard{}
		err = rows.Scan(
			&s.Id,
			&s.AccountId,
			&s.Name,
			&s.Money,
			&s.Avatar,
			&s.Title,
			&s.Details,
			&s.Multiple,
			&s.Vip,
			&s.Discount,
			&s.Serial,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanRightCardUpgradeScan(rows *sql.Rows) (s *ImsVxFanRightCardUpgrade, err error) {
	if rows.Next() {
		s = &ImsVxFanRightCardUpgrade{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.UpgradeRatio,
			&s.FromUid,
			&s.RightCardAmount,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
			&s.Cards,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanRightCardUpgradeScanAll(rows *sql.Rows) (ss []*ImsVxFanRightCardUpgrade, err error) {
	for rows.Next() {
		s := &ImsVxFanRightCardUpgrade{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.UpgradeRatio,
			&s.FromUid,
			&s.RightCardAmount,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
			&s.Cards,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanRightCardUsedScan(rows *sql.Rows) (s *ImsVxFanRightCardUsed, err error) {
	if rows.Next() {
		s = &ImsVxFanRightCardUsed{}
		err = rows.Scan(
			&s.Id,
			&s.Cid,
			&s.Uid,
			&s.Serial,
			&s.Name,
			&s.Money,
			&s.Multiple,
			&s.Amount,
			&s.Days,
			&s.AmountRemain,
			&s.DaysRemain,
			&s.DayRatio,
			&s.DayMoney,
			&s.AmountAchieved,
			&s.AmountAchievedDays,
			&s.Discount,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
			&s.Tips,
			&s.Card,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanRightCardUsedScanAll(rows *sql.Rows) (ss []*ImsVxFanRightCardUsed, err error) {
	for rows.Next() {
		s := &ImsVxFanRightCardUsed{}
		err = rows.Scan(
			&s.Id,
			&s.Cid,
			&s.Uid,
			&s.Serial,
			&s.Name,
			&s.Money,
			&s.Multiple,
			&s.Amount,
			&s.Days,
			&s.AmountRemain,
			&s.DaysRemain,
			&s.DayRatio,
			&s.DayMoney,
			&s.AmountAchieved,
			&s.AmountAchievedDays,
			&s.Discount,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
			&s.Tips,
			&s.Card,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanRightCardUsingScan(rows *sql.Rows) (s *ImsVxFanRightCardUsing, err error) {
	if rows.Next() {
		s = &ImsVxFanRightCardUsing{}
		err = rows.Scan(
			&s.Id,
			&s.Cid,
			&s.Uid,
			&s.Serial,
			&s.Name,
			&s.Money,
			&s.Multiple,
			&s.Amount,
			&s.Days,
			&s.AmountRemain,
			&s.DaysRemain,
			&s.DayRatio,
			&s.DayMoney,
			&s.AmountAchieved,
			&s.AmountAchievedDays,
			&s.Discount,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
			&s.Tips,
			&s.Card,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanRightCardUsingScanAll(rows *sql.Rows) (ss []*ImsVxFanRightCardUsing, err error) {
	for rows.Next() {
		s := &ImsVxFanRightCardUsing{}
		err = rows.Scan(
			&s.Id,
			&s.Cid,
			&s.Uid,
			&s.Serial,
			&s.Name,
			&s.Money,
			&s.Multiple,
			&s.Amount,
			&s.Days,
			&s.AmountRemain,
			&s.DaysRemain,
			&s.DayRatio,
			&s.DayMoney,
			&s.AmountAchieved,
			&s.AmountAchievedDays,
			&s.Discount,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
			&s.Tips,
			&s.Card,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanSignScan(rows *sql.Rows) (s *ImsVxFanSign, err error) {
	if rows.Next() {
		s = &ImsVxFanSign{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Val1,
			&s.Val2,
			&s.Val3,
			&s.Day,
			&s.ContinuedDays,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanSignScanAll(rows *sql.Rows) (ss []*ImsVxFanSign, err error) {
	for rows.Next() {
		s := &ImsVxFanSign{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Val1,
			&s.Val2,
			&s.Val3,
			&s.Day,
			&s.ContinuedDays,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanSignAdvertisingScan(rows *sql.Rows) (s *ImsVxFanSignAdvertising, err error) {
	if rows.Next() {
		s = &ImsVxFanSignAdvertising{}
		err = rows.Scan(
			&s.Id,
			&s.AccountId,
			&s.Name,
			&s.Avatar,
			&s.Video,
			&s.Details,
			&s.Duration,
			&s.Category,
			&s.CategoryMsg,
			&s.Uid,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanSignAdvertisingScanAll(rows *sql.Rows) (ss []*ImsVxFanSignAdvertising, err error) {
	for rows.Next() {
		s := &ImsVxFanSignAdvertising{}
		err = rows.Scan(
			&s.Id,
			&s.AccountId,
			&s.Name,
			&s.Avatar,
			&s.Video,
			&s.Details,
			&s.Duration,
			&s.Category,
			&s.CategoryMsg,
			&s.Uid,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanVipScan(rows *sql.Rows) (s *ImsVxFanVip, err error) {
	if rows.Next() {
		s = &ImsVxFanVip{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Vip,
			&s.VipMsg,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.ExpiredAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanVipScanAll(rows *sql.Rows) (ss []*ImsVxFanVip, err error) {
	for rows.Next() {
		s := &ImsVxFanVip{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Vip,
			&s.VipMsg,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.ExpiredAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxFanWithdrawLogScan(rows *sql.Rows) (s *ImsVxFanWithdrawLog, err error) {
	if rows.Next() {
		s = &ImsVxFanWithdrawLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Val1,
			&s.Val2,
			&s.Val3,
			&s.Category,
			&s.CategoryMsg,
			&s.OrderId,
			&s.OrderNo,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxFanWithdrawLogScanAll(rows *sql.Rows) (ss []*ImsVxFanWithdrawLog, err error) {
	for rows.Next() {
		s := &ImsVxFanWithdrawLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.Val1,
			&s.Val2,
			&s.Val3,
			&s.Category,
			&s.CategoryMsg,
			&s.OrderId,
			&s.OrderNo,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxRedBagScan(rows *sql.Rows) (s *ImsVxRedBag, err error) {
	if rows.Next() {
		s = &ImsVxRedBag{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.AccountId,
			&s.AccountAppId,
			&s.MchId,
			&s.Openid,
			&s.Amount,
			&s.SendListId,
			&s.MchNo,
			&s.ReturnCode,
			&s.ReturnMsg,
			&s.ResultCode,
			&s.ErrCode,
			&s.ErrCodeDes,
			&s.Day,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxRedBagScanAll(rows *sql.Rows) (ss []*ImsVxRedBag, err error) {
	for rows.Next() {
		s := &ImsVxRedBag{}
		err = rows.Scan(
			&s.Id,
			&s.Uid,
			&s.AccountId,
			&s.AccountAppId,
			&s.MchId,
			&s.Openid,
			&s.Amount,
			&s.SendListId,
			&s.MchNo,
			&s.ReturnCode,
			&s.ReturnMsg,
			&s.ResultCode,
			&s.ErrCode,
			&s.ErrCodeDes,
			&s.Day,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsVxRedBagQuotaCardScan(rows *sql.Rows) (s *ImsVxRedBagQuotaCard, err error) {
	if rows.Next() {
		s = &ImsVxRedBagQuotaCard{}
		err = rows.Scan(
			&s.Id,
			&s.AccountId,
			&s.Name,
			&s.Avatar,
			&s.Title,
			&s.Details,
			&s.Money,
			&s.Prepare,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsVxRedBagQuotaCardScanAll(rows *sql.Rows) (ss []*ImsVxRedBagQuotaCard, err error) {
	for rows.Next() {
		s := &ImsVxRedBagQuotaCard{}
		err = rows.Scan(
			&s.Id,
			&s.AccountId,
			&s.Name,
			&s.Avatar,
			&s.Title,
			&s.Details,
			&s.Money,
			&s.Prepare,
			&s.State,
			&s.StateMsg,
			&s.CreatedAt,
			&s.UpdatedAt,
			&s.Note,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsWechatAttachmentScan(rows *sql.Rows) (s *ImsWechatAttachment, err error) {
	if rows.Next() {
		s = &ImsWechatAttachment{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Uid,
			&s.Filename,
			&s.Attachment,
			&s.MediaId,
			&s.Width,
			&s.Height,
			&s.Type,
			&s.Model,
			&s.Tag,
			&s.Createtime,
			&s.ModuleUploadDir,
			&s.GroupId,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsWechatAttachmentScanAll(rows *sql.Rows) (ss []*ImsWechatAttachment, err error) {
	for rows.Next() {
		s := &ImsWechatAttachment{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Acid,
			&s.Uid,
			&s.Filename,
			&s.Attachment,
			&s.MediaId,
			&s.Width,
			&s.Height,
			&s.Type,
			&s.Model,
			&s.Tag,
			&s.Createtime,
			&s.ModuleUploadDir,
			&s.GroupId,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsWechatNewsScan(rows *sql.Rows) (s *ImsWechatNews, err error) {
	if rows.Next() {
		s = &ImsWechatNews{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.AttachId,
			&s.ThumbMediaId,
			&s.ThumbUrl,
			&s.Title,
			&s.Author,
			&s.Digest,
			&s.Content,
			&s.ContentSourceUrl,
			&s.ShowCoverPic,
			&s.Url,
			&s.Displayorder,
			&s.NeedOpenComment,
			&s.OnlyFansCanComment,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsWechatNewsScanAll(rows *sql.Rows) (ss []*ImsWechatNews, err error) {
	for rows.Next() {
		s := &ImsWechatNews{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.AttachId,
			&s.ThumbMediaId,
			&s.ThumbUrl,
			&s.Title,
			&s.Author,
			&s.Digest,
			&s.Content,
			&s.ContentSourceUrl,
			&s.ShowCoverPic,
			&s.Url,
			&s.Displayorder,
			&s.NeedOpenComment,
			&s.OnlyFansCanComment,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsWxappGeneralAnalysisScan(rows *sql.Rows) (s *ImsWxappGeneralAnalysis, err error) {
	if rows.Next() {
		s = &ImsWxappGeneralAnalysis{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.SessionCnt,
			&s.VisitPv,
			&s.VisitUv,
			&s.VisitUvNew,
			&s.Type,
			&s.StayTimeUv,
			&s.StayTimeSession,
			&s.VisitDepth,
			&s.RefDate,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsWxappGeneralAnalysisScanAll(rows *sql.Rows) (ss []*ImsWxappGeneralAnalysis, err error) {
	for rows.Next() {
		s := &ImsWxappGeneralAnalysis{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.SessionCnt,
			&s.VisitPv,
			&s.VisitUv,
			&s.VisitUvNew,
			&s.Type,
			&s.StayTimeUv,
			&s.StayTimeSession,
			&s.VisitDepth,
			&s.RefDate,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsWxappRegisterVersionScan(rows *sql.Rows) (s *ImsWxappRegisterVersion, err error) {
	if rows.Next() {
		s = &ImsWxappRegisterVersion{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.VersionId,
			&s.Auditid,
			&s.Version,
			&s.Description,
			&s.Status,
			&s.Reason,
			&s.UploadTime,
			&s.AuditInfo,
			&s.SubmitInfo,
			&s.Developer,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsWxappRegisterVersionScanAll(rows *sql.Rows) (ss []*ImsWxappRegisterVersion, err error) {
	for rows.Next() {
		s := &ImsWxappRegisterVersion{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.VersionId,
			&s.Auditid,
			&s.Version,
			&s.Description,
			&s.Status,
			&s.Reason,
			&s.UploadTime,
			&s.AuditInfo,
			&s.SubmitInfo,
			&s.Developer,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsWxappUndocodeauditLogScan(rows *sql.Rows) (s *ImsWxappUndocodeauditLog, err error) {
	if rows.Next() {
		s = &ImsWxappUndocodeauditLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.VersionId,
			&s.Auditid,
			&s.RevokeTime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsWxappUndocodeauditLogScanAll(rows *sql.Rows) (ss []*ImsWxappUndocodeauditLog, err error) {
	for rows.Next() {
		s := &ImsWxappUndocodeauditLog{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.VersionId,
			&s.Auditid,
			&s.RevokeTime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsWxappVersionsScan(rows *sql.Rows) (s *ImsWxappVersions, err error) {
	if rows.Next() {
		s = &ImsWxappVersions{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Version,
			&s.Description,
			&s.Modules,
			&s.DesignMethod,
			&s.Template,
			&s.Quickmenu,
			&s.Createtime,
			&s.Type,
			&s.EntryId,
			&s.Appjson,
			&s.DefaultAppjson,
			&s.UseDefault,
			&s.LastModules,
			&s.Tominiprogram,
			&s.UploadTime,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsWxappVersionsScanAll(rows *sql.Rows) (ss []*ImsWxappVersions, err error) {
	for rows.Next() {
		s := &ImsWxappVersions{}
		err = rows.Scan(
			&s.Id,
			&s.Uniacid,
			&s.Multiid,
			&s.Version,
			&s.Description,
			&s.Modules,
			&s.DesignMethod,
			&s.Template,
			&s.Quickmenu,
			&s.Createtime,
			&s.Type,
			&s.EntryId,
			&s.Appjson,
			&s.DefaultAppjson,
			&s.UseDefault,
			&s.LastModules,
			&s.Tominiprogram,
			&s.UploadTime,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

func ImsWxcardReplyScan(rows *sql.Rows) (s *ImsWxcardReply, err error) {
	if rows.Next() {
		s = &ImsWxcardReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.CardId,
			&s.Cid,
			&s.BrandName,
			&s.LogoUrl,
			&s.Success,
			&s.Error,
		)
		if err != nil {
			return
		}
	}
	return
}

func ImsWxcardReplyScanAll(rows *sql.Rows) (ss []*ImsWxcardReply, err error) {
	for rows.Next() {
		s := &ImsWxcardReply{}
		err = rows.Scan(
			&s.Id,
			&s.Rid,
			&s.Title,
			&s.CardId,
			&s.Cid,
			&s.BrandName,
			&s.LogoUrl,
			&s.Success,
			&s.Error,
		)
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}

// ImsAccount ims_account
type ImsAccount struct {
	Acid                    uint   `json:"acid"`                       //
	Uniacid                 uint   `json:"uniacid"`                    //
	Hash                    string `json:"hash"`                       //
	Type                    uint8  `json:"type"`                       //
	Isconnect               int8   `json:"isconnect"`                  //
	Isdeleted               uint8  `json:"isdeleted"`                  //
	Endtime                 uint   `json:"endtime"`                    //
	SendAccountExpireStatus int8   `json:"send_account_expire_status"` //
	SendApiExpireStatus     int8   `json:"send_api_expire_status"`     //
}

// ImsAccountAliapp ims_account_aliapp
type ImsAccountAliapp struct {
	Acid        int    `json:"acid"`        //
	Uniacid     int    `json:"uniacid"`     //
	Level       uint8  `json:"level"`       //
	Name        string `json:"name"`        //
	Description string `json:"description"` //
	Key         string `json:"key"`         //
}

// ImsAccountBaiduapp ims_account_baiduapp
type ImsAccountBaiduapp struct {
	Acid        int    `json:"acid"`        //
	Uniacid     int    `json:"uniacid"`     //
	Name        string `json:"name"`        //
	Appid       string `json:"appid"`       //
	Key         string `json:"key"`         //
	Secret      string `json:"secret"`      //
	Description string `json:"description"` //
}

// ImsAccountPhoneapp ims_account_phoneapp
type ImsAccountPhoneapp struct {
	Acid    int    `json:"acid"`    //
	Uniacid int    `json:"uniacid"` //
	Name    string `json:"name"`    //
}

// ImsAccountToutiaoapp ims_account_toutiaoapp
type ImsAccountToutiaoapp struct {
	Acid        int    `json:"acid"`        //
	Uniacid     int    `json:"uniacid"`     //
	Name        string `json:"name"`        //
	Appid       string `json:"appid"`       //
	Key         string `json:"key"`         //
	Secret      string `json:"secret"`      //
	Description string `json:"description"` //
}

// ImsAccountWebapp ims_account_webapp
type ImsAccountWebapp struct {
	Acid    int     `json:"acid"`    //
	Uniacid *int    `json:"uniacid"` //
	Name    *string `json:"name"`    //
}

// ImsAccountWechats ims_account_wechats
type ImsAccountWechats struct {
	Acid             uint   `json:"acid"`               //
	Uniacid          uint   `json:"uniacid"`            //
	Token            string `json:"token"`              //
	Encodingaeskey   string `json:"encodingaeskey"`     //
	Level            uint8  `json:"level"`              //
	Name             string `json:"name"`               //
	Account          string `json:"account"`            //
	Original         string `json:"original"`           //
	Signature        string `json:"signature"`          //
	Country          string `json:"country"`            //
	Province         string `json:"province"`           //
	City             string `json:"city"`               //
	Username         string `json:"username"`           //
	Password         string `json:"password"`           //
	Lastupdate       uint   `json:"lastupdate"`         //
	Key              string `json:"key"`                //
	Secret           string `json:"secret"`             //
	Styleid          uint   `json:"styleid"`            //
	Subscribeurl     string `json:"subscribeurl"`       //
	AuthRefreshToken string `json:"auth_refresh_token"` //
}

// ImsAccountWxapp ims_account_wxapp
type ImsAccountWxapp struct {
	Acid             uint    `json:"acid"`               //
	Uniacid          int     `json:"uniacid"`            //
	Token            string  `json:"token"`              //
	Encodingaeskey   string  `json:"encodingaeskey"`     //
	Level            int8    `json:"level"`              //
	Account          string  `json:"account"`            //
	Original         string  `json:"original"`           //
	Key              string  `json:"key"`                //
	Secret           string  `json:"secret"`             //
	Name             string  `json:"name"`               //
	Appdomain        string  `json:"appdomain"`          //
	AuthRefreshToken *string `json:"auth_refresh_token"` //
}

// ImsAccountXzapp ims_account_xzapp
type ImsAccountXzapp struct {
	Acid           int    `json:"acid"`           //
	Uniacid        int    `json:"uniacid"`        //
	Name           string `json:"name"`           //
	Original       string `json:"original"`       //
	Lastupdate     int    `json:"lastupdate"`     //
	Styleid        int    `json:"styleid"`        //
	Createtime     int    `json:"createtime"`     //
	Token          string `json:"token"`          //
	Encodingaeskey string `json:"encodingaeskey"` //
	XzappId        string `json:"xzapp_id"`       //
	Level          uint8  `json:"level"`          //
	Key            string `json:"key"`            //
	Secret         string `json:"secret"`         //
}

// ImsActivityClerks ims_activity_clerks
type ImsActivityClerks struct {
	Id       uint   `json:"id"`       //
	Uniacid  uint   `json:"uniacid"`  //
	Uid      uint   `json:"uid"`      //
	Storeid  uint   `json:"storeid"`  //
	Name     string `json:"name"`     //
	Password string `json:"password"` //
	Mobile   string `json:"mobile"`   //
	Openid   string `json:"openid"`   //
	Nickname string `json:"nickname"` //
}

// ImsActivityClerkMenu ims_activity_clerk_menu
type ImsActivityClerkMenu struct {
	Id           int    `json:"id"`           //
	Uniacid      int    `json:"uniacid"`      //
	Displayorder int    `json:"displayorder"` //
	Pid          int    `json:"pid"`          //
	GroupName    string `json:"group_name"`   //
	Title        string `json:"title"`        //
	Icon         string `json:"icon"`         //
	Url          string `json:"url"`          //
	Type         string `json:"type"`         //
	Permission   string `json:"permission"`   //
	System       int    `json:"system"`       //
}

// ImsAddress ims_address 系统地址配置表
type ImsAddress struct {
	Id    int64  `json:"id"`    // vv
	Pid   int64  `json:"pid"`   // 上级
	Name  string `json:"name"`  // 地区名称
	Level int64  `json:"level"` // 地区等级 分省市县区
}

// ImsArticleCategory ims_article_category
type ImsArticleCategory struct {
	Id           uint   `json:"id"`           //
	Title        string `json:"title"`        //
	Displayorder uint8  `json:"displayorder"` //
	Type         string `json:"type"`         //
}

// ImsArticleComment ims_article_comment
type ImsArticleComment struct {
	Id         int     `json:"id"`         //
	Articleid  int     `json:"articleid"`  //
	Parentid   int     `json:"parentid"`   //
	Uid        int     `json:"uid"`        //
	Content    *string `json:"content"`    //
	IsLike     int8    `json:"is_like"`    //
	IsReply    int8    `json:"is_reply"`   //
	LikeNum    int     `json:"like_num"`   //
	Createtime int     `json:"createtime"` //
}

// ImsArticleNews ims_article_news
type ImsArticleNews struct {
	Id           uint   `json:"id"`           //
	Cateid       uint   `json:"cateid"`       //
	Title        string `json:"title"`        //
	Content      string `json:"content"`      //
	Thumb        string `json:"thumb"`        //
	Source       string `json:"source"`       //
	Author       string `json:"author"`       //
	Displayorder uint8  `json:"displayorder"` //
	IsDisplay    uint8  `json:"is_display"`   //
	IsShowHome   uint8  `json:"is_show_home"` //
	Createtime   uint   `json:"createtime"`   //
	Click        uint   `json:"click"`        //
}

// ImsArticleNotice ims_article_notice
type ImsArticleNotice struct {
	Id           uint   `json:"id"`           //
	Cateid       uint   `json:"cateid"`       //
	Title        string `json:"title"`        //
	Content      string `json:"content"`      //
	Displayorder uint8  `json:"displayorder"` //
	IsDisplay    uint8  `json:"is_display"`   //
	IsShowHome   uint8  `json:"is_show_home"` //
	Createtime   uint   `json:"createtime"`   //
	Click        uint   `json:"click"`        //
	Style        string `json:"style"`        //
	Group        string `json:"group"`        //
}

// ImsArticleUnreadNotice ims_article_unread_notice
type ImsArticleUnreadNotice struct {
	Id       uint  `json:"id"`        //
	NoticeId uint  `json:"notice_id"` //
	Uid      uint  `json:"uid"`       //
	IsNew    uint8 `json:"is_new"`    //
}

// ImsAttachmentGroup ims_attachment_group
type ImsAttachmentGroup struct {
	Id      int    `json:"id"`      //
	Pid     int    `json:"pid"`     //
	Name    string `json:"name"`    //
	Uniacid int    `json:"uniacid"` //
	Uid     int    `json:"uid"`     //
	Type    int8   `json:"type"`    //
}

// ImsBasicReply ims_basic_reply
type ImsBasicReply struct {
	Id      uint   `json:"id"`      //
	Rid     uint   `json:"rid"`     //
	Content string `json:"content"` //
}

// ImsBusiness ims_business
type ImsBusiness struct {
	Id         uint   `json:"id"`         //
	Weid       uint   `json:"weid"`       //
	Title      string `json:"title"`      //
	Thumb      string `json:"thumb"`      //
	Content    string `json:"content"`    //
	Phone      string `json:"phone"`      //
	Qq         string `json:"qq"`         //
	Province   string `json:"province"`   //
	City       string `json:"city"`       //
	Dist       string `json:"dist"`       //
	Address    string `json:"address"`    //
	Lng        string `json:"lng"`        //
	Lat        string `json:"lat"`        //
	Industry1  string `json:"industry1"`  //
	Industry2  string `json:"industry2"`  //
	Createtime int    `json:"createtime"` //
}

// ImsCoreAttachment ims_core_attachment
type ImsCoreAttachment struct {
	Id              uint   `json:"id"`                //
	Uniacid         uint   `json:"uniacid"`           //
	Uid             uint   `json:"uid"`               //
	Filename        string `json:"filename"`          //
	Attachment      string `json:"attachment"`        //
	Type            uint8  `json:"type"`              //
	Createtime      uint   `json:"createtime"`        //
	ModuleUploadDir string `json:"module_upload_dir"` //
	GroupId         int    `json:"group_id"`          //
	Displayorder    int    `json:"displayorder"`      //
}

// ImsCoreCache ims_core_cache
type ImsCoreCache struct {
	Key   string `json:"key"`   //
	Value string `json:"value"` //
}

// ImsCoreCron ims_core_cron
type ImsCoreCron struct {
	Id          uint   `json:"id"`          //
	Cloudid     uint   `json:"cloudid"`     //
	Module      string `json:"module"`      //
	Uniacid     uint   `json:"uniacid"`     //
	Type        uint8  `json:"type"`        //
	Name        string `json:"name"`        //
	Filename    string `json:"filename"`    //
	Lastruntime uint   `json:"lastruntime"` //
	Nextruntime uint   `json:"nextruntime"` //
	Weekday     int8   `json:"weekday"`     //
	Day         int8   `json:"day"`         //
	Hour        int8   `json:"hour"`        //
	Minute      string `json:"minute"`      //
	Extra       string `json:"extra"`       //
	Status      uint8  `json:"status"`      //
	Createtime  uint   `json:"createtime"`  //
}

// ImsCoreCronRecord ims_core_cron_record
type ImsCoreCronRecord struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Module     string `json:"module"`     //
	Type       string `json:"type"`       //
	Tid        uint   `json:"tid"`        //
	Note       string `json:"note"`       //
	Tag        string `json:"tag"`        //
	Createtime uint   `json:"createtime"` //
}

// ImsCoreJob ims_core_job
type ImsCoreJob struct {
	Id         int    `json:"id"`         //
	Type       int8   `json:"type"`       //
	Uniacid    int    `json:"uniacid"`    //
	Payload    string `json:"payload"`    //
	Status     int8   `json:"status"`     //
	Title      string `json:"title"`      //
	Handled    int    `json:"handled"`    //
	Total      int    `json:"total"`      //
	Createtime int    `json:"createtime"` //
	Updatetime int    `json:"updatetime"` //
	Endtime    int    `json:"endtime"`    //
	Uid        int    `json:"uid"`        //
	Isdeleted  *int8  `json:"isdeleted"`  //
}

// ImsCoreMenu ims_core_menu
type ImsCoreMenu struct {
	Id             uint   `json:"id"`              //
	Pid            uint   `json:"pid"`             //
	Title          string `json:"title"`           //
	Name           string `json:"name"`            //
	Url            string `json:"url"`             //
	AppendTitle    string `json:"append_title"`    //
	AppendUrl      string `json:"append_url"`      //
	Displayorder   uint8  `json:"displayorder"`    //
	Type           string `json:"type"`            //
	IsDisplay      uint8  `json:"is_display"`      //
	IsSystem       uint8  `json:"is_system"`       //
	PermissionName string `json:"permission_name"` //
	GroupName      string `json:"group_name"`      //
	Icon           string `json:"icon"`            //
}

// ImsCoreMenuShortcut ims_core_menu_shortcut
type ImsCoreMenuShortcut struct {
	Id           int    `json:"id"`           //
	Uid          int    `json:"uid"`          //
	Uniacid      int    `json:"uniacid"`      //
	Modulename   string `json:"modulename"`   //
	Displayorder int    `json:"displayorder"` //
	Position     string `json:"position"`     //
	Updatetime   int    `json:"updatetime"`   //
}

// ImsCorePaylog ims_core_paylog
type ImsCorePaylog struct {
	Plid        uint64  `json:"plid"`         //
	Type        string  `json:"type"`         //
	Uniacid     int     `json:"uniacid"`      //
	Acid        int     `json:"acid"`         //
	Openid      string  `json:"openid"`       //
	Uniontid    string  `json:"uniontid"`     //
	Tid         string  `json:"tid"`          //
	Fee         float64 `json:"fee"`          //
	Status      int8    `json:"status"`       //
	Module      string  `json:"module"`       //
	Tag         string  `json:"tag"`          //
	IsUsecard   uint8   `json:"is_usecard"`   //
	CardType    uint8   `json:"card_type"`    //
	CardId      string  `json:"card_id"`      //
	CardFee     float64 `json:"card_fee"`     //
	EncryptCode string  `json:"encrypt_code"` //
	IsWish      int8    `json:"is_wish"`      //
}

// ImsCorePerformance ims_core_performance
type ImsCorePerformance struct {
	Id         uint   `json:"id"`         //
	Type       int8   `json:"type"`       //
	Runtime    string `json:"runtime"`    //
	Runurl     string `json:"runurl"`     //
	Runsql     string `json:"runsql"`     //
	Createtime int    `json:"createtime"` //
}

// ImsCoreQueue ims_core_queue
type ImsCoreQueue struct {
	Qid      uint64 `json:"qid"`      //
	Uniacid  uint   `json:"uniacid"`  //
	Acid     uint   `json:"acid"`     //
	Message  string `json:"message"`  //
	Params   string `json:"params"`   //
	Keyword  string `json:"keyword"`  //
	Response string `json:"response"` //
	Module   string `json:"module"`   //
	Type     uint8  `json:"type"`     //
	Dateline uint   `json:"dateline"` //
}

// ImsCoreRefundlog ims_core_refundlog
type ImsCoreRefundlog struct {
	Id             int     `json:"id"`              //
	Uniacid        int     `json:"uniacid"`         //
	RefundUniontid string  `json:"refund_uniontid"` //
	Reason         string  `json:"reason"`          //
	Uniontid       string  `json:"uniontid"`        //
	Fee            float64 `json:"fee"`             //
	Status         int     `json:"status"`          //
	IsWish         int8    `json:"is_wish"`         //
}

// ImsCoreResource ims_core_resource
type ImsCoreResource struct {
	Mid      int    `json:"mid"`      //
	Uniacid  uint   `json:"uniacid"`  //
	MediaId  string `json:"media_id"` //
	Trunk    uint   `json:"trunk"`    //
	Type     string `json:"type"`     //
	Dateline uint   `json:"dateline"` //
}

// ImsCoreSendsmsLog ims_core_sendsms_log
type ImsCoreSendsmsLog struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Mobile     string `json:"mobile"`     //
	Content    string `json:"content"`    //
	Result     string `json:"result"`     //
	Createtime uint   `json:"createtime"` //
}

// ImsCoreSessions ims_core_sessions
type ImsCoreSessions struct {
	Sid        string `json:"sid"`        //
	Uniacid    uint   `json:"uniacid"`    //
	Openid     string `json:"openid"`     //
	Data       string `json:"data"`       //
	Expiretime uint   `json:"expiretime"` //
}

// ImsCoreSettings ims_core_settings
type ImsCoreSettings struct {
	Key   string `json:"key"`   //
	Value string `json:"value"` //
}

// ImsCouponLocation ims_coupon_location
type ImsCouponLocation struct {
	Id           uint   `json:"id"`            //
	Uniacid      uint   `json:"uniacid"`       //
	Acid         uint   `json:"acid"`          //
	Sid          uint   `json:"sid"`           //
	LocationId   uint   `json:"location_id"`   //
	BusinessName string `json:"business_name"` //
	BranchName   string `json:"branch_name"`   //
	Category     string `json:"category"`      //
	Province     string `json:"province"`      //
	City         string `json:"city"`          //
	District     string `json:"district"`      //
	Address      string `json:"address"`       //
	Longitude    string `json:"longitude"`     //
	Latitude     string `json:"latitude"`      //
	Telephone    string `json:"telephone"`     //
	PhotoList    string `json:"photo_list"`    //
	AvgPrice     uint   `json:"avg_price"`     //
	OpenTime     string `json:"open_time"`     //
	Recommend    string `json:"recommend"`     //
	Special      string `json:"special"`       //
	Introduction string `json:"introduction"`  //
	OffsetType   uint8  `json:"offset_type"`   //
	Status       uint8  `json:"status"`        //
	Message      string `json:"message"`       //
}

// ImsCoverReply ims_cover_reply
type ImsCoverReply struct {
	Id          uint   `json:"id"`          //
	Uniacid     uint   `json:"uniacid"`     //
	Multiid     uint   `json:"multiid"`     //
	Rid         uint   `json:"rid"`         //
	Module      string `json:"module"`      //
	Do          string `json:"do"`          //
	Title       string `json:"title"`       //
	Description string `json:"description"` //
	Thumb       string `json:"thumb"`       //
	Url         string `json:"url"`         //
}

// ImsCustomReply ims_custom_reply
type ImsCustomReply struct {
	Id     uint `json:"id"`     //
	Rid    uint `json:"rid"`    //
	Start1 int  `json:"start1"` //
	End1   int  `json:"end1"`   //
	Start2 int  `json:"start2"` //
	End2   int  `json:"end2"`   //
}

// ImsImagesReply ims_images_reply
type ImsImagesReply struct {
	Id          uint   `json:"id"`          //
	Rid         uint   `json:"rid"`         //
	Title       string `json:"title"`       //
	Description string `json:"description"` //
	Mediaid     string `json:"mediaid"`     //
	Createtime  int    `json:"createtime"`  //
}

// ImsMcCashRecord ims_mc_cash_record
type ImsMcCashRecord struct {
	Id         uint    `json:"id"`          //
	Uniacid    uint    `json:"uniacid"`     //
	Uid        uint    `json:"uid"`         //
	ClerkId    uint    `json:"clerk_id"`    //
	StoreId    uint    `json:"store_id"`    //
	ClerkType  uint8   `json:"clerk_type"`  //
	Fee        float64 `json:"fee"`         //
	FinalFee   float64 `json:"final_fee"`   //
	Credit1    uint    `json:"credit1"`     //
	Credit1Fee float64 `json:"credit1_fee"` //
	Credit2    float64 `json:"credit2"`     //
	Cash       float64 `json:"cash"`        //
	ReturnCash float64 `json:"return_cash"` //
	FinalCash  float64 `json:"final_cash"`  //
	Remark     string  `json:"remark"`      //
	Createtime uint    `json:"createtime"`  //
	TradeType  string  `json:"trade_type"`  //
}

// ImsMcChatsRecord ims_mc_chats_record
type ImsMcChatsRecord struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Acid       uint   `json:"acid"`       //
	Flag       uint8  `json:"flag"`       //
	Openid     string `json:"openid"`     //
	Msgtype    string `json:"msgtype"`    //
	Content    string `json:"content"`    //
	Createtime uint   `json:"createtime"` //
}

// ImsMcCreditsRecharge ims_mc_credits_recharge
type ImsMcCreditsRecharge struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Uid        uint   `json:"uid"`        //
	Openid     string `json:"openid"`     //
	Tid        string `json:"tid"`        //
	Transid    string `json:"transid"`    //
	Fee        string `json:"fee"`        //
	Type       string `json:"type"`       //
	Tag        string `json:"tag"`        //
	Status     int8   `json:"status"`     //
	Createtime uint   `json:"createtime"` //
	Backtype   uint8  `json:"backtype"`   //
}

// ImsMcCreditsRecord ims_mc_credits_record
type ImsMcCreditsRecord struct {
	Id          int     `json:"id"`           //
	Uid         uint    `json:"uid"`          //
	Uniacid     int     `json:"uniacid"`      //
	Credittype  string  `json:"credittype"`   //
	Num         float64 `json:"num"`          //
	Operator    uint    `json:"operator"`     //
	Module      string  `json:"module"`       //
	ClerkId     uint    `json:"clerk_id"`     //
	StoreId     uint    `json:"store_id"`     //
	ClerkType   uint8   `json:"clerk_type"`   //
	Createtime  uint    `json:"createtime"`   //
	Remark      string  `json:"remark"`       //
	RealUniacid int     `json:"real_uniacid"` //
}

// ImsMcFansGroups ims_mc_fans_groups
type ImsMcFansGroups struct {
	Id      uint   `json:"id"`      //
	Uniacid uint   `json:"uniacid"` //
	Acid    uint   `json:"acid"`    //
	Groups  string `json:"groups"`  //
}

// ImsMcFansTag ims_mc_fans_tag
type ImsMcFansTag struct {
	Id             uint    `json:"id"`              //
	Uniacid        *int    `json:"uniacid"`         //
	Fanid          int     `json:"fanid"`           //
	Openid         string  `json:"openid"`          //
	Subscribe      *int    `json:"subscribe"`       //
	Nickname       *string `json:"nickname"`        //
	Sex            *int    `json:"sex"`             //
	Language       *string `json:"language"`        //
	City           *string `json:"city"`            //
	Province       *string `json:"province"`        //
	Country        *string `json:"country"`         //
	Headimgurl     *string `json:"headimgurl"`      //
	SubscribeTime  int     `json:"subscribe_time"`  //
	Unionid        *string `json:"unionid"`         //
	Remark         *string `json:"remark"`          //
	Groupid        *string `json:"groupid"`         //
	TagidList      *string `json:"tagid_list"`      //
	SubscribeScene *string `json:"subscribe_scene"` //
	QrSceneStr     *string `json:"qr_scene_str"`    //
	QrScene        *string `json:"qr_scene"`        //
}

// ImsMcFansTagMapping ims_mc_fans_tag_mapping
type ImsMcFansTagMapping struct {
	Id    uint   `json:"id"`    //
	Fanid uint   `json:"fanid"` //
	Tagid string `json:"tagid"` //
}

// ImsMcGroups ims_mc_groups
type ImsMcGroups struct {
	Groupid   int    `json:"groupid"`   //
	Uniacid   int    `json:"uniacid"`   //
	Title     string `json:"title"`     //
	Credit    uint   `json:"credit"`    //
	Isdefault int8   `json:"isdefault"` //
}

// ImsMcHandsel ims_mc_handsel
type ImsMcHandsel struct {
	Id          uint   `json:"id"`           //
	Uniacid     int    `json:"uniacid"`      //
	Touid       uint   `json:"touid"`        //
	Fromuid     string `json:"fromuid"`      //
	Module      string `json:"module"`       //
	Sign        string `json:"sign"`         //
	Action      string `json:"action"`       //
	CreditValue uint   `json:"credit_value"` //
	Createtime  uint   `json:"createtime"`   //
}

// ImsMcMappingFans ims_mc_mapping_fans
type ImsMcMappingFans struct {
	Fanid        uint   `json:"fanid"`        //
	Acid         uint   `json:"acid"`         //
	Uniacid      uint   `json:"uniacid"`      //
	Uid          uint   `json:"uid"`          //
	Openid       string `json:"openid"`       //
	Nickname     string `json:"nickname"`     //
	Groupid      string `json:"groupid"`      //
	Salt         string `json:"salt"`         //
	Follow       uint8  `json:"follow"`       //
	Followtime   uint   `json:"followtime"`   //
	Unfollowtime uint   `json:"unfollowtime"` //
	Tag          string `json:"tag"`          //
	Updatetime   *uint  `json:"updatetime"`   //
	Unionid      string `json:"unionid"`      //
	UserFrom     int8   `json:"user_from"`    //
}

// ImsMcMassRecord ims_mc_mass_record
type ImsMcMassRecord struct {
	Id            uint   `json:"id"`            //
	Uniacid       uint   `json:"uniacid"`       //
	Acid          uint   `json:"acid"`          //
	Groupname     string `json:"groupname"`     //
	Fansnum       uint   `json:"fansnum"`       //
	Msgtype       string `json:"msgtype"`       //
	Content       string `json:"content"`       //
	Group         int    `json:"group"`         //
	AttachId      uint   `json:"attach_id"`     //
	MediaId       string `json:"media_id"`      //
	Type          uint8  `json:"type"`          //
	Status        uint8  `json:"status"`        //
	CronId        uint   `json:"cron_id"`       //
	Sendtime      uint   `json:"sendtime"`      //
	Finalsendtime uint   `json:"finalsendtime"` //
	Createtime    uint   `json:"createtime"`    //
	MsgId         string `json:"msg_id"`        //
	MsgDataId     string `json:"msg_data_id"`   //
}

// ImsMcMembers ims_mc_members
type ImsMcMembers struct {
	Uid             uint    `json:"uid"`             //
	Uniacid         uint    `json:"uniacid"`         //
	Mobile          string  `json:"mobile"`          //
	Email           string  `json:"email"`           //
	Password        string  `json:"password"`        //
	Salt            string  `json:"salt"`            //
	Groupid         int     `json:"groupid"`         //
	Credit1         float64 `json:"credit1"`         //
	Credit2         float64 `json:"credit2"`         //
	Credit3         float64 `json:"credit3"`         //
	Credit4         float64 `json:"credit4"`         //
	Credit5         float64 `json:"credit5"`         //
	Credit6         float64 `json:"credit6"`         //
	Createtime      uint    `json:"createtime"`      //
	Realname        string  `json:"realname"`        //
	Nickname        string  `json:"nickname"`        //
	Avatar          string  `json:"avatar"`          //
	Qq              string  `json:"qq"`              //
	Vip             uint8   `json:"vip"`             //
	Gender          int8    `json:"gender"`          //
	Birthyear       uint16  `json:"birthyear"`       //
	Birthmonth      uint8   `json:"birthmonth"`      //
	Birthday        uint8   `json:"birthday"`        //
	Constellation   string  `json:"constellation"`   //
	Zodiac          string  `json:"zodiac"`          //
	Telephone       string  `json:"telephone"`       //
	Idcard          string  `json:"idcard"`          //
	Studentid       string  `json:"studentid"`       //
	Grade           string  `json:"grade"`           //
	Address         string  `json:"address"`         //
	Zipcode         string  `json:"zipcode"`         //
	Nationality     string  `json:"nationality"`     //
	Resideprovince  string  `json:"resideprovince"`  //
	Residecity      string  `json:"residecity"`      //
	Residedist      string  `json:"residedist"`      //
	Graduateschool  string  `json:"graduateschool"`  //
	Company         string  `json:"company"`         //
	Education       string  `json:"education"`       //
	Occupation      string  `json:"occupation"`      //
	Position        string  `json:"position"`        //
	Revenue         string  `json:"revenue"`         //
	Affectivestatus string  `json:"affectivestatus"` //
	Lookingfor      string  `json:"lookingfor"`      //
	Bloodtype       string  `json:"bloodtype"`       //
	Height          string  `json:"height"`          //
	Weight          string  `json:"weight"`          //
	Alipay          string  `json:"alipay"`          //
	Msn             string  `json:"msn"`             //
	Taobao          string  `json:"taobao"`          //
	Site            string  `json:"site"`            //
	Bio             string  `json:"bio"`             //
	Interest        string  `json:"interest"`        //
	PayPassword     string  `json:"pay_password"`    //
	UserFrom        int8    `json:"user_from"`       //
}

// ImsMcMemberAddress ims_mc_member_address
type ImsMcMemberAddress struct {
	Id        uint   `json:"id"`        //
	Uniacid   uint   `json:"uniacid"`   //
	Uid       uint   `json:"uid"`       //
	Username  string `json:"username"`  //
	Mobile    string `json:"mobile"`    //
	Zipcode   string `json:"zipcode"`   //
	Province  string `json:"province"`  //
	City      string `json:"city"`      //
	District  string `json:"district"`  //
	Address   string `json:"address"`   //
	Isdefault uint8  `json:"isdefault"` //
}

// ImsMcMemberFields ims_mc_member_fields
type ImsMcMemberFields struct {
	Id           uint   `json:"id"`           //
	Uniacid      int    `json:"uniacid"`      //
	Fieldid      int    `json:"fieldid"`      //
	Title        string `json:"title"`        //
	Available    int8   `json:"available"`    //
	Displayorder int16  `json:"displayorder"` //
}

// ImsMcMemberProperty ims_mc_member_property
type ImsMcMemberProperty struct {
	Id       uint   `json:"id"`       //
	Uniacid  int    `json:"uniacid"`  //
	Property string `json:"property"` //
}

// ImsMcOauthFans ims_mc_oauth_fans
type ImsMcOauthFans struct {
	Id          uint   `json:"id"`           //
	OauthOpenid string `json:"oauth_openid"` //
	Acid        uint   `json:"acid"`         //
	Uid         uint   `json:"uid"`          //
	Openid      string `json:"openid"`       //
}

// ImsMenuEvent ims_menu_event
type ImsMenuEvent struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Keyword    string `json:"keyword"`    //
	Type       string `json:"type"`       //
	Picmd5     string `json:"picmd5"`     //
	Openid     string `json:"openid"`     //
	Createtime uint   `json:"createtime"` //
}

// ImsMessageNoticeLog ims_message_notice_log
type ImsMessageNoticeLog struct {
	Id         int    `json:"id"`          //
	Message    string `json:"message"`     //
	IsRead     int8   `json:"is_read"`     //
	Uid        int    `json:"uid"`         //
	Sign       string `json:"sign"`        //
	Type       int8   `json:"type"`        //
	Status     *int8  `json:"status"`      //
	CreateTime int    `json:"create_time"` //
	EndTime    int    `json:"end_time"`    //
	Url        string `json:"url"`         //
}

// ImsMobilenumber ims_mobilenumber
type ImsMobilenumber struct {
	Id       int   `json:"id"`       //
	Rid      int   `json:"rid"`      //
	Enabled  uint8 `json:"enabled"`  //
	Dateline *int  `json:"dateline"` //
}

// ImsModules ims_modules
type ImsModules struct {
	Mid               uint   `json:"mid"`                //
	Name              string `json:"name"`               //
	ApplicationType   uint8  `json:"application_type"`   //
	Type              string `json:"type"`               //
	Title             string `json:"title"`              //
	Version           string `json:"version"`            //
	Ability           string `json:"ability"`            //
	Description       string `json:"description"`        //
	Author            string `json:"author"`             //
	Url               string `json:"url"`                //
	Settings          int8   `json:"settings"`           //
	Subscribes        string `json:"subscribes"`         //
	Handles           string `json:"handles"`            //
	Isrulefields      int8   `json:"isrulefields"`       //
	Issystem          uint8  `json:"issystem"`           //
	Target            uint   `json:"target"`             //
	Iscard            uint8  `json:"iscard"`             //
	Permissions       string `json:"permissions"`        //
	TitleInitial      string `json:"title_initial"`      //
	WxappSupport      int8   `json:"wxapp_support"`      //
	WelcomeSupport    int    `json:"welcome_support"`    //
	OauthType         int8   `json:"oauth_type"`         //
	WebappSupport     int8   `json:"webapp_support"`     //
	PhoneappSupport   int8   `json:"phoneapp_support"`   //
	AccountSupport    int8   `json:"account_support"`    //
	XzappSupport      int8   `json:"xzapp_support"`      //
	AliappSupport     int8   `json:"aliapp_support"`     //
	Logo              string `json:"logo"`               //
	BaiduappSupport   int8   `json:"baiduapp_support"`   //
	ToutiaoappSupport int8   `json:"toutiaoapp_support"` //
	From              string `json:"from"`               //
	CloudRecord       int8   `json:"cloud_record"`       //
	Sections          uint   `json:"sections"`           //
	Label             string `json:"label"`              //
}

// ImsModulesBindings ims_modules_bindings
type ImsModulesBindings struct {
	Eid          int    `json:"eid"`          //
	Module       string `json:"module"`       //
	Entry        string `json:"entry"`        //
	Call         string `json:"call"`         //
	Title        string `json:"title"`        //
	Do           string `json:"do"`           //
	State        string `json:"state"`        //
	Direct       int    `json:"direct"`       //
	Url          string `json:"url"`          //
	Icon         string `json:"icon"`         //
	Displayorder uint8  `json:"displayorder"` //
	Multilevel   int8   `json:"multilevel"`   //
	Parent       string `json:"parent"`       //
}

// ImsModulesCloud ims_modules_cloud
type ImsModulesCloud struct {
	Id                int    `json:"id"`                 //
	Name              string `json:"name"`               //
	ApplicationType   uint8  `json:"application_type"`   //
	Title             string `json:"title"`              //
	TitleInitial      string `json:"title_initial"`      //
	Logo              string `json:"logo"`               //
	Version           string `json:"version"`            //
	InstallStatus     int8   `json:"install_status"`     //
	AccountSupport    int8   `json:"account_support"`    //
	WxappSupport      int8   `json:"wxapp_support"`      //
	WebappSupport     int8   `json:"webapp_support"`     //
	PhoneappSupport   int8   `json:"phoneapp_support"`   //
	WelcomeSupport    int8   `json:"welcome_support"`    //
	MainModuleName    string `json:"main_module_name"`   //
	MainModuleLogo    string `json:"main_module_logo"`   //
	HasNewVersion     int8   `json:"has_new_version"`    //
	HasNewBranch      int8   `json:"has_new_branch"`     //
	IsBan             int8   `json:"is_ban"`             //
	Lastupdatetime    int    `json:"lastupdatetime"`     //
	XzappSupport      int8   `json:"xzapp_support"`      //
	CloudId           int    `json:"cloud_id"`           //
	AliappSupport     int8   `json:"aliapp_support"`     //
	BaiduappSupport   int8   `json:"baiduapp_support"`   //
	ToutiaoappSupport int8   `json:"toutiaoapp_support"` //
	Buytime           int    `json:"buytime"`            //
	ModuleStatus      int8   `json:"module_status"`      //
	Label             string `json:"label"`              //
}

// ImsModulesIgnore ims_modules_ignore
type ImsModulesIgnore struct {
	Id      int    `json:"id"`      //
	Name    string `json:"name"`    //
	Version string `json:"version"` //
}

// ImsModulesPlugin ims_modules_plugin
type ImsModulesPlugin struct {
	Id         int     `json:"id"`          //
	Name       *string `json:"name"`        //
	MainModule *string `json:"main_module"` //
}

// ImsModulesPluginRank ims_modules_plugin_rank
type ImsModulesPluginRank struct {
	Id             int    `json:"id"`               //
	Uniacid        int    `json:"uniacid"`          //
	Uid            int    `json:"uid"`              //
	Rank           int    `json:"rank"`             //
	PluginName     string `json:"plugin_name"`      //
	MainModuleName string `json:"main_module_name"` //
}

// ImsModulesRank ims_modules_rank
type ImsModulesRank struct {
	Id         uint   `json:"id"`          //
	ModuleName string `json:"module_name"` //
	Uid        int    `json:"uid"`         //
	Rank       int    `json:"rank"`        //
	Uniacid    int    `json:"uniacid"`     //
}

// ImsModulesRecycle ims_modules_recycle
type ImsModulesRecycle struct {
	Id                int    `json:"id"`                 //
	Name              string `json:"name"`               //
	Type              int8   `json:"type"`               //
	AccountSupport    int8   `json:"account_support"`    //
	WxappSupport      int8   `json:"wxapp_support"`      //
	WelcomeSupport    int8   `json:"welcome_support"`    //
	WebappSupport     int8   `json:"webapp_support"`     //
	PhoneappSupport   int8   `json:"phoneapp_support"`   //
	XzappSupport      int8   `json:"xzapp_support"`      //
	AliappSupport     int8   `json:"aliapp_support"`     //
	BaiduappSupport   int8   `json:"baiduapp_support"`   //
	ToutiaoappSupport int8   `json:"toutiaoapp_support"` //
}

// ImsMusicReply ims_music_reply
type ImsMusicReply struct {
	Id          uint   `json:"id"`          //
	Rid         uint   `json:"rid"`         //
	Title       string `json:"title"`       //
	Description string `json:"description"` //
	Url         string `json:"url"`         //
	Hqurl       string `json:"hqurl"`       //
}

// ImsNewsReply ims_news_reply
type ImsNewsReply struct {
	Id           uint   `json:"id"`           //
	Rid          uint   `json:"rid"`          //
	ParentId     int    `json:"parent_id"`    //
	Title        string `json:"title"`        //
	Author       string `json:"author"`       //
	Description  string `json:"description"`  //
	Thumb        string `json:"thumb"`        //
	Content      string `json:"content"`      //
	Url          string `json:"url"`          //
	Displayorder int    `json:"displayorder"` //
	Incontent    int8   `json:"incontent"`    //
	Createtime   int    `json:"createtime"`   //
	MediaId      string `json:"media_id"`     //
}

// ImsPhoneappVersions ims_phoneapp_versions
type ImsPhoneappVersions struct {
	Id          int    `json:"id"`          //
	Uniacid     int    `json:"uniacid"`     //
	Version     string `json:"version"`     //
	Description string `json:"description"` //
	Modules     string `json:"modules"`     //
	Createtime  int    `json:"createtime"`  //
}

// ImsProfileFields ims_profile_fields
type ImsProfileFields struct {
	Id             uint   `json:"id"`             //
	Field          string `json:"field"`          //
	Available      int8   `json:"available"`      //
	Title          string `json:"title"`          //
	Description    string `json:"description"`    //
	Displayorder   int16  `json:"displayorder"`   //
	Required       int8   `json:"required"`       //
	Unchangeable   int8   `json:"unchangeable"`   //
	Showinregister int8   `json:"showinregister"` //
	FieldLength    int    `json:"field_length"`   //
}

// ImsQrcode ims_qrcode
type ImsQrcode struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Acid       uint   `json:"acid"`       //
	Type       string `json:"type"`       //
	Extra      uint   `json:"extra"`      //
	Qrcid      int64  `json:"qrcid"`      //
	SceneStr   string `json:"scene_str"`  //
	Name       string `json:"name"`       //
	Keyword    string `json:"keyword"`    //
	Model      uint8  `json:"model"`      //
	Ticket     string `json:"ticket"`     //
	Url        string `json:"url"`        //
	Expire     uint   `json:"expire"`     //
	Subnum     uint   `json:"subnum"`     //
	Createtime uint   `json:"createtime"` //
	Status     uint8  `json:"status"`     //
}

// ImsQrcodeStat ims_qrcode_stat
type ImsQrcodeStat struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Acid       uint   `json:"acid"`       //
	Qid        uint   `json:"qid"`        //
	Openid     string `json:"openid"`     //
	Type       uint8  `json:"type"`       //
	Qrcid      uint64 `json:"qrcid"`      //
	SceneStr   string `json:"scene_str"`  //
	Name       string `json:"name"`       //
	Createtime uint   `json:"createtime"` //
}

// ImsRule ims_rule
type ImsRule struct {
	Id           uint   `json:"id"`           //
	Uniacid      uint   `json:"uniacid"`      //
	Name         string `json:"name"`         //
	Module       string `json:"module"`       //
	Displayorder uint   `json:"displayorder"` //
	Status       uint8  `json:"status"`       //
	Containtype  string `json:"containtype"`  //
}

// ImsRuleKeyword ims_rule_keyword
type ImsRuleKeyword struct {
	Id           uint   `json:"id"`           //
	Rid          uint   `json:"rid"`          //
	Uniacid      uint   `json:"uniacid"`      //
	Module       string `json:"module"`       //
	Content      string `json:"content"`      //
	Type         uint8  `json:"type"`         //
	Displayorder uint8  `json:"displayorder"` //
	Status       uint8  `json:"status"`       //
}

// ImsSiteArticle ims_site_article
type ImsSiteArticle struct {
	Id           uint   `json:"id"`           //
	Uniacid      uint   `json:"uniacid"`      //
	Rid          uint   `json:"rid"`          //
	Kid          uint   `json:"kid"`          //
	Iscommend    int8   `json:"iscommend"`    //
	Ishot        uint8  `json:"ishot"`        //
	Pcate        uint   `json:"pcate"`        //
	Ccate        uint   `json:"ccate"`        //
	Template     string `json:"template"`     //
	Title        string `json:"title"`        //
	Description  string `json:"description"`  //
	Content      string `json:"content"`      //
	Thumb        string `json:"thumb"`        //
	Incontent    int8   `json:"incontent"`    //
	Source       string `json:"source"`       //
	Author       string `json:"author"`       //
	Displayorder uint   `json:"displayorder"` //
	Linkurl      string `json:"linkurl"`      //
	Createtime   uint   `json:"createtime"`   //
	Edittime     int    `json:"edittime"`     //
	Click        uint   `json:"click"`        //
	Type         string `json:"type"`         //
	Credit       string `json:"credit"`       //
}

// ImsSiteArticleComment ims_site_article_comment
type ImsSiteArticleComment struct {
	Id         int     `json:"id"`         //
	Uniacid    int     `json:"uniacid"`    //
	Articleid  int     `json:"articleid"`  //
	Parentid   int     `json:"parentid"`   //
	Uid        int     `json:"uid"`        //
	Openid     string  `json:"openid"`     //
	Content    *string `json:"content"`    //
	IsRead     int8    `json:"is_read"`    //
	Iscomment  int8    `json:"iscomment"`  //
	Createtime int     `json:"createtime"` //
}

// ImsSiteCategory ims_site_category
type ImsSiteCategory struct {
	Id           uint   `json:"id"`           //
	Uniacid      uint   `json:"uniacid"`      //
	Nid          uint   `json:"nid"`          //
	Name         string `json:"name"`         //
	Parentid     uint   `json:"parentid"`     //
	Displayorder uint8  `json:"displayorder"` //
	Enabled      uint8  `json:"enabled"`      //
	Icon         string `json:"icon"`         //
	Description  string `json:"description"`  //
	Styleid      uint   `json:"styleid"`      //
	Linkurl      string `json:"linkurl"`      //
	Ishomepage   int8   `json:"ishomepage"`   //
	Icontype     uint8  `json:"icontype"`     //
	Css          string `json:"css"`          //
	Multiid      int    `json:"multiid"`      //
}

// ImsSiteMulti ims_site_multi
type ImsSiteMulti struct {
	Id       uint   `json:"id"`        //
	Uniacid  uint   `json:"uniacid"`   //
	Title    string `json:"title"`     //
	Styleid  uint   `json:"styleid"`   //
	SiteInfo string `json:"site_info"` //
	Status   uint8  `json:"status"`    //
	Bindhost string `json:"bindhost"`  //
}

// ImsSiteNav ims_site_nav
type ImsSiteNav struct {
	Id           uint   `json:"id"`           //
	Uniacid      uint   `json:"uniacid"`      //
	Multiid      uint   `json:"multiid"`      //
	Section      int8   `json:"section"`      //
	Module       string `json:"module"`       //
	Displayorder uint16 `json:"displayorder"` //
	Name         string `json:"name"`         //
	Description  string `json:"description"`  //
	Position     int8   `json:"position"`     //
	Url          string `json:"url"`          //
	Icon         string `json:"icon"`         //
	Css          string `json:"css"`          //
	Status       uint8  `json:"status"`       //
	Categoryid   uint   `json:"categoryid"`   //
}

// ImsSitePage ims_site_page
type ImsSitePage struct {
	Id          uint   `json:"id"`          //
	Uniacid     uint   `json:"uniacid"`     //
	Multiid     uint   `json:"multiid"`     //
	Title       string `json:"title"`       //
	Description string `json:"description"` //
	Params      string `json:"params"`      //
	Html        string `json:"html"`        //
	Multipage   string `json:"multipage"`   //
	Type        uint8  `json:"type"`        //
	Status      uint8  `json:"status"`      //
	Createtime  uint   `json:"createtime"`  //
	Goodnum     uint   `json:"goodnum"`     //
}

// ImsSiteSlide ims_site_slide
type ImsSiteSlide struct {
	Id           uint   `json:"id"`           //
	Uniacid      uint   `json:"uniacid"`      //
	Multiid      uint   `json:"multiid"`      //
	Title        string `json:"title"`        //
	Url          string `json:"url"`          //
	Thumb        string `json:"thumb"`        //
	Displayorder int8   `json:"displayorder"` //
}

// ImsSiteStoreCashLog ims_site_store_cash_log
type ImsSiteStoreCashLog struct {
	Id         int     `json:"id"`          //
	FounderUid int     `json:"founder_uid"` //
	Number     string  `json:"number"`      //
	Amount     float64 `json:"amount"`      //
	Status     int8    `json:"status"`      //
	CreateTime int     `json:"create_time"` //
}

// ImsSiteStoreCashOrder ims_site_store_cash_order
type ImsSiteStoreCashOrder struct {
	Id          int     `json:"id"`           //
	Number      string  `json:"number"`       //
	FounderUid  int     `json:"founder_uid"`  //
	OrderId     int     `json:"order_id"`     //
	GoodsId     int     `json:"goods_id"`     //
	OrderAmount float64 `json:"order_amount"` //
	CashLogId   int     `json:"cash_log_id"`  //
	Status      int8    `json:"status"`       //
	CreateTime  int     `json:"create_time"`  //
}

// ImsSiteStoreCreateAccount ims_site_store_create_account
type ImsSiteStoreCreateAccount struct {
	Id      int  `json:"id"`      //
	Uid     int  `json:"uid"`     //
	Uniacid int  `json:"uniacid"` //
	Type    int8 `json:"type"`    //
	Endtime int  `json:"endtime"` //
}

// ImsSiteStoreGoods ims_site_store_goods
type ImsSiteStoreGoods struct {
	Id             uint    `json:"id"`               //
	Type           int8    `json:"type"`             //
	Title          string  `json:"title"`            //
	Module         string  `json:"module"`           //
	AccountNum     int     `json:"account_num"`      //
	WxappNum       int     `json:"wxapp_num"`        //
	Price          float64 `json:"price"`            //
	Unit           string  `json:"unit"`             //
	Slide          string  `json:"slide"`            //
	CategoryId     int     `json:"category_id"`      //
	TitleInitial   string  `json:"title_initial"`    //
	Status         int8    `json:"status"`           //
	Createtime     int     `json:"createtime"`       //
	Synopsis       string  `json:"synopsis"`         //
	Description    string  `json:"description"`      //
	ModuleGroup    int     `json:"module_group"`     //
	ApiNum         int     `json:"api_num"`          //
	UserGroupPrice string  `json:"user_group_price"` //
	UserGroup      int     `json:"user_group"`       //
	AccountGroup   int     `json:"account_group"`    //
	IsWish         int8    `json:"is_wish"`          //
	Logo           string  `json:"logo"`             //
	PlatformNum    int     `json:"platform_num"`     //
	AliappNum      int     `json:"aliapp_num"`       //
	BaiduappNum    int     `json:"baiduapp_num"`     //
	PhoneappNum    int     `json:"phoneapp_num"`     //
	ToutiaoappNum  int     `json:"toutiaoapp_num"`   //
	WebappNum      int     `json:"webapp_num"`       //
	XzappNum       int     `json:"xzapp_num"`        //
}

// ImsSiteStoreGoodsCloud ims_site_store_goods_cloud
type ImsSiteStoreGoodsCloud struct {
	Id         int    `json:"id"`          //
	CloudId    int    `json:"cloud_id"`    //
	Name       string `json:"name"`        //
	Title      string `json:"title"`       //
	Logo       string `json:"logo"`        //
	WishBranch int    `json:"wish_branch"` //
	IsEdited   int8   `json:"is_edited"`   //
	Isdeleted  int8   `json:"isdeleted"`   //
	Branchs    string `json:"branchs"`     //
}

// ImsSiteStoreOrder ims_site_store_order
type ImsSiteStoreOrder struct {
	Id          uint    `json:"id"`          //
	Orderid     string  `json:"orderid"`     //
	Goodsid     int     `json:"goodsid"`     //
	Duration    int     `json:"duration"`    //
	Buyer       string  `json:"buyer"`       //
	Buyerid     int     `json:"buyerid"`     //
	Amount      float64 `json:"amount"`      //
	Type        int8    `json:"type"`        //
	Changeprice int8    `json:"changeprice"` //
	Createtime  int     `json:"createtime"`  //
	Uniacid     int     `json:"uniacid"`     //
	Endtime     int     `json:"endtime"`     //
	Wxapp       int     `json:"wxapp"`       //
	IsWish      int8    `json:"is_wish"`     //
}

// ImsSiteStyles ims_site_styles
type ImsSiteStyles struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Templateid uint   `json:"templateid"` //
	Name       string `json:"name"`       //
}

// ImsSiteStylesVars ims_site_styles_vars
type ImsSiteStylesVars struct {
	Id          uint   `json:"id"`          //
	Uniacid     uint   `json:"uniacid"`     //
	Templateid  uint   `json:"templateid"`  //
	Styleid     uint   `json:"styleid"`     //
	Variable    string `json:"variable"`    //
	Content     string `json:"content"`     //
	Description string `json:"description"` //
}

// ImsSiteTemplates ims_site_templates
type ImsSiteTemplates struct {
	Id          uint   `json:"id"`          //
	Name        string `json:"name"`        //
	Title       string `json:"title"`       //
	Version     string `json:"version"`     //
	Description string `json:"description"` //
	Author      string `json:"author"`      //
	Url         string `json:"url"`         //
	Type        string `json:"type"`        //
	Sections    uint   `json:"sections"`    //
}

// ImsStatFans ims_stat_fans
type ImsStatFans struct {
	Id       uint   `json:"id"`       //
	Uniacid  uint   `json:"uniacid"`  //
	New      uint   `json:"new"`      //
	Cancel   uint   `json:"cancel"`   //
	Cumulate int    `json:"cumulate"` //
	Date     string `json:"date"`     //
}

// ImsStatKeyword ims_stat_keyword
type ImsStatKeyword struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Rid        string `json:"rid"`        //
	Kid        uint   `json:"kid"`        //
	Hit        uint   `json:"hit"`        //
	Lastupdate uint   `json:"lastupdate"` //
	Createtime uint   `json:"createtime"` //
}

// ImsStatMsgHistory ims_stat_msg_history
type ImsStatMsgHistory struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Rid        uint   `json:"rid"`        //
	Kid        uint   `json:"kid"`        //
	FromUser   string `json:"from_user"`  //
	Module     string `json:"module"`     //
	Message    string `json:"message"`    //
	Type       string `json:"type"`       //
	Createtime uint   `json:"createtime"` //
}

// ImsStatRule ims_stat_rule
type ImsStatRule struct {
	Id         uint `json:"id"`         //
	Uniacid    uint `json:"uniacid"`    //
	Rid        uint `json:"rid"`        //
	Hit        uint `json:"hit"`        //
	Lastupdate uint `json:"lastupdate"` //
	Createtime uint `json:"createtime"` //
}

// ImsStatVisit ims_stat_visit
type ImsStatVisit struct {
	Id      uint   `json:"id"`       //
	Uniacid int    `json:"uniacid"`  //
	Module  string `json:"module"`   //
	Count   uint   `json:"count"`    //
	Date    uint   `json:"date"`     //
	Type    string `json:"type"`     //
	IpCount int    `json:"ip_count"` //
}

// ImsStatVisitIp ims_stat_visit_ip
type ImsStatVisitIp struct {
	Id      int    `json:"id"`      //
	Ip      int64  `json:"ip"`      //
	Uniacid int    `json:"uniacid"` //
	Type    string `json:"type"`    //
	Module  string `json:"module"`  //
	Date    int    `json:"date"`    //
}

// ImsSystemStatVisit ims_system_stat_visit
type ImsSystemStatVisit struct {
	Id           int    `json:"id"`           //
	Uniacid      int    `json:"uniacid"`      //
	Modulename   string `json:"modulename"`   //
	Uid          int    `json:"uid"`          //
	Displayorder int    `json:"displayorder"` //
	Createtime   int    `json:"createtime"`   //
	Updatetime   int    `json:"updatetime"`   //
}

// ImsSystemWelcomeBinddomain ims_system_welcome_binddomain
type ImsSystemWelcomeBinddomain struct {
	Id         int    `json:"id"`          //
	Uid        int    `json:"uid"`         //
	ModuleName string `json:"module_name"` //
	Domain     string `json:"domain"`      //
	Createtime int    `json:"createtime"`  //
}

// ImsUniAccount ims_uni_account
type ImsUniAccount struct {
	Uniacid      uint   `json:"uniacid"`       //
	Groupid      int    `json:"groupid"`       //
	Name         string `json:"name"`          //
	Description  string `json:"description"`   //
	DefaultAcid  uint   `json:"default_acid"`  //
	Rank         *int   `json:"rank"`          //
	TitleInitial string `json:"title_initial"` //
	Createtime   int    `json:"createtime"`    //
	Logo         string `json:"logo"`          //
	Qrcode       string `json:"qrcode"`        //
	CreateUid    int    `json:"create_uid"`    //
}

// ImsUniAccountExtraModules ims_uni_account_extra_modules
type ImsUniAccountExtraModules struct {
	Id      uint   `json:"id"`      //
	Uniacid uint   `json:"uniacid"` //
	Modules string `json:"modules"` //
}

// ImsUniAccountGroup ims_uni_account_group
type ImsUniAccountGroup struct {
	Id      uint `json:"id"`      //
	Uniacid uint `json:"uniacid"` //
	Groupid int  `json:"groupid"` //
}

// ImsUniAccountMenus ims_uni_account_menus
type ImsUniAccountMenus struct {
	Id                 uint   `json:"id"`                   //
	Uniacid            uint   `json:"uniacid"`              //
	Menuid             uint   `json:"menuid"`               //
	Type               uint8  `json:"type"`                 //
	Title              string `json:"title"`                //
	Sex                uint8  `json:"sex"`                  //
	GroupId            int    `json:"group_id"`             //
	ClientPlatformType uint8  `json:"client_platform_type"` //
	Area               string `json:"area"`                 //
	Data               string `json:"data"`                 //
	Status             uint8  `json:"status"`               //
	Createtime         uint   `json:"createtime"`           //
	Isdeleted          uint8  `json:"isdeleted"`            //
}

// ImsUniAccountModules ims_uni_account_modules
type ImsUniAccountModules struct {
	Id           uint   `json:"id"`           //
	Uniacid      uint   `json:"uniacid"`      //
	Module       string `json:"module"`       //
	Enabled      uint8  `json:"enabled"`      //
	Settings     string `json:"settings"`     //
	Shortcut     uint8  `json:"shortcut"`     //
	Displayorder uint   `json:"displayorder"` //
}

// ImsUniAccountModulesShortcut ims_uni_account_modules_shortcut
type ImsUniAccountModulesShortcut struct {
	Id         int    `json:"id"`          //
	Title      string `json:"title"`       //
	Url        string `json:"url"`         //
	Icon       string `json:"icon"`        //
	Uniacid    int    `json:"uniacid"`     //
	VersionId  int    `json:"version_id"`  //
	ModuleName string `json:"module_name"` //
}

// ImsUniAccountUsers ims_uni_account_users
type ImsUniAccountUsers struct {
	Id      uint   `json:"id"`      //
	Uniacid uint   `json:"uniacid"` //
	Uid     uint   `json:"uid"`     //
	Role    string `json:"role"`    //
	Rank    uint8  `json:"rank"`    //
}

// ImsUniGroup ims_uni_group
type ImsUniGroup struct {
	Id        uint   `json:"id"`        //
	OwnerUid  int    `json:"owner_uid"` //
	Name      string `json:"name"`      //
	Modules   string `json:"modules"`   //
	Templates string `json:"templates"` //
	Uniacid   uint   `json:"uniacid"`   //
	Uid       int    `json:"uid"`       //
}

// ImsUniLinkUniacid ims_uni_link_uniacid
type ImsUniLinkUniacid struct {
	Id          int    `json:"id"`           //
	Uniacid     int    `json:"uniacid"`      //
	LinkUniacid int    `json:"link_uniacid"` //
	VersionId   int    `json:"version_id"`   //
	ModuleName  string `json:"module_name"`  //
}

// ImsUniModules ims_uni_modules
type ImsUniModules struct {
	Id         int    `json:"id"`          //
	Uniacid    int    `json:"uniacid"`     //
	ModuleName string `json:"module_name"` //
}

// ImsUniSettings ims_uni_settings
type ImsUniSettings struct {
	Uniacid         uint   `json:"uniacid"`          //
	Passport        string `json:"passport"`         //
	Oauth           string `json:"oauth"`            //
	JsauthAcid      uint   `json:"jsauth_acid"`      //
	Notify          string `json:"notify"`           //
	Creditnames     string `json:"creditnames"`      //
	Creditbehaviors string `json:"creditbehaviors"`  //
	Welcome         string `json:"welcome"`          //
	Default         string `json:"default"`          //
	DefaultMessage  string `json:"default_message"`  //
	Payment         string `json:"payment"`          //
	Stat            string `json:"stat"`             //
	DefaultSite     *uint  `json:"default_site"`     //
	Sync            uint8  `json:"sync"`             //
	Recharge        string `json:"recharge"`         //
	Tplnotice       string `json:"tplnotice"`        //
	Grouplevel      uint8  `json:"grouplevel"`       //
	Mcplugin        string `json:"mcplugin"`         //
	ExchangeEnable  uint8  `json:"exchange_enable"`  //
	CouponType      uint8  `json:"coupon_type"`      //
	Menuset         string `json:"menuset"`          //
	Statistics      string `json:"statistics"`       //
	BindDomain      string `json:"bind_domain"`      //
	CommentStatus   int8   `json:"comment_status"`   //
	ReplySetting    int8   `json:"reply_setting"`    //
	DefaultModule   string `json:"default_module"`   //
	AttachmentLimit int    `json:"attachment_limit"` //
	AttachmentSize  string `json:"attachment_size"`  //
	SyncMember      int8   `json:"sync_member"`      //
	Remote          string `json:"remote"`           //
}

// ImsUniVerifycode ims_uni_verifycode
type ImsUniVerifycode struct {
	Id          uint   `json:"id"`           //
	Uniacid     uint   `json:"uniacid"`      //
	Receiver    string `json:"receiver"`     //
	Verifycode  string `json:"verifycode"`   //
	Total       uint8  `json:"total"`        //
	Createtime  uint   `json:"createtime"`   //
	FailedCount *int   `json:"failed_count"` //
}

// ImsUserapiCache ims_userapi_cache
type ImsUserapiCache struct {
	Id         uint   `json:"id"`         //
	Key        string `json:"key"`        //
	Content    string `json:"content"`    //
	Lastupdate uint   `json:"lastupdate"` //
}

// ImsUserapiReply ims_userapi_reply
type ImsUserapiReply struct {
	Id          uint   `json:"id"`           //
	Rid         uint   `json:"rid"`          //
	Description string `json:"description"`  //
	Apiurl      string `json:"apiurl"`       //
	Token       string `json:"token"`        //
	DefaultText string `json:"default_text"` //
	Cachetime   uint   `json:"cachetime"`    //
}

// ImsUsers ims_users
type ImsUsers struct {
	Uid            uint   `json:"uid"`             //
	OwnerUid       int    `json:"owner_uid"`       //
	Groupid        uint   `json:"groupid"`         //
	FounderGroupid int8   `json:"founder_groupid"` //
	Username       string `json:"username"`        //
	Password       string `json:"password"`        //
	Salt           string `json:"salt"`            //
	Type           uint8  `json:"type"`            //
	Status         int8   `json:"status"`          //
	Joindate       uint   `json:"joindate"`        //
	Joinip         string `json:"joinip"`          //
	Lastvisit      uint   `json:"lastvisit"`       //
	Lastip         string `json:"lastip"`          //
	Remark         string `json:"remark"`          //
	Starttime      uint   `json:"starttime"`       //
	Endtime        uint   `json:"endtime"`         //
	RegisterType   int8   `json:"register_type"`   //
	Openid         string `json:"openid"`          //
	WelcomeLink    int8   `json:"welcome_link"`    //
	NoticeSetting  string `json:"notice_setting"`  //
	IsBind         int8   `json:"is_bind"`         //
}

// ImsUsersBind ims_users_bind
type ImsUsersBind struct {
	Id            int    `json:"id"`             //
	Uid           int    `json:"uid"`            //
	BindSign      string `json:"bind_sign"`      //
	ThirdType     int8   `json:"third_type"`     //
	ThirdNickname string `json:"third_nickname"` //
}

// ImsUsersCreateGroup ims_users_create_group
type ImsUsersCreateGroup struct {
	Id            int    `json:"id"`            //
	GroupName     string `json:"group_name"`    //
	Maxaccount    int    `json:"maxaccount"`    //
	Maxwxapp      int    `json:"maxwxapp"`      //
	Maxwebapp     int    `json:"maxwebapp"`     //
	Maxphoneapp   int    `json:"maxphoneapp"`   //
	Maxxzapp      int    `json:"maxxzapp"`      //
	Maxaliapp     int    `json:"maxaliapp"`     //
	Createtime    int    `json:"createtime"`    //
	Maxbaiduapp   int    `json:"maxbaiduapp"`   //
	Maxtoutiaoapp int    `json:"maxtoutiaoapp"` //
}

// ImsUsersExtraGroup ims_users_extra_group
type ImsUsersExtraGroup struct {
	Id            int `json:"id"`              //
	Uid           int `json:"uid"`             //
	UniGroupId    int `json:"uni_group_id"`    //
	CreateGroupId int `json:"create_group_id"` //
}

// ImsUsersExtraLimit ims_users_extra_limit
type ImsUsersExtraLimit struct {
	Id            int `json:"id"`            //
	Uid           int `json:"uid"`           //
	Maxaccount    int `json:"maxaccount"`    //
	Maxwxapp      int `json:"maxwxapp"`      //
	Maxwebapp     int `json:"maxwebapp"`     //
	Maxphoneapp   int `json:"maxphoneapp"`   //
	Maxxzapp      int `json:"maxxzapp"`      //
	Maxaliapp     int `json:"maxaliapp"`     //
	Timelimit     int `json:"timelimit"`     //
	Maxbaiduapp   int `json:"maxbaiduapp"`   //
	Maxtoutiaoapp int `json:"maxtoutiaoapp"` //
}

// ImsUsersExtraModules ims_users_extra_modules
type ImsUsersExtraModules struct {
	Id         int    `json:"id"`          //
	Uid        int    `json:"uid"`         //
	ModuleName string `json:"module_name"` //
	Support    string `json:"support"`     //
}

// ImsUsersExtraTemplates ims_users_extra_templates
type ImsUsersExtraTemplates struct {
	Id         int `json:"id"`          //
	Uid        int `json:"uid"`         //
	TemplateId int `json:"template_id"` //
}

// ImsUsersFailedLogin ims_users_failed_login
type ImsUsersFailedLogin struct {
	Id         uint   `json:"id"`         //
	Ip         string `json:"ip"`         //
	Username   string `json:"username"`   //
	Count      uint8  `json:"count"`      //
	Lastupdate uint   `json:"lastupdate"` //
}

// ImsUsersFounderGroup ims_users_founder_group
type ImsUsersFounderGroup struct {
	Id            uint   `json:"id"`            //
	Name          string `json:"name"`          //
	Package       string `json:"package"`       //
	Maxaccount    uint   `json:"maxaccount"`    //
	Timelimit     uint   `json:"timelimit"`     //
	Maxwxapp      uint   `json:"maxwxapp"`      //
	Maxwebapp     int    `json:"maxwebapp"`     //
	Maxphoneapp   int    `json:"maxphoneapp"`   //
	Maxxzapp      int    `json:"maxxzapp"`      //
	Maxaliapp     int    `json:"maxaliapp"`     //
	Maxbaiduapp   int    `json:"maxbaiduapp"`   //
	Maxtoutiaoapp int    `json:"maxtoutiaoapp"` //
}

// ImsUsersFounderOwnCreateGroups ims_users_founder_own_create_groups
type ImsUsersFounderOwnCreateGroups struct {
	Id            int `json:"id"`              //
	FounderUid    int `json:"founder_uid"`     //
	CreateGroupId int `json:"create_group_id"` //
}

// ImsUsersFounderOwnUniGroups ims_users_founder_own_uni_groups
type ImsUsersFounderOwnUniGroups struct {
	Id         int `json:"id"`           //
	FounderUid int `json:"founder_uid"`  //
	UniGroupId int `json:"uni_group_id"` //
}

// ImsUsersFounderOwnUsers ims_users_founder_own_users
type ImsUsersFounderOwnUsers struct {
	Id         int `json:"id"`          //
	Uid        int `json:"uid"`         //
	FounderUid int `json:"founder_uid"` //
}

// ImsUsersFounderOwnUsersGroups ims_users_founder_own_users_groups
type ImsUsersFounderOwnUsersGroups struct {
	Id           int `json:"id"`             //
	FounderUid   int `json:"founder_uid"`    //
	UsersGroupId int `json:"users_group_id"` //
}

// ImsUsersGroup ims_users_group
type ImsUsersGroup struct {
	Id            uint   `json:"id"`            //
	OwnerUid      int    `json:"owner_uid"`     //
	Name          string `json:"name"`          //
	Package       string `json:"package"`       //
	Maxaccount    uint   `json:"maxaccount"`    //
	Timelimit     uint   `json:"timelimit"`     //
	Maxwxapp      uint   `json:"maxwxapp"`      //
	Maxwebapp     int    `json:"maxwebapp"`     //
	Maxphoneapp   int    `json:"maxphoneapp"`   //
	Maxxzapp      int    `json:"maxxzapp"`      //
	Maxaliapp     int    `json:"maxaliapp"`     //
	Maxbaiduapp   int    `json:"maxbaiduapp"`   //
	Maxtoutiaoapp int    `json:"maxtoutiaoapp"` //
}

// ImsUsersInvitation ims_users_invitation
type ImsUsersInvitation struct {
	Id         uint   `json:"id"`         //
	Code       string `json:"code"`       //
	Fromuid    uint   `json:"fromuid"`    //
	Inviteuid  uint   `json:"inviteuid"`  //
	Createtime uint   `json:"createtime"` //
}

// ImsUsersLastuse ims_users_lastuse
type ImsUsersLastuse struct {
	Id         int    `json:"id"`         //
	Uid        int    `json:"uid"`        //
	Uniacid    int    `json:"uniacid"`    //
	Modulename string `json:"modulename"` //
	Type       string `json:"type"`       //
}

// ImsUsersLoginLogs ims_users_login_logs
type ImsUsersLoginLogs struct {
	Id         uint   `json:"id"`         //
	Uid        uint   `json:"uid"`        //
	Ip         string `json:"ip"`         //
	City       string `json:"city"`       //
	LoginAt    uint   `json:"login_at"`   //
	Createtime uint   `json:"createtime"` //
}

// ImsUsersOperateHistory ims_users_operate_history
type ImsUsersOperateHistory struct {
	Id         uint   `json:"id"`          //
	Type       uint8  `json:"type"`        //
	Uid        uint   `json:"uid"`         //
	Uniacid    uint   `json:"uniacid"`     //
	ModuleName string `json:"module_name"` //
	Createtime int    `json:"createtime"`  //
}

// ImsUsersOperateStar ims_users_operate_star
type ImsUsersOperateStar struct {
	Id         uint   `json:"id"`          //
	Type       uint8  `json:"type"`        //
	Uid        uint   `json:"uid"`         //
	Uniacid    uint   `json:"uniacid"`     //
	ModuleName string `json:"module_name"` //
	Rank       int    `json:"rank"`        //
	Createtime int    `json:"createtime"`  //
}

// ImsUsersPermission ims_users_permission
type ImsUsersPermission struct {
	Id         uint   `json:"id"`         //
	Uniacid    uint   `json:"uniacid"`    //
	Uid        uint   `json:"uid"`        //
	Type       string `json:"type"`       //
	Permission string `json:"permission"` //
	Url        string `json:"url"`        //
	Modules    string `json:"modules"`    //
	Templates  string `json:"templates"`  //
}

// ImsUsersProfile ims_users_profile
type ImsUsersProfile struct {
	Id                 uint   `json:"id"`                    //
	Uid                uint   `json:"uid"`                   //
	Createtime         uint   `json:"createtime"`            //
	Edittime           int    `json:"edittime"`              //
	Realname           string `json:"realname"`              //
	Nickname           string `json:"nickname"`              //
	Avatar             string `json:"avatar"`                //
	Qq                 string `json:"qq"`                    //
	Mobile             string `json:"mobile"`                //
	Fakeid             string `json:"fakeid"`                //
	Vip                uint8  `json:"vip"`                   //
	Gender             int8   `json:"gender"`                //
	Birthyear          uint16 `json:"birthyear"`             //
	Birthmonth         uint8  `json:"birthmonth"`            //
	Birthday           uint8  `json:"birthday"`              //
	Constellation      string `json:"constellation"`         //
	Zodiac             string `json:"zodiac"`                //
	Telephone          string `json:"telephone"`             //
	Idcard             string `json:"idcard"`                //
	Studentid          string `json:"studentid"`             //
	Grade              string `json:"grade"`                 //
	Address            string `json:"address"`               //
	Zipcode            string `json:"zipcode"`               //
	Nationality        string `json:"nationality"`           //
	Resideprovince     string `json:"resideprovince"`        //
	Residecity         string `json:"residecity"`            //
	Residedist         string `json:"residedist"`            //
	Graduateschool     string `json:"graduateschool"`        //
	Company            string `json:"company"`               //
	Education          string `json:"education"`             //
	Occupation         string `json:"occupation"`            //
	Position           string `json:"position"`              //
	Revenue            string `json:"revenue"`               //
	Affectivestatus    string `json:"affectivestatus"`       //
	Lookingfor         string `json:"lookingfor"`            //
	Bloodtype          string `json:"bloodtype"`             //
	Height             string `json:"height"`                //
	Weight             string `json:"weight"`                //
	Alipay             string `json:"alipay"`                //
	Msn                string `json:"msn"`                   //
	Email              string `json:"email"`                 //
	Taobao             string `json:"taobao"`                //
	Site               string `json:"site"`                  //
	Bio                string `json:"bio"`                   //
	Interest           string `json:"interest"`              //
	Workerid           string `json:"workerid"`              //
	IsSendMobileStatus int8   `json:"is_send_mobile_status"` //
	SendExpireStatus   int8   `json:"send_expire_status"`    //
}

// ImsVideoReply ims_video_reply
type ImsVideoReply struct {
	Id          uint   `json:"id"`          //
	Rid         uint   `json:"rid"`         //
	Title       string `json:"title"`       //
	Description string `json:"description"` //
	Mediaid     string `json:"mediaid"`     //
	Createtime  int    `json:"createtime"`  //
}

// ImsVoiceReply ims_voice_reply
type ImsVoiceReply struct {
	Id         uint   `json:"id"`         //
	Rid        uint   `json:"rid"`        //
	Title      string `json:"title"`      //
	Mediaid    string `json:"mediaid"`    //
	Createtime int    `json:"createtime"` //
}

// ImsVxFan ims_vx_fan 微信公众号粉丝表
type ImsVxFan struct {
	Id              int64  `json:"id"`                // id
	Pid             int64  `json:"pid"`               // 推荐人id
	AccountId       int64  `json:"account_id"`        // 微信公众账号id
	AccountCategory string `json:"account_category"`  // 微信账号类型
	Openid          string `json:"openid"`            // 粉丝openid
	GroupId         int64  `json:"group_id"`          // 粉丝group_id
	UnionId         string `json:"union_id"`          // 粉丝union_id
	Mobile          string `json:"mobile"`            // 粉丝手机号
	Category        string `json:"category"`          // 粉丝类型
	CategoryMsg     string `json:"category_msg"`      // 类型描述
	Subscribe       string `json:"subscribe"`         // 粉丝是否关注公众号
	Nickname        string `json:"nickname"`          // 粉丝昵称
	Sex             string `json:"sex"`               // 粉丝性别
	Language        string `json:"language"`          // 粉丝语言
	City            string `json:"city"`              // 粉丝城市
	Province        string `json:"province"`          // 粉丝省州
	Country         string `json:"country"`           // 粉丝国家
	Avatar          string `json:"avatar"`            // 粉丝头像
	SubscribeAt     int64  `json:"subscribe_at"`      // 粉丝关注时间
	State           int    `json:"state"`             // 状态
	StateMsg        string `json:"state_msg"`         // 状态描述
	CreatedAt       int64  `json:"created_at"`        // 创建时间戳
	UpdatedAt       int64  `json:"updated_at"`        // 更新时间戳
	Note            string `json:"note"`              // 备注
	RecommendQrCode string `json:"recommend_qr_code"` // 推广二维码
}

// ImsVxFanAccount ims_vx_fan_account 微信粉丝账户表
type ImsVxFanAccount struct {
	Id              int64  `json:"id"`                // id
	Uid             int64  `json:"uid"`               // 粉丝id
	Balance         int64  `json:"balance"`           // 零钱余额
	Prepare         int64  `json:"prepare"`           // 可提现额度
	Withdraw        int64  `json:"withdraw"`          // 累计提现金额
	LastWithdrawAt  int64  `json:"last_withdraw_at"`  // 最后一次提现时间戳
	LastWithdrawDay int64  `json:"last_withdraw_day"` // 最后一次提现日,如:20010101
	State           int    `json:"state"`             // 状态
	StateMsg        string `json:"state_msg"`         // 状态描述
	CreatedAt       int64  `json:"created_at"`        // 创建时间戳
	UpdatedAt       int64  `json:"updated_at"`        // 更新时间戳
	Note            string `json:"note"`              // 备注
}

// ImsVxFanAddress ims_vx_fan_address 微信公众号粉丝地址表
type ImsVxFanAddress struct {
	Id         int64  `json:"id"`          // id
	Uid        int64  `json:"uid"`         // 粉丝id
	Consignee  string `json:"consignee"`   // 收货人姓名
	Mobile     string `json:"mobile"`      // 收货人手机号
	Address1   string `json:"address1"`    // 收货人地址1
	Address2   string `json:"address2"`    // 收货人地址2
	Address3   string `json:"address3"`    // 收货人地址3
	Address4   string `json:"address4"`    // 收货人地址4
	PostalCode string `json:"postal_code"` // 邮政编码
	IsDefault  int8   `json:"is_default"`  // 是否默认
	State      int    `json:"state"`       // 状态
	StateMsg   string `json:"state_msg"`   // 状态描述
	CreatedAt  int64  `json:"created_at"`  // 创建时间戳
	UpdatedAt  int64  `json:"updated_at"`  // 更新时间戳
	DeletedAt  int64  `json:"deleted_at"`  // 删除时间戳
	Note       string `json:"note"`        // 备注
}

// ImsVxFanBalanceLog ims_vx_fan_balance_log 微信粉丝账户零钱明细表
type ImsVxFanBalanceLog struct {
	Id          int64  `json:"id"`           // id
	Uid         int64  `json:"uid"`          // 粉丝id
	Val1        int64  `json:"val1"`         // 变化前的值
	Val2        int64  `json:"val2"`         // 本次变化值
	Val3        int64  `json:"val3"`         // 变化前的后
	Category    int    `json:"category"`     // 明细类型
	CategoryMsg string `json:"category_msg"` // 明细类型描述
	OrderId     int64  `json:"order_id"`     // 对应订单id
	OrderNo     string `json:"order_no"`     // 对应订单号
	State       int    `json:"state"`        // 状态
	StateMsg    string `json:"state_msg"`    // 状态描述
	CreatedAt   int64  `json:"created_at"`   // 创建时间戳
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间戳
	Note        string `json:"note"`         // 备注
}

// ImsVxFanGoods ims_vx_fan_goods 微信商品表
type ImsVxFanGoods struct {
	Id              int64  `json:"id"`               // id
	AccountId       int64  `json:"account_id"`       // 微信公众账号id
	Avatar          string `json:"avatar"`           // 商品标识图
	Name            string `json:"name"`             // 商品名称
	Title           string `json:"title"`            // 标题
	TitleDesc       string `json:"title_desc"`       // 副标题
	Label           string `json:"label"`            // 标签
	Details         string `json:"details"`          // 商品详情
	Video           string `json:"video"`            // 商品视频地址
	PriceNominal    int64  `json:"price_nominal"`    // 商品名义价,单位:分
	Price           int64  `json:"price"`            // 商品单价,单位:分
	Postage         int64  `json:"postage"`          // 邮费,单位:分
	Stock           int64  `json:"stock"`            // 商品库存数量
	Sold            int64  `json:"sold"`             // 商品已卖数量
	Serial          int    `json:"serial"`           // 序号
	Category        string `json:"category"`         // 商品类型
	CategoryMsg     string `json:"category_msg"`     // 商品类型描述
	Uid             int64  `json:"uid"`              // 所属用户id
	DeliveryAddress string `json:"delivery_address"` // 发货地址
	Recommend       int8   `json:"recommend"`        // 是否推荐
	State           int    `json:"state"`            // 状态 0:待上架,1:已上架,2:已下架
	StateMsg        string `json:"state_msg"`        // 状态描述
	CreatedAt       int64  `json:"created_at"`       // 创建时间戳
	UpdatedAt       int64  `json:"updated_at"`       // 更新时间戳
	DeletedAt       int64  `json:"deleted_at"`       // 删除时间戳
	Note            string `json:"note"`             // 备注
}

// ImsVxFanGoodsRule ims_vx_fan_goods_rule 微信商品规格表
type ImsVxFanGoodsRule struct {
	Id           int64  `json:"id"`            // id
	GoodsId      int64  `json:"goods_id"`      // 商品id
	Attributes   string `json:"attributes"`    // 规格
	Serial       int    `json:"serial"`        // 序号,默认升序
	Avatar       string `json:"avatar"`        // 规格主图
	Details      string `json:"details"`       // 规格详情
	PriceNominal int64  `json:"price_nominal"` // 商品名义价,单位:分
	Price        int64  `json:"price"`         // 商品单价,单位:分
	Stock        int64  `json:"stock"`         // 商品库存数量
	Sold         int64  `json:"sold"`          // 商品已卖数量
	State        int    `json:"state"`         // 状态 0:待上架,1:已上架,2:已下架
	StateMsg     string `json:"state_msg"`     // 状态描述
	CreatedAt    int64  `json:"created_at"`    // 创建时间戳
	UpdatedAt    int64  `json:"updated_at"`    // 更新时间戳
	DeletedAt    int64  `json:"deleted_at"`    // 删除时间戳
	Note         string `json:"note"`          // 备注
}

// ImsVxFanOrderBagQuotaCard ims_vx_fan_order_bag_quota_card 微信粉丝订单红包额度卡表
type ImsVxFanOrderBagQuotaCard struct {
	Id           int64   `json:"id"`            // id
	Uid          int64   `json:"uid"`           // 会员id
	AccountId    int64   `json:"account_id"`    // 公众号id
	GoodsId      int64   `json:"goods_id"`      // 商品id
	GoodsNum     int64   `json:"goods_num"`     // 商品数量
	GoodsPrice   int64   `json:"goods_price"`   // 商品单价
	GoodsTitle   *string `json:"goods_title"`   // 商品标题
	OrderNo      string  `json:"order_no"`      // 订单号
	Amount       int64   `json:"amount"`        // 订单金额,单位:分
	PayAmount    int64   `json:"pay_amount"`    // 支付金额,单位:分
	PayType      string  `json:"pay_type"`      // 支付方式
	PayState     string  `json:"pay_state"`     // 支付状态
	PayNo        string  `json:"pay_no"`        // 支付订单号
	PayDetails   string  `json:"pay_details"`   // 支付详细信息
	GoodsDetails string  `json:"goods_details"` // 商品详细信息
	UserDetails  string  `json:"user_details"`  // 会员详细信息
	State        int     `json:"state"`         // 状态
	State1       string  `json:"state1"`        // 状态描述
	CreatedAt    int64   `json:"created_at"`    // 创建时间戳
	UpdatedAt    int64   `json:"updated_at"`    // 更新时间戳
	Note         string  `json:"note"`          // 备注
}

// ImsVxFanOrderGoods ims_vx_fan_order_goods 微信粉丝订单商品表
type ImsVxFanOrderGoods struct {
	Id             int64  `json:"id"`              // id
	Uid            int64  `json:"uid"`             // 会员id
	AccountId      int64  `json:"account_id"`      // 公众号id
	GoodsId        int64  `json:"goods_id"`        // 商品id
	GoodsNum       int64  `json:"goods_num"`       // 商品数量
	GoodsPrice     int64  `json:"goods_price"`     // 商品单价
	OrderNo        string `json:"order_no"`        // 订单号
	Amount         int64  `json:"amount"`          // 订单金额,单位:分
	DiscountAmount int64  `json:"discount_amount"` // 优惠金额,单位:分
	Discount       string `json:"discount"`        // 优惠信息
	PayAmount      int64  `json:"pay_amount"`      // 支付金额,单位:分
	PayType        string `json:"pay_type"`        // 支付方式
	PayState       string `json:"pay_state"`       // 支付状态
	PayNo          string `json:"pay_no"`          // 支付订单号
	PayDetails     string `json:"pay_details"`     // 支付详细信息
	UserDetails    string `json:"user_details"`    // 会员信息
	GoodsDetails   string `json:"goods_details"`   // 商品详细信息
	AddressDetails string `json:"address_details"` // 收货地址信息
	AddressId      int64  `json:"address_id"`      // 收货地址id
	Address        string `json:"address"`         // 收货地址
	State          int    `json:"state"`           // 状态
	State1         string `json:"state1"`          // 状态描述
	CreatedAt      int64  `json:"created_at"`      // 创建时间戳
	UpdatedAt      int64  `json:"updated_at"`      // 更新时间戳
	Note           string `json:"note"`            // 备注
}

// ImsVxFanOrderRightCard ims_vx_fan_order_right_card 微信粉丝订单权益卡表
type ImsVxFanOrderRightCard struct {
	Id           int64   `json:"id"`            // id
	Uid          int64   `json:"uid"`           // 会员id
	AccountId    int64   `json:"account_id"`    // 公众号id
	GoodsId      int64   `json:"goods_id"`      // 商品id
	GoodsNum     int64   `json:"goods_num"`     // 商品数量
	GoodsPrice   int64   `json:"goods_price"`   // 商品单价
	GoodsTitle   *string `json:"goods_title"`   // 商品标题
	OrderNo      string  `json:"order_no"`      // 订单号
	Amount       int64   `json:"amount"`        // 订单金额,单位:分
	PayAmount    int64   `json:"pay_amount"`    // 支付金额,单位:分
	PayType      string  `json:"pay_type"`      // 支付方式
	PayState     string  `json:"pay_state"`     // 支付状态
	PayNo        string  `json:"pay_no"`        // 支付订单号
	PayDetails   string  `json:"pay_details"`   // 支付详细信息
	GoodsDetails string  `json:"goods_details"` // 商品详细信息
	UserDetails  string  `json:"user_details"`  // 会员详细信息
	State        int     `json:"state"`         // 状态
	State1       string  `json:"state1"`        // 状态描述
	CreatedAt    int64   `json:"created_at"`    // 创建时间戳
	UpdatedAt    int64   `json:"updated_at"`    // 更新时间戳
	Note         string  `json:"note"`          // 备注
}

// ImsVxFanPrepareLog ims_vx_fan_prepare_log 微信粉丝账户零钱额度明细表
type ImsVxFanPrepareLog struct {
	Id          int64  `json:"id"`           // id
	Uid         int64  `json:"uid"`          // 粉丝id
	Val1        int64  `json:"val1"`         // 变化前的值
	Val2        int64  `json:"val2"`         // 本次变化值
	Val3        int64  `json:"val3"`         // 变化前的后
	Category    int    `json:"category"`     // 明细类型
	CategoryMsg string `json:"category_msg"` // 明细类型描述
	OrderId     int64  `json:"order_id"`     // 对应订单id
	OrderNo     string `json:"order_no"`     // 对应订单号
	State       int    `json:"state"`        // 状态
	StateMsg    string `json:"state_msg"`    // 状态描述
	CreatedAt   int64  `json:"created_at"`   // 创建时间戳
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间戳
	Note        string `json:"note"`         // 备注
}

// ImsVxFanRightCard ims_vx_fan_right_card 微信粉丝权益卡表
type ImsVxFanRightCard struct {
	Id        int64   `json:"id"`         // id
	AccountId int64   `json:"account_id"` // 微信公众账号id
	Name      string  `json:"name"`       // 权益卡名称
	Money     int64   `json:"money"`      // 权益卡单价,单位:分
	Avatar    string  `json:"avatar"`     // 权益卡标识图
	Title     string  `json:"title"`      // 标题
	Details   string  `json:"details"`    // 权益卡详情
	Multiple  float64 `json:"multiple"`   // 返佣倍数
	Vip       string  `json:"vip"`        // 为粉丝提供的权益vip等级
	Discount  float64 `json:"discount"`   // 商城购物折扣
	Serial    int     `json:"serial"`     // 序号
	State     int     `json:"state"`      // 状态
	StateMsg  string  `json:"state_msg"`  // 状态描述
	CreatedAt int64   `json:"created_at"` // 创建时间戳
	UpdatedAt int64   `json:"updated_at"` // 更新时间戳
	Note      string  `json:"note"`       // 备注
}

// ImsVxFanRightCardUpgrade ims_vx_fan_right_card_upgrade 微信粉丝权益卡升级表
type ImsVxFanRightCardUpgrade struct {
	Id              int64   `json:"id"`                // id
	Uid             int64   `json:"uid"`               // 会员id
	UpgradeRatio    float64 `json:"upgrade_ratio"`     // 升级返佣比例
	FromUid         int64   `json:"from_uid"`          // 来自会员id
	RightCardAmount int64   `json:"right_card_amount"` // 权益卡金额
	State           int     `json:"state"`             // 状态
	StateMsg        string  `json:"state_msg"`         // 状态描述
	CreatedAt       int64   `json:"created_at"`        // 创建时间戳
	UpdatedAt       int64   `json:"updated_at"`        // 更新时间戳
	Note            string  `json:"note"`              // 备注
	Cards           string  `json:"cards"`             // 权益卡信息
}

// ImsVxFanRightCardUsed ims_vx_fan_right_card_used 微信粉丝权益卡已使用表
type ImsVxFanRightCardUsed struct {
	Id                 int64   `json:"id"`                   // id
	Cid                int64   `json:"cid"`                  // 权益卡id
	Uid                int64   `json:"uid"`                  // 会员id
	Serial             string  `json:"serial"`               // 订单号
	Name               string  `json:"name"`                 // 权益卡名称
	Money              int64   `json:"money"`                // 权益卡金额,单位:分
	Multiple           float64 `json:"multiple"`             // 返佣倍数
	Amount             int64   `json:"amount"`               // 权益卡总返佣金额,单位:分
	Days               int64   `json:"days"`                 // 权益卡总返天数
	AmountRemain       int64   `json:"amount_remain"`        // 权益卡剩余总返佣金额,单位:分
	DaysRemain         int64   `json:"days_remain"`          // 权益卡剩余总返天数
	DayRatio           float64 `json:"day_ratio"`            // 每日返佣比例
	DayMoney           int64   `json:"day_money"`            // 每日返佣金额,单位:分
	AmountAchieved     int64   `json:"amount_achieved"`      // 累计已返金额,单位:分
	AmountAchievedDays int64   `json:"amount_achieved_days"` // 已返天数
	Discount           float64 `json:"discount"`             // 商城购物折扣
	State              int     `json:"state"`                // 状态
	StateMsg           string  `json:"state_msg"`            // 状态描述
	CreatedAt          int64   `json:"created_at"`           // 创建时间戳
	UpdatedAt          int64   `json:"updated_at"`           // 更新时间戳
	Note               string  `json:"note"`                 // 备注
	Tips               string  `json:"tips"`                 // 提示
	Card               *string `json:"card"`                 // 权益卡信息
}

// ImsVxFanRightCardUsing ims_vx_fan_right_card_using 微信粉丝权益卡使用中表
type ImsVxFanRightCardUsing struct {
	Id                 int64   `json:"id"`                   // id
	Cid                int64   `json:"cid"`                  // 权益卡id
	Uid                int64   `json:"uid"`                  // 会员id
	Serial             string  `json:"serial"`               // 订单号
	Name               string  `json:"name"`                 // 权益卡名称
	Money              int64   `json:"money"`                // 权益卡金额,单位:分
	Multiple           float64 `json:"multiple"`             // 返佣倍数
	Amount             int64   `json:"amount"`               // 权益卡总返佣金额,单位:分
	Days               int64   `json:"days"`                 // 权益卡总返天数
	AmountRemain       int64   `json:"amount_remain"`        // 权益卡剩余总返佣金额,单位:分
	DaysRemain         int64   `json:"days_remain"`          // 权益卡剩余总返天数
	DayRatio           float64 `json:"day_ratio"`            // 每日返佣比例
	DayMoney           int64   `json:"day_money"`            // 每日返佣金额,单位:分
	AmountAchieved     int64   `json:"amount_achieved"`      // 累计已返金额,单位:分
	AmountAchievedDays int64   `json:"amount_achieved_days"` // 已返天数
	Discount           float64 `json:"discount"`             // 商城购物折扣
	State              int     `json:"state"`                // 状态
	StateMsg           string  `json:"state_msg"`            // 状态描述
	CreatedAt          int64   `json:"created_at"`           // 创建时间戳
	UpdatedAt          int64   `json:"updated_at"`           // 更新时间戳
	Note               string  `json:"note"`                 // 备注
	Tips               string  `json:"tips"`                 // 提示
	Card               *string `json:"card"`                 // 权益卡信息
}

// ImsVxFanSign ims_vx_fan_sign 微信粉丝签到表
type ImsVxFanSign struct {
	Id            int64  `json:"id"`             // id
	Uid           int64  `json:"uid"`            // 粉丝id
	Val1          int64  `json:"val1"`           // 变化前的值
	Val2          int64  `json:"val2"`           // 本次变化值
	Val3          int64  `json:"val3"`           // 变化前的后
	Day           int64  `json:"day"`            // 签到日,如:20010101
	ContinuedDays int64  `json:"continued_days"` // 持续日数
	State         int    `json:"state"`          // 状态
	StateMsg      string `json:"state_msg"`      // 状态描述
	CreatedAt     int64  `json:"created_at"`     // 创建时间戳
	UpdatedAt     int64  `json:"updated_at"`     // 更新时间戳
	Note          string `json:"note"`           // 备注
}

// ImsVxFanSignAdvertising ims_vx_fan_sign_advertising 微信签到广告表
type ImsVxFanSignAdvertising struct {
	Id          int64   `json:"id"`           // id
	AccountId   int64   `json:"account_id"`   // 微信公众账号id
	Name        string  `json:"name"`         // 广告名称
	Avatar      string  `json:"avatar"`       // 广告标识图
	Video       string  `json:"video"`        // 广告视频地址
	Details     *string `json:"details"`      // 广告详情
	Duration    int64   `json:"duration"`     // 广告时长:秒
	Category    string  `json:"category"`     // 广告类型
	CategoryMsg string  `json:"category_msg"` // 广告类型描述
	Uid         int64   `json:"uid"`          // 所属用户id
	State       int     `json:"state"`        // 状态
	StateMsg    string  `json:"state_msg"`    // 状态描述
	CreatedAt   int64   `json:"created_at"`   // 创建时间戳
	UpdatedAt   int64   `json:"updated_at"`   // 更新时间戳
	Note        string  `json:"note"`         // 备注
}

// ImsVxFanVip ims_vx_fan_vip 微信粉丝等级表
type ImsVxFanVip struct {
	Id        int64  `json:"id"`         // id
	Uid       int64  `json:"uid"`        // 粉丝id
	Vip       string `json:"vip"`        // 粉丝等级
	VipMsg    string `json:"vip_msg"`    // 粉丝等级描述
	State     int    `json:"state"`      // 状态
	StateMsg  string `json:"state_msg"`  // 状态描述
	CreatedAt int64  `json:"created_at"` // 创建时间戳
	UpdatedAt int64  `json:"updated_at"` // 更新时间戳
	ExpiredAt int64  `json:"expired_at"` // vip等级过期时间戳
	Note      string `json:"note"`       // 备注
}

// ImsVxFanWithdrawLog ims_vx_fan_withdraw_log 微信粉丝账户零钱提现明细表
type ImsVxFanWithdrawLog struct {
	Id          int64  `json:"id"`           // id
	Uid         int64  `json:"uid"`          // 粉丝id
	Val1        int64  `json:"val1"`         // 变化前的值
	Val2        int64  `json:"val2"`         // 本次变化值
	Val3        int64  `json:"val3"`         // 变化前的后
	Category    int    `json:"category"`     // 明细类型
	CategoryMsg string `json:"category_msg"` // 明细类型描述
	OrderId     int64  `json:"order_id"`     // 对应订单id
	OrderNo     string `json:"order_no"`     // 对应订单号
	State       int    `json:"state"`        // 状态
	StateMsg    string `json:"state_msg"`    // 状态描述
	CreatedAt   int64  `json:"created_at"`   // 创建时间戳
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间戳
	Note        string `json:"note"`         // 备注
}

// ImsVxRedBag ims_vx_red_bag 微信公众号红包发送记录表
type ImsVxRedBag struct {
	Id           int64  `json:"id"`             // id
	Uid          int64  `json:"uid"`            // 会员id
	AccountId    int64  `json:"account_id"`     // 微信公众账号id
	AccountAppId string `json:"account_app_id"` // 微信公众账号app_id
	MchId        string `json:"mch_id"`         // 微信商户号id
	Openid       string `json:"openid"`         // 粉丝openid
	Amount       int64  `json:"amount"`         // 红包金额,单位:分
	SendListId   string `json:"send_list_id"`   // 发送队列id
	MchNo        string `json:"mch_no"`         // 商户订单号
	ReturnCode   string `json:"return_code"`    // 状态码
	ReturnMsg    string `json:"return_msg"`     // 状态信息
	ResultCode   string `json:"result_code"`    // 业务结果
	ErrCode      string `json:"err_code"`       // 错误代码
	ErrCodeDes   string `json:"err_code_des"`   // 错误代码描述
	Day          int64  `json:"day"`            // 红包发送日
	State        int    `json:"state"`          // 状态
	StateMsg     string `json:"state_msg"`      // 状态描述
	CreatedAt    int64  `json:"created_at"`     // 创建时间戳
	UpdatedAt    int64  `json:"updated_at"`     // 更新时间戳
	Note         string `json:"note"`           // 备注
}

// ImsVxRedBagQuotaCard ims_vx_red_bag_quota_card 微信红包额度卡表
type ImsVxRedBagQuotaCard struct {
	Id        int64  `json:"id"`         // id
	AccountId int64  `json:"account_id"` // 微信公众账号id
	Name      string `json:"name"`       // 红包额度卡名称
	Avatar    string `json:"avatar"`     // 红包额度卡标识图
	Title     string `json:"title"`      // 标题
	Details   string `json:"details"`    // 红包额度卡详情
	Money     int64  `json:"money"`      // 红包额度卡单价,单位:分
	Prepare   int64  `json:"prepare"`    // 红包额度,单位:分
	State     int    `json:"state"`      // 状态
	StateMsg  string `json:"state_msg"`  // 状态描述
	CreatedAt int64  `json:"created_at"` // 创建时间戳
	UpdatedAt int64  `json:"updated_at"` // 更新时间戳
	Note      string `json:"note"`       // 备注
}

// ImsWechatAttachment ims_wechat_attachment
type ImsWechatAttachment struct {
	Id              uint   `json:"id"`                //
	Uniacid         uint   `json:"uniacid"`           //
	Acid            uint   `json:"acid"`              //
	Uid             uint   `json:"uid"`               //
	Filename        string `json:"filename"`          //
	Attachment      string `json:"attachment"`        //
	MediaId         string `json:"media_id"`          //
	Width           uint   `json:"width"`             //
	Height          uint   `json:"height"`            //
	Type            string `json:"type"`              //
	Model           string `json:"model"`             //
	Tag             string `json:"tag"`               //
	Createtime      uint   `json:"createtime"`        //
	ModuleUploadDir string `json:"module_upload_dir"` //
	GroupId         int    `json:"group_id"`          //
}

// ImsWechatNews ims_wechat_news
type ImsWechatNews struct {
	Id                 uint   `json:"id"`                    //
	Uniacid            *uint  `json:"uniacid"`               //
	AttachId           uint   `json:"attach_id"`             //
	ThumbMediaId       string `json:"thumb_media_id"`        //
	ThumbUrl           string `json:"thumb_url"`             //
	Title              string `json:"title"`                 //
	Author             string `json:"author"`                //
	Digest             string `json:"digest"`                //
	Content            string `json:"content"`               //
	ContentSourceUrl   string `json:"content_source_url"`    //
	ShowCoverPic       uint8  `json:"show_cover_pic"`        //
	Url                string `json:"url"`                   //
	Displayorder       int    `json:"displayorder"`          //
	NeedOpenComment    int8   `json:"need_open_comment"`     //
	OnlyFansCanComment int8   `json:"only_fans_can_comment"` //
}

// ImsWxappGeneralAnalysis ims_wxapp_general_analysis
type ImsWxappGeneralAnalysis struct {
	Id              uint   `json:"id"`                //
	Uniacid         int    `json:"uniacid"`           //
	SessionCnt      int    `json:"session_cnt"`       //
	VisitPv         int    `json:"visit_pv"`          //
	VisitUv         int    `json:"visit_uv"`          //
	VisitUvNew      int    `json:"visit_uv_new"`      //
	Type            int8   `json:"type"`              //
	StayTimeUv      string `json:"stay_time_uv"`      //
	StayTimeSession string `json:"stay_time_session"` //
	VisitDepth      string `json:"visit_depth"`       //
	RefDate         string `json:"ref_date"`          //
}

// ImsWxappRegisterVersion ims_wxapp_register_version
type ImsWxappRegisterVersion struct {
	Id          uint   `json:"id"`          //
	Uniacid     int    `json:"uniacid"`     //
	VersionId   int    `json:"version_id"`  //
	Auditid     int    `json:"auditid"`     //
	Version     string `json:"version"`     //
	Description string `json:"description"` //
	Status      int8   `json:"status"`      //
	Reason      string `json:"reason"`      //
	UploadTime  int    `json:"upload_time"` //
	AuditInfo   string `json:"audit_info"`  //
	SubmitInfo  string `json:"submit_info"` //
	Developer   string `json:"developer"`   //
}

// ImsWxappUndocodeauditLog ims_wxapp_undocodeaudit_log
type ImsWxappUndocodeauditLog struct {
	Id         uint `json:"id"`          //
	Uniacid    int  `json:"uniacid"`     //
	VersionId  int  `json:"version_id"`  //
	Auditid    int  `json:"auditid"`     //
	RevokeTime int  `json:"revoke_time"` //
}

// ImsWxappVersions ims_wxapp_versions
type ImsWxappVersions struct {
	Id             uint    `json:"id"`              //
	Uniacid        uint    `json:"uniacid"`         //
	Multiid        uint    `json:"multiid"`         //
	Version        string  `json:"version"`         //
	Description    string  `json:"description"`     //
	Modules        string  `json:"modules"`         //
	DesignMethod   int8    `json:"design_method"`   //
	Template       int     `json:"template"`        //
	Quickmenu      string  `json:"quickmenu"`       //
	Createtime     int     `json:"createtime"`      //
	Type           int     `json:"type"`            //
	EntryId        int     `json:"entry_id"`        //
	Appjson        string  `json:"appjson"`         //
	DefaultAppjson string  `json:"default_appjson"` //
	UseDefault     int8    `json:"use_default"`     //
	LastModules    *string `json:"last_modules"`    //
	Tominiprogram  string  `json:"tominiprogram"`   //
	UploadTime     int     `json:"upload_time"`     //
}

// ImsWxcardReply ims_wxcard_reply
type ImsWxcardReply struct {
	Id        uint   `json:"id"`         //
	Rid       uint   `json:"rid"`        //
	Title     string `json:"title"`      //
	CardId    string `json:"card_id"`    //
	Cid       uint   `json:"cid"`        //
	BrandName string `json:"brand_name"` //
	LogoUrl   string `json:"logo_url"`   //
	Success   string `json:"success"`    //
	Error     string `json:"error"`      //
}
