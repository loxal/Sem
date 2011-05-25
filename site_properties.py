
# User related constants
AUTHOR = _('Alexander Orlov')
TITLE = _('synergy of loxal')
# sub title
TITLE_DESC = _('...more than just the sum of its parts.')
# administrative content, also responsible for site content
IMPRINT = '''
    Alexander Orlov

    Rablstr. 12
    81669 Munich
    Germany
'''

# if you want to change the favicon.ico, you have to do it manually,
# by replacing the favicon.ico file in the static folder

MAIL_RECEIVER = 'alexander.orlov@loxal.net'
MAIL_GAE_ACCOUNT = 'alexander.orlov@loxal.net'
# to built an user relationship and bind the user to the site
GADGET_OPENSOCIAL = '''
<!-- Include the Google Friend Connect javascript library. -->
<script type="text/javascript" src="http://www.google.com/friendconnect/script/friendconnect.js"></script>
<!-- Define the div tag where the gadget will be inserted. -->
<div id="div-1234017250892" style="width:200px;border:1px solid #cccccc;"></div>
<!-- Render the gadget into a div. -->
<script type="text/javascript">
var skin = {};
skin['HEIGHT'] = '200';
skin['BORDER_COLOR'] = '#cccccc';
skin['ENDCAP_BG_COLOR'] = '#e0ecff';
skin['ENDCAP_TEXT_COLOR'] = '#333333';
skin['ENDCAP_LINK_COLOR'] = '#0000cc';
skin['ALTERNATE_BG_COLOR'] = '#ffffff';
skin['CONTENT_BG_COLOR'] = '#ffffff';
skin['CONTENT_LINK_COLOR'] = '#0000cc';
skin['CONTENT_TEXT_COLOR'] = '#333333';
skin['CONTENT_SECONDARY_LINK_COLOR'] = '#7777cc';
skin['CONTENT_SECONDARY_TEXT_COLOR'] = '#666666';
skin['CONTENT_HEADLINE_COLOR'] = '#333333';
google.friendconnect.container.setParentUrl('/' /* location of rpc_relay.html and canvas.html */);
google.friendconnect.container.renderMembersGadget(
 { id: 'div-1234017250892',
   site: '14005686145006081345' },
  skin);
</script>
'''
# to animate the user to interact with the site and provide feedback 
GADGET_OPENSOCIAL_FEEDBACK = '''
    <!-- Include the Google Friend Connect javascript library. -->
    <script type="text/javascript" src="http://www.google.com/friendconnect/script/friendconnect.js"></script>
    <!-- Define the div tag where the gadget will be inserted. -->
    <div id="div-1234017525776" style="width:222px;border:1px solid #cccccc;"></div>
    <!-- Render the gadget into a div. -->
    <script type="text/javascript">
    var skin = {};
    skin['BORDER_COLOR'] = '#cccccc';
    skin['ENDCAP_BG_COLOR'] = '#e0ecff';
    skin['ENDCAP_TEXT_COLOR'] = '#333333';
    skin['ENDCAP_LINK_COLOR'] = '#0000cc';
    skin['ALTERNATE_BG_COLOR'] = '#ffffff';
    skin['CONTENT_BG_COLOR'] = '#ffffff';
    skin['CONTENT_LINK_COLOR'] = '#0000cc';
    skin['CONTENT_TEXT_COLOR'] = '#333333';
    skin['CONTENT_SECONDARY_LINK_COLOR'] = '#7777cc';
    skin['CONTENT_SECONDARY_TEXT_COLOR'] = '#666666';
    skin['CONTENT_HEADLINE_COLOR'] = '#333333';
    skin['DEFAULT_COMMENT_TEXT'] = 'How are you?';
    skin['HEADER_TEXT'] = 'Impressions';
    skin['POSTS_PER_PAGE'] = '5';
    google.friendconnect.container.setParentUrl('/' /* location of rpc_relay.html and canvas.html */);
    google.friendconnect.container.renderWallGadget(
     { id: 'div-1234017525776',
       site: '14005686145006081345',
       'view-params':{"disableMinMax":"true","scope":"SITE","allowAnonymousPost":"true","features":"video,comment","startMaximized":"true"}
     },
      skin);
    </script>
'''
# to involve users in interactive tasks on the site
GADGET_OPENSOCIAL_INTERACTIVITY = '''
    <!-- Include the Google Friend Connect javascript library. -->
    <script type="text/javascript" src="http://www.google.com/friendconnect/script/friendconnect.js"></script>
    <!-- Define the div tag where the gadget will be inserted. -->
    <div id="div-1234020036810" style="width:300px;border:1px solid #cccccc;"></div>
    <!-- Render the gadget into a div. -->
    <script type="text/javascript">
    google.friendconnect.container.setParentUrl('/' /* location of rpc_relay.html and canvas.html */);
    google.friendconnect.container.renderOpenSocialGadget(
     { id: 'div-1234020036810',
       url:'http://os.ilike.com/gadget/playlist',
       site: '18234698991911944501' });
    </script>
'''
# main (financial) resources gaining entity
SPONSOR = '''
    <!-- Google AdSense/ -->
    <script type="text/javascript"><!--
        google_ad_client = "pub-7462302732712971";
        /* sol.header.half-banner */
        google_ad_slot = "4241123125";
        google_ad_width = 234;
        google_ad_height = 60;
        //-->
    </script>
    <script type="text/javascript"
            src="http://pagead2.googlesyndication.com/pagead/show_ads.js">
    </script>
'''
# (financial) resources gaining entity
SUPPORT_CAMPAIGN = '''
    <!-- PayPal Donate/ -->
    <form action="https://www.paypal.com/cgi-bin/webscr" method="post">
    <div>
    <input type="hidden" name="cmd" value="_s-xclick"/>
    <input type="hidden" name="hosted_button_id" value="2578449"/>
    <input style="border: 0em; background: transparent; float: right;" type="image" src="https://www.paypal.com/en_US/i/btn/btn_donate_SM.gif" name="submit" alt="Donate! :)" title="Donate! :)"/>
    <img src="https://www.paypal.com/en_US/i/scr/pixel.gif" width="0" height="0" alt="."/>
    </div>
    </form>
'''
# bookmark/share/redistribution functionality
BOOKMARK = '''
<a href="http://www.addthis.com/bookmark.php?v=250&amp;username=loxal" class="addthis_button_compact">Share</a>
<script type="text/javascript">var addthis_config = {"data_track_clickback":true};</script>
<script type="text/javascript" src="http://s7.addthis.com/js/250/addthis_widget.js#username=loxal"></script>

'''
# to track the visitors
ANALYTICS = '''
        <script type="text/javascript"> <!--Google Analytics Tracker-->
            var _gaq = _gaq || [];
            _gaq.push(
            ['_setAccount', 'UA-7363751-1'],
            ['_trackPageview']
            );
            (function() {
            var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
            ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') +
            '.google-analytics.com/ga.js';
            (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(ga);
            })();
        </script>
'''

