package locales

import "decred.org/dcrdex/client/intl"

var ZhCN = map[string]*intl.Translation{
	"Language":                       {T: "zh-CN"},
	"Markets":                        {T: "市场"},
	"Wallets":                        {T: "钱包"},
	"Notifications":                  {T: "通知"},
	"Recent Activity":                {T: "近期活动"},
	"Sign Out":                       {T: "退出"},
	"Order History":                  {T: "历史订单"},
	"load from file":                 {T: "从文件加载"},
	"loaded from file":               {T: "从文件加载"},
	"defaults":                       {T: "默认"},
	"Wallet Password":                {T: "钱包密码"},
	"w_password_helper":              {T: "这是您使用钱包软件配置的密码。"},
	"w_password_tooltip":             {T: "如果您的钱包不需要密码，请将密码清空。"},
	"App Password":                   {T: "应用程序密码"},
	"Add":                            {T: "添加"},
	"Unlock":                         {T: "解锁"},
	"Rescan":                         {T: "重新扫描"},
	"Wallet":                         {T: "钱包"},
	"app_password_reminder":          {T: "执行敏感的钱包操作时始终需要您的应用程序密码。"},
	"DEX Address":                    {T: "DEX 地址"},
	"TLS Certificate":                {T: "TLS 证书"},
	"remove":                         {T: "删除"},
	"add a file":                     {T: "添加文件"},
	"Submit":                         {T: "提交"},
	"Confirm Registration":           {T: "确认注册"},
	"app_pw_reg":                     {T: "输入应用程序密码以确认注册DEX。"},
	"reg_confirm_submit":             {T: `当您提交此表单时，将从您的钱包中花费资金来支付注册费。`},
	"provided_markets":               {T: "DEX 提供以下市场："},
	"accepted_fee_assets":            {T: "DEX 接受以下费用："},
	"base_header":                    {T: "基础"},
	"quote_header":                   {T: "报价"},
	"lot_size_headsup":               {T: "所有交易都是手数的倍数。"},
	"Password":                       {T: "密码"},
	"Register":                       {T: "注册"},
	"Authorize Export":               {T: "授权导出"},
	"export_app_pw_msg":              {T: "输入您的应用程序密码以确认帐户导出"},
	"Disable Account":                {T: "禁用帐户"},
	"disable_dex_server":             {T: "DEX 服务器可在以后的任何时间重新启用（您无需支付费用），只需再次添加它即可。"},
	"Authorize Import":               {T: "授权导入"},
	"app_pw_import_msg":              {T: "输入您的应用程序密码以确认账户导入"},
	"Account File":                   {T: "账户档案"},
	"Change Application Password":    {T: "更改应用程序密码"},
	"Current Password":               {T: "当前密码"},
	"New Password":                   {T: "新密码"},
	"Confirm New Password":           {T: "确认新密码"},
	"cancel_no_pw":                   {T: "取消剩余订单"},
	"cancel_remain":                  {T: "在取消订单之前，剩余金额可能会发生变化。"},
	"Log In":                         {T: "登录"},
	"epoch":                          {T: "epoch"},
	"price":                          {T: "价格"},
	"volume":                         {T: "交易量"},
	"buys":                           {T: "买入"},
	"sells":                          {T: "卖出"},
	"Buy Orders":                     {T: "买单"},
	"Quantity":                       {T: "数量"},
	"Rate":                           {T: "汇率"},
	"Limit Order":                    {T: "限价单"},
	"Market Order":                   {T: "市价单"},
	"reg_status_msg":                 {T: `为了在<span id="regStatusDex" class="text-break"></span>进行交易，注册费支付需要<span id="confReq"></span>确认。`},
	"Buy":                            {T: "买"},
	"Sell":                           {T: "卖"},
	"lot_size":                       {T: "手数"},
	"Rate Step":                      {T: "当前汇率"},
	"Max":                            {T: "最大"},
	"lot":                            {T: "很多"},
	"min trade is about":             {T: "最小交易额"},
	"immediate_explanation":          {T: "如果订单在下一个匹配周期内不能完全匹配，不匹配的数量将不会被再次匹配。需要重新设置订单。"},
	"Immediate or cancel":            {T: "立即或取消"},
	"Balances":                       {T: "余额"},
	"outdated_tooltip":               {T: "余额可能已过期。连接钱包刷新。"},
	"available":                      {T: "可用"},
	"connect_refresh_tooltip":        {T: "点击连接刷新"},
	"add_a_wallet":                   {T: `添加一个 <span data-tmpl="addWalletSymbol"></span> 钱包`},
	"locked":                         {T: "锁定"},
	"immature":                       {T: "不成熟"},
	"Sell Orders":                    {T: "卖单"},
	"Your Orders":                    {T: "你的订单"},
	"Recent Matches":                 {T: "最近的匹配"},
	"Type":                           {T: "类型"},
	"Side":                           {T: "侧面"},
	"Age":                            {T: "年龄"},
	"Filled":                         {T: "已满"},
	"Settled":                        {T: "稳定的"},
	"Status":                         {T: "状态"},
	"view order history":             {T: "查看历史订单"},
	"cancel_order":                   {T: "取消订单"},
	"order details":                  {T: "订单详情"},
	"verify_order":                   {T: `验证 <span id="vSideHeader"></span> 订单`},
	"You are submitting an order to": {T: "您正在提交订单至"},
	"at a rate of":                   {T: "汇率"},
	"for a total of":                 {T: "共计"},
	"verify_market":                  {T: "这是一个市价单，将匹配交易列表中的最佳限价单。根据当前市场的缺口率，您可能会收到大约"},
	"auth_order_app_pw":              {T: "输入您的应用程序密码授权此订单。"},
	"lots":                           {T: "手数"},
	"order_disclaimer": {T: `<span class="red">请注意</span>：交易需要时间结算，您不能关闭 DEX 客户端软件，
		 以及<span data-quote-ticker></span> 或 <span data-base-ticker></span> 钱包软件，直到
		结算完成。结算可能在几分钟内完成，也可能需要几个小时完成。`},
	"Order":                     {T: "订单"},
	"see all orders":            {T: "查看所有订单"},
	"Exchange":                  {T: "交易"},
	"Market":                    {T: "市场"},
	"Offering":                  {T: "提供"},
	"Asking":                    {T: "询问"},
	"Fees":                      {T: "费用"},
	"order_fees_tooltip":        {T: "链上交易费用，通常由矿工收取。Decred DEX 不收取任何交易费用。"},
	"Matches":                   {T: "匹配"},
	"Match ID":                  {T: "匹配 ID"},
	"Time":                      {T: "时间"},
	"ago":                       {T: "以前"},
	"Cancellation":              {T: "取消"},
	"Order Portion":             {T: "订单部分"},
	"you":                       {T: "你"},
	"them":                      {T: "他们"},
	"Redemption":                {T: "赎回"},
	"Refund":                    {T: "退款"},
	"Funding Coins":             {T: "资金硬币"},
	"Exchanges":                 {T: "交易"},
	"apply":                     {T: "申请"},
	"Assets":                    {T: "资产"},
	"Trade":                     {T: "交易"},
	"Set App Password":          {T: "设置应用程序密码"},
	"reg_set_app_pw_msg":        {T: "设置您的应用程序密码。此密码将保护您的 DEX 帐户密钥和连接的钱包。"},
	"Password Again":            {T: "再次输入密码"},
	"Add a DEX":                 {T: "添加 DEX"},
	"Pick a server":             {T: "选择服务器"},
	"reg_ssl_needed":            {T: "看起来我们没有此 DEX 的 SSL 证书。添加服务器的证书以继续。"},
	"Dark Mode":                 {T: "暗模式"},
	"Show pop-up notifications": {T: "显示弹出通知"},
	"Account ID":                {T: "账户 ID"},
	"Export Account":            {T: "退出账户"},
	"simultaneous_servers_msg":  {T: "Decred DEX 客户端支持同时使用任意数量的 DEX 服务器。"},
	"Change App Password":       {T: "更改应用程序密码"},
	"Build ID":                  {T: "构建 ID"},
	"Connect":                   {T: "连接"},
	"Send":                      {T: "发送"},
	"Deposit":                   {T: "存款"},
	"Lock":                      {T: "锁定"},
	"New Deposit Address":       {T: "新存款地址"},
	"Address":                   {T: "地址"},
	"Amount":                    {T: "金额"},
	"Authorize the transfer with your app password.": {T: "输入您的应用程序密码授权转移。"},
	"Reconfigure":                 {T: "重新配置"},
	"pw_change_instructions":      {T: "修改下面的密码不会改变你钱包软件的密码。直接用钱包软件修改密码后，使用此表单更新DEX客户端。"},
	"New Wallet Password":         {T: "新的钱包密码"},
	"pw_change_warn":              {T: "注意：在进行活跃交易时更改为不同的钱包可能会导致资金丢失。"},
	"Show more options":           {T: "显示更多选项"},
	"seed_implore_msg":            {T: "您应该仔细记下您的应用程序种子并保存一份副本。如果您无法访问本机或关键应用程序文件，该种子可用于恢复您的 DEX 账户和原生钱包。一些旧账户不能从种子中恢复，无论是旧的还是新的，最好将帐户密钥与种子分开备份。"},
	"View Application Seed":       {T: "查看应用程序种子"},
	"Remember my password":        {T: "记住密码"},
	"pw_for_seed":                 {T: "输入您的应用程序密码以显示您的种子。请确保没有其他人可以看到您的屏幕。"},
	"Asset":                       {T: "资产"},
	"Balance":                     {T: "余额"},
	"Actions":                     {T: "行为记录"},
	"Restoration Seed":            {T: "恢复种子"},
	"Restore from seed":           {T: "从种子中恢复"},
	"Import Account":              {T: "导入账户"},
	"no_wallet":                   {T: "没有钱包"},
	"create_a_x_wallet":           {T: "创建 <span data-asset-name=1></span> 钱包"},
	"dont_share":                  {T: "不要分享，不要丢失。"},
	"Show Me":                     {T: "查看"},
	"Wallet Settings":             {T: "钱包设置"},
	"add_a_x_wallet":              {T: `添加一个 <img data-tmpl="assetLogo" class="small-icon mx-1"> <span data-tmpl="assetName"></span> 钱包`},
	"off":                         {T: "关"},
	"Export Trades":               {T: "退出交易"},
	"change the wallet type":      {T: "更改钱包类型"},
	"confirmations":               {T: "确认"},
	"pick a different asset":      {T: "选择其它的资产"},
	"Create":                      {T: "创建"},
	"1 Sync the Blockchain":       {T: "1: 同步区块链"},
	"Progress":                    {T: "进度"},
	"remaining":                   {T: "剩余"},
	"Your Deposit Address":        {T: "您钱包的存款地址"},
	"add a different server":      {T: "添加不同的服务器"},
	"Add a custom server":         {T: "添加自定义服务器"},
	"plus tx fees":                {T: "+ 交易费"},
	"Export Seed":                 {T: "退出种子"},
	"Total":                       {T: "总计"},
	"Trading":                     {T: "交易"},
	"Receiving Approximately":     {T: "大约接收"},
	"Fee Projection":              {T: "费用预测"},
	"details":                     {T: "细节"},
	"to":                          {T: "到"},
	"Options":                     {T: "选项"},
	"fee_projection_tooltip":      {T: "如果在您的订单匹配之前网络条件没有改变，则总费用应该在这个范围内。"},
	"unlock_for_details":          {T: "解锁您的钱包以检索订单详细信息和其它选项。"},
	"estimate_unavailable":        {T: "订单估算和选项不可用"},
	"Fee Details":                 {T: "费用明细"},
	"estimate_market_conditions":  {T: "最佳和最坏情况的估计基于当前的网络状况，并且可能会在订单匹配时发生变化。"},
	"Best Case Fees":              {T: "最佳案例费用"},
	"best_case_conditions":        {T: "最好的案例费用发生在整个订单在单个匹配中被消耗掉时。"},
	"Swap":                        {T: "交换"},
	"Redeem":                      {T: "兑换"},
	"Worst Case Fees":             {T: "最坏情况费用"},
	"worst_case_conditions":       {T: "如果订单在多个时期内一次匹配一批，则可能发生最坏的情况。"},
	"Maximum Possible Swap Fees":  {T: "可能出现的最大交换费用"},
	"max_fee_conditions":          {T: "这是您交易而支付的最高费用。费用通常按此费率的一小部分评估。一旦您下订单，最高金额不会更改。"},
	"wallet_logs":                 {T: "钱包日志"},
	"accelerate_order":            {T: "加速订单"},
	"acceleration_text":           {T: "如果你的交易卡住了，你可以尝试通过额外的交易来加速它们。当现有未确认交易的费率变得太低而无法在下一个区块中开采时，这很有帮助。当您提交此表单时，将创建一个交易，提高网络的费用。选择一个足以包含在下一个区块中的费率。请查看区块浏览器用以确认。"},
	"effective_swap_tx_rate":      {T: "有效交易费率"},
	"current_fee":                 {T: "当前建议费率"},
	"accelerate_success":          {T: `成功提交交易：<span id="accelerateTxID"></span>`},
	"accelerate":                  {T: "加速"},
	"acceleration_transactions":   {T: "加速交易"},
	"acceleration_cost_msg":       {T: `将有效费率提高到 <span id="feeRateEstimate"></span> 将花费 <span id="feeEstimate"></span>`},
	"recent_acceleration_msg":     {T: `你最近的加速是 <span id="recentAccelerationTime"></span> 分钟前！您确定要加速吗？`},
	"recent_swap_msg":             {T: `您最早的未确认交换交易在 <span id="recentSwapTime"></span> 分钟前提交！您确定要加速吗？`},
	"early_acceleration_help_msg": {T: `这不会对您的订单造成损害，但您可能会浪费金钱。仅当现有未确认交易的费率变得太低而无法在下一个区块中开采时，加速才有帮助，但如果区块只是缓慢开采，则无济于事。您可以通过关闭此弹出窗口并单击您以前的交易在区块浏览器中确认这一点。`},
	"recover":                     {T: "恢复"},
	"recover_wallet":              {T: "恢复钱包"},
	"recover_warning":             {T: "恢复钱包会将所有钱包数据移至备份文件夹。您必须等到钱包与网络重新同步，这可能需要很长时间才能再次使用钱包。"},
	"wallet_actively_used":        {T: "钱包正在使用中！"},
	"confirm_force_message":       {T: "此钱包正在主动管理订单。执行此操作后，重新同步您的钱包需要很长时间，可能导致订单失败。仅在必要时才执行此操作！"},
	"confirm":                     {T: "确认"},
	"cancel":                      {T: "取消"},
	"Update TLS Certificate":      {T: "更新 TLS 证书"},
	"registered dexes":            {T: "注册的 dexes:"},
	"successful_cert_update":      {T: "成功更新证书！"},
	"update dex host":             {T: "更新 DEX 主机"},
	"copied":                      {T: "复制！"},
	"export_wallet":               {T: "导出钱包"},
	"pw_for_wallet_seed":          {T: "输入您的应用程序密码以显示钱包种子。确保没有其他人可以看到您的屏幕。如果有人可以访问钱包种子，他们将能够窃取您的所有资金。"},
	"export_wallet_disclaimer":    {T: `<span class="warning-text">在 DEX 中运行交易时使用外部恢复的钱包可能会导致交易失败和资金损失。</span> 建议您不要导出你的钱包，除非你是一个有经验的用户并且你知道在做什么。`},
	"export_wallet_msg":           {T: "以下是在第三方钱包中进行恢复所需的种子。当您在 DEX 上进行活跃交易时，请勿使用您的外部钱包进行交易。"},
	"clipboard_warning":           {T: "复制/粘贴钱包种子是一个潜在的安全风险。这样做需要您自担风险。"},
	"fiat_exchange_rate_sources":  {T: "法定汇率来源"},
	"Synchronizing":               {T: "同步"},
	"wallet_wait_synced":          {T: "同步后会创建钱包"},
	"Create a Wallet":             {T: "创建钱包"},
	"Receive":                     {T: "接收"},
	":title:ready":                {T: "准备就绪"},
	"ready":                       {T: "已准备"},
	"Wallet Type":                 {T: "钱包类型"},
	"Peer Count":                  {T: "节点数"},
	"Sync Progress":               {T: "同步进度"},
	"Settings":                    {T: "设置"},
	"asset_name Markets":          {T: "<span data-asset-name=1></span> 市场"},
	"Host":                        {T: "主机"},
	"No Recent Activity":          {T: "最近没有活动"},
	"Recent asset_name Activity":  {T: "最近的 <span data-asset-name=1></span> 活动"},
	"other_actions":               {T: "其它活动"},
	"estimated_fee":               {T: "预估费用"},
	"estimated_total_spend":       {T: "预估总支出"},
	"estimated_balance":           {T: "完成之后的预估余额"},
	"max_estimated_send":          {T: "预估最多发送"},
	"max_estimated_send_fee":      {T: "预估最多发送费用"},
	"sending":                     {T: "发送"},
	"transfer":                    {T: "交易"},
	"invalid_address":             {T: "注意：提供一个有效的地址。"},
	"max_estimated_send_tooltip":  {T: "这是您在选中“从发送的金额中减去费用”的情况下将收到的预估金额。如果没有减去费用复选框，这是您可以发送的最大预估金额。"},
	"Maker Swap":                  {T: "挂单"},
	"Taker Swap":                  {T: "吃单"},
	"Maker Redemption":            {T: "挂单赎回"},
	"Taker Redemption":            {T: "吃单赎回"},
	"Pending":                     {T: "发送中"},
	"disable_wallet":              {T: "禁用钱包"},
	"enable_wallet":               {T: "启用钱包"},
	"disable_wallet_warning":      {T: "注意：该钱包在你启动DEX客户端软件时是不会连接的，只有开启后才能使用。"},
	"enable_wallet_message":       {T: "钱包将恢复运行，可能需要一些时间来同步。"},
	"Disabled":                    {T: "已禁用"},
}
