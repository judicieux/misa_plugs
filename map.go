package l9plugins

import (
        "github.com/LeakIX/l9format"
        "github.com/judicieux/misa_plugs/web"
)

var webPlugins = []l9format.WebPluginInterface{
        web.AwsCredentialsHttpPlugin{},
        web.BootstrapHttpPlugin{},
        web.CodeJsonReply{},
        web.ConfigAssetsHttpPlugin{},
        web.ConfigEnvRubyHttpPlugin{},
        web.ConfigJsonHttp{},
        web.ConfigMediaHttpPlugin{},
        web.ConfigStaticHttpPlugin{},
        web.ConfigSubPathApiHttpPlugin{},
        web.ConfigSubPathDevHttpPlugin{},
        web.ConfigSubPathJsHttpPlugin{},
        web.ConfigSubPathMediaHttpPlugin{},
        web.ConfigSubPathWwwHttpPlugin{},
        web.ConfigYmlHttpPlugin{},
        web.DashboardPhpInfoHttpPlugin{},
        web.DataConfigHttpPlugin{},
        web.DockerComposeProdHttpPlugin{},
        web.DotEnvHttpPlugin{},
        web.DotGitlabciHttpPlugin{},
        web.EnvBackupHttpPlugin{},
        web.EnvBakHttpPlugin{},
        web.EnvDevHttpPlugin{},
        web.EnvDevLocalHttpPlugin{},
        web.EnvDevelopmentLocalHttpPlugin{},
        web.EnvExampleHttpPlugin{},
        web.EnvLiveHttpPlugin{},
        web.EnvLocalHttpPlugin{},
        web.EnvOldHttpPlugin{},
        web.EnvProdHttpPlugin{},
        web.EnvProdLocalHttpPlugin{},
        web.EnvProductionHttpPlugin{},
        web.EnvProductionLocalHttpPlugin{},
        web.EnvRcHttpPlugin{},
        web.EnvSampleHttpPlugin{},
        web.EnvSaveHttpPlugin{},
        web.EnvStageHttpPlugin{},
        web.GitConfigHttpPlugin{},
        web.LinusadminPhpInfoHttpPlugin{},
        web.OldPhpInfoHttpPlugin{},
        web.PInfoHttpPlugin{},
        web.ProfilerPhpInfoHttpPlugin{},
        web.PublicConfigHttpPlugin{},
        web.ServerConfigHttpPlugin{},
        web.SrcConfigHttpPlugin{},
        web.SubApiEnvHttpPlugin{},
        web.SubConfigEnvHttpPlugin{},
        web.TiretPhpInfoHttpPlugin{},
        web.TravisYmlHttpPlugin{},
        web.UnderscorePhpInfoHttpPlugin{},
}

func GetWebPlugins() []l9format.WebPluginInterface {
        return webPlugins
}