# domain search
SEARCH = '''
    <form action="http://www.loxal.net/search.html" id="cse-search-box">
      <div>
        <input type="hidden" name="cx" value="partner-pub-7462302732712971:88unvq-4d13" />
        <input type="hidden" name="cof" value="FORID:9" />
        <input type="hidden" name="ie" value="UTF-8" />
        <input type="text" name="q" size="22" />
        <input type="submit" name="sa" value="Search" />
      </div>
    </form>
    <script type="text/javascript" src="http://www.google.com/coop/cse/brand?form=cse-search-box&amp;lang=en"></script>
'''
# domain search results which are displayed on a separate page
SEARCH_RESULTS = '''
    <div id="cse-search-results"></div>
    <script type="text/javascript">
      var googleSearchIframeName = "cse-search-results";
      var googleSearchFormName = "cse-search-box";
      var googleSearchFrameWidth = 800;
      var googleSearchDomain = "www.google.com";
      var googleSearchPath = "/cse";
    </script>
    <script type="text/javascript" src="http://www.google.com/afsonline/show_afs_search.js"></script>
'''

# not user related configuration constants
TPL_DIR = '../templates'
TPL_MAIN = 'main.html'
TPL_MAIN_SERVICE = '../' + TPL_DIR + '/' + TPL_MAIN
TPL_SVC_DIR = 'svc'
TPL_404_NOT_FOUND = '/404-Not-Found.html'
RES_DIR = '/static'

# main properties of every template
template = {
'res_dir'                           : RES_DIR,
'tpl_main'                          : TPL_MAIN,
'tpl_main_service'                  : TPL_MAIN_SERVICE,
'author'                            : AUTHOR,
'title'                             : TITLE,
'title_desc'                        : TITLE_DESC,
'imprint'                           : IMPRINT,
'gadget_opensocial'                 : GADGET_OPENSOCIAL,
'gadget_opensocial_feedback'        : GADGET_OPENSOCIAL_FEEDBACK,
'gadget_opensocial_interactivity'   : GADGET_OPENSOCIAL_INTERACTIVITY,
'sponsor'                           : SPONSOR,
'analytics'                         : ANALYTICS,
'search'                            : SEARCH,
'search_results'                    : SEARCH_RESULTS,
'support_campaign'                  : SUPPORT_CAMPAIGN,
'bookmark'                          : BOOKMARK,
}
